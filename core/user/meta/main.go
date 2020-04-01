package meta

import (
	"context"
	"time"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/ZeroTechh/VelocityCore/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// generates meta data for a user
func generate(userID string) types.Meta {
	return types.Meta{
		UserID:             userID,
		AccountStatus:      accountStatuses.Str("unverified"),
		AccountCreationUTC: time.Now().Unix(),
	}
}

// NewMeta returns a new meta handler struct
func NewMeta() *Meta {
	meta := Meta{}
	meta.init()
	return &meta
}

// Meta is used to handle user meta data
type Meta struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// init initializes client and database
func (meta *Meta) init() {
	meta.client = utils.CreateMongoDB(dbConfig.Str("address"), log)
	meta.database = meta.client.Database(dbConfig.Str("db"))
	meta.collection = meta.database.Collection(metaCollection)
}

// Create creates and adds new meta data for a user
func (meta Meta) Create(userID string) {
	data := generate(userID)
	meta.collection.InsertOne(context.TODO(), data)
}

// Get returns user's meta data
func (meta Meta) Get(userID string) (data types.Meta) {
	filter := types.Meta{UserID: userID}
	meta.collection.FindOne(context.TODO(), filter).Decode(&data)
	return
}

// Activate marks user as verified
func (meta Meta) Activate(userID string) {
	update := types.Meta{
		AccountStatus: accountStatuses.Str("verified"),
	}
	filter := types.Meta{UserID: userID}
	meta.collection.UpdateOne(
		context.TODO(),
		filter,
		map[string]types.Meta{"$set": update},
	)
}
