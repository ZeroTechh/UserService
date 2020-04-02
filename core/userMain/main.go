package userMain

import (
	"context"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/ZeroTechh/UserService/core/validate"
	"github.com/ZeroTechh/VelocityCore/logger"
	"github.com/ZeroTechh/VelocityCore/utils"
	"github.com/ZeroTechh/hades"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// all the configs
	config         = hades.GetConfig("main.yaml", []string{"config", "../../config"})
	dbConfig       = config.Map("database")
	mainCollection = dbConfig.Map("collections").Str("main")
	log            = logger.GetLogger(
		config.Map("service").Str("lowLevelLogFile"),
		config.Map("service").Bool("debug"),
	)
	messages       = config.Map("messages")
	invalidDataMsg = messages.Str("invalidUserData")
)

// New returns a new main handler struct
func New() *Main {
	main := Main{}
	main.init()
	return &main
}

// Main is used to handle user main data
type Main struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// init initializes client and database
func (main *Main) init() {
	main.client = utils.CreateMongoDB(dbConfig.Str("address"), log)
	main.database = main.client.Database(dbConfig.Str("db"))
	main.collection = main.database.Collection(mainCollection)
}

func (main Main) generateID() string {
	userIDExists := true
	var userID uuid.UUID

	for userIDExists {
		userID, _ = uuid.NewRandom()
		userIDExists = main.Exists(
			types.Main{UserID: userID.String()},
		)
	}

	return userID.String()
}

// Exists checks if user with certain field exists
func (main Main) Exists(filter types.Main) bool {
	return main.Get(filter) != types.Main{}
}

// Create is used to add new main data
func (main Main) Create(data types.Main) string {
	if !validate.IsValid(data, "main", false) {
		return invalidDataMsg
	}

	usernameExists := main.Exists(types.Main{Username: data.Username})
	emailExists := main.Exists(types.Main{Email: data.Email})
	if usernameExists {
		return messages.Str("usernameExists")
	} else if emailExists {
		return messages.Str("emailExists")
	}

	main.collection.InsertOne(context.TODO(), data)
	return ""
}

// Get is used to a users data
func (main Main) Get(filter types.Main) (data types.Main) {
	main.collection.FindOne(context.TODO(), filter).Decode(&data)
	return
}

// Update updates user's main data
func (main Main) Update(filter types.Main, update types.Main) string {
	if !validate.IsValid(update, "main", true) {
		return invalidDataMsg
	}
	main.collection.UpdateOne(
		context.TODO(),
		filter,
		map[string]types.Main{"$set": update},
	)
	return ""
}

// Auth is used to authenticate username, email or password
func (main Main) Auth(username, email, password string) bool {
	filter := types.Main{Username: username, Email: email}
	data := main.Get(filter)
	return data.Password == password && data != types.Main{}
}
