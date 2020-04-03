package serviceHandler

import (
	"context"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/ZeroTechh/UserService/core/userExtra"
	"github.com/ZeroTechh/UserService/core/userMain"
	"github.com/ZeroTechh/UserService/core/userMeta"
	"github.com/ZeroTechh/VelocityCore/logger"
	proto "github.com/ZeroTechh/VelocityCore/proto/UserService"
	"github.com/ZeroTechh/blaze"
	"github.com/ZeroTechh/hades"
	"go.uber.org/zap"
)

var (
	config = hades.GetConfig("main.yaml", []string{"config", "../config"})
	log    = logger.GetLogger(
		config.Map("service").Str("logFile"),
		config.Map("service").Bool("debug"),
	)
)

// New returns a new service handler
func New() *Handler {
	handler := Handler{}
	handler.init()
	return &handler
}

// Handler is used to handle all user service functions
type Handler struct {
	main  *userMain.Main
	extra *userExtra.Extra
	meta  *userMeta.Meta
}

// Init is used to initialize
func (handler *Handler) init() {
	handler.main = userMain.New()
	handler.extra = userExtra.New()
	handler.meta = userMeta.NewMeta()
}

// Add is used to handle Add function
func (handler Handler) Add(ctx context.Context, request *proto.AddRequest) (*proto.AddResponse, error) {
	funcLog := blaze.NewFuncLog(
		"service.Add",
		log,
		zap.Any("Request", request),
	)
	funcLog.Started()
	main, extra, _ := copyFromRequest(
		request.MainData, request.ExtraData, nil)

	userID := handler.main.GenerateID()
	main.UserID = userID
	extra.UserID = userID

	handler.main.Create(main)
	handler.extra.Create(extra)
	handler.meta.Create(userID)

	funcLog.Completed(zap.String("UserID", userID))
	return &proto.AddResponse{UserID: userID}, nil
}

// Get is used to handle Get function
func (handler Handler) Get(ctx context.Context, request *proto.GetRequest) (*proto.GetResponse, error) {
	funcLog := blaze.NewFuncLog(
		"service.Get",
		log,
		zap.Any("Request", request),
	)
	funcLog.Started()

	var (
		main  types.Main
		extra types.Extra
		meta  types.Meta
	)

	switch request.Type.String() {
	case "extra":
		extra = handler.extra.Get(request.UserID)
	case "meta":
		meta = handler.meta.Get(request.UserID)
	default:
		filter := types.Main{
			UserID:   request.UserID,
			Username: request.Username,
			Email:    request.Email,
		}
		main = handler.main.Get(filter)
	}

	funcLog.Completed(
		zap.Any("Main", main),
		zap.Any("Extra", main),
		zap.Any("Meta", meta),
	)
	return copyToGetResponse(main, extra, meta), nil
}

func (handler Handler) Get(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {