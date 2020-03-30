package user

import (
	"context"

	"github.com/ZeroTechh/VelocityCore/utils"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/ZeroTechh/UserService/core/validate"
)

// NewUser creates a new user struct
func NewUser() *User {
	user := User{}
	user.init()
	return &user
}

// User is used to handle user data
type User struct {
	client   *mongo.Client
	database *mongo.Database
	ctx      context.Context
}

// init initializes client and database
func (user *User) init() {
	user.client = utils.CreateMongoDB(dbConfig.Str("address"), log)
	user.database = user.client.Database(dbConfig.Str("db"))
	user.ctx = context.TODO()
}

// generateID is used to generate an unique user id
func (user User) generateID() string {
	var userID string
	userIDExists := true

	// Creates a new ID and checks if it is already taken
	for userIDExists {
		userID = generateUUID()
		userIDExists = user.exists(
			types.UserMain{UserID: userID}, mainColl)
	}

	return userID
}

// coll returns a collection based on its name
func (user User) coll(name string) *mongo.Collection {
	return user.database.Collection(name)
}

// checks if a user with certain user with certain filter exists
func (user User) exists(filter interface{}, collection string) bool {
	var data interface{}
	user.Get(filter, collection, &data)
	return data != nil
}

// verify verifies user data by checking that username and email dont exist
func (user User) uniqueFieldExists(main types.UserMain) (bool, string) {
	usernameExists := user.exists(
		types.UserMain{Username: main.Username}, mainColl)
	emailExists := user.exists(
		types.UserMain{Email: main.Email}, mainColl)

	if usernameExists {
		return true, messages.Str("usernameExists")
	} else if emailExists {
		return true, messages.Str("emailExists")
	}

	return false, ""
}

// validate validates user data
func (user User) validate(main types.UserMain, extra types.UserExtra) (bool, string) {
	isMainValid := validate.IsValid(main, types.Main, false)
	isExtraValid := validate.IsValid(extra, types.Extra, false)
	if !(isMainValid && isExtraValid) {
		return false, messages.Str("invalidUserData")
	}
	return true, ""
}
