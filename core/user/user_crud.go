package user

import (
	"github.com/ZeroTechh/UserService/core/types"
	"github.com/ZeroTechh/UserService/core/validate"
)

// Add adds a new user then returns user ID or a message
func (user User) Add(main types.UserMain, extra types.UserExtra) (string, string) {
	// Validating data
	valid, msg := user.validate(main, extra)
	if !valid {
		return "", msg
	}

	// Checking if unique fields such as email or username already exists
	uniqueFieldExists, msg := user.uniqueFieldExists(main)
	if uniqueFieldExists {
		return "", msg
	}

	// Generating meta data
	meta := generateMeta()

	// Generating and adding user ID
	userID := user.generateID()
	main.UserID = userID
	extra.UserID = userID
	meta.UserID = userID

	// Adding into database
	user.coll(mainColl).InsertOne(user.ctx, main)
	user.coll(extraColl).InsertOne(user.ctx, extra)
	user.coll(metaColl).InsertOne(user.ctx, meta)

	return userID, ""
}

// Get gets a user based on a filter
func (user User) Get(filter interface{}, collection string, data interface{}) string {
	err := user.coll(collection).FindOne(user.ctx, filter).Decode(data)
	if err != nil {
		return messages.Str("invalidUserData")
	}
	return ""
}

// Update is used to update user data in database
func (user User) Update(filter interface{}, update interface{}, dataType string) string {
	// Checking if update is valid
	if !validate.IsValid(update, dataType, true) {
		return messages.Str("invalidUserData")
	}

	// Updating data in database
	user.coll(collections.Str(dataType)).UpdateOne(
		user.ctx,
		filter,
		map[string]interface{}{"$set": update},
	)

	return ""
}

// Auth validates user's username or email and password
func (user User) Auth(filter interface{}, password string) (bool, string) {
	var userData types.UserMain
	msg := user.Get(filter, mainColl, &userData)
	if msg != "" {
		return false, ""
	}

	// TODO Add password hashing
	if userData.Password == password {
		return true, userData.UserID
	}
	return false, ""
}

// Activate marks user as activated
func (user User) Activate(userID string) string {
	filter := types.UserMeta{UserID: userID}
	update := types.UserMeta{
		AccountStatus: accountStatuses.Str("verified"),
	}
	return user.Update(filter, update, types.Meta)
}
