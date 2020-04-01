package userExtra

import (
	"context"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/ZeroTechh/UserService/core/validate"
	"github.com/ZeroTechh/VelocityCore/logger"
	"github.com/ZeroTechh/VelocityCore/utils"
	"github.com/ZeroTechh/hades"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// all the configs
	config          = hades.GetConfig("main.yaml", []string{"config", "../../../config"})
	dbConfig        = config.Map("database")
	extraCollection = dbConfig.Map("collections").Str("extra")
	log             = logger.GetLogger(
		config.Map("service").Str("lowLevelLogFile"),
		config.Map("service").Bool("debug"),
	)
	invalidDataMsg = config.Map("messages").Str("invalidUserData")
)

// NewExtra returns a new extra handler struct
func NewExtra() *Extra {
	extra := Extra{}
	extra.init()
	return &extra
}

// Extra is used to handle user extra data
type Extra struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// init initializes client and database
func (extra *Extra) init() {
	extra.client = utils.CreateMongoDB(dbConfig.Str("address"), log)
	extra.database = extra.client.Database(dbConfig.Str("db"))
	extra.collection = extra.database.Collection(extraCollection)
}

// Create is used to add new extra data
func (extra Extra) Create(data types.Extra) string {
	if !validate.IsValid(data, "extra", false) {
		return invalidDataMsg
	}
	extra.collection.InsertOne(context.TODO(), data)
	return ""
}

// Get is used to a users data
func (extra Extra) Get(userID string) (data types.Extra) {
	extra.collection.FindOne(
		context.TODO(),
		types.Extra{UserID: userID},
	).Decode(&data)
	return
}

// Update updates user's extra data
func (extra Extra) Update(userID string, update types.Extra) string {
	if !validate.IsValid(update, "extra", true) {
		return invalidDataMsg
	}
	filter := types.Extra{UserID: userID}
	extra.collection.UpdateOne(
		context.TODO(),
		filter,
		map[string]types.Extra{"$set": update},
	)
	return ""
}
