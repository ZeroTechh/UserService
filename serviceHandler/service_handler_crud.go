package serviceHandler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserService"
	"github.com/ZeroTechh/blaze"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

// Add is used to handle Add function
func (userService UserService) Add(
	ctx context.Context,
	request *proto.AddRequest) (*proto.AddResponse, error) {
	funcLog := blaze.NewFuncLog(
		"service.Add",
		log,
		zap.Any("Request", request),
	)
	funcLog.Started()

	main, extra, _ := copyFromRequest(
		request.MainData, request.ExtraData, nil)

	userID, msg := userService.user.Add(main, extra)
	if msg != "" {
		funcLog.Completed(zap.String("Cause", msg))
		return &proto.AddResponse{Message: msg}, nil
	}

	funcLog.Completed(
		zap.Any("Main Data", main),
		zap.Any("Extra Data", extra),
		zap.String("UserID", userID),
	)

	return &proto.AddResponse{UserID: userID}, nil
}

// Get is used to handle Get function
func (userService UserService) Get(
	ctx context.Context, request *proto.GetRequest) (
	*proto.GetResponse, error) {
	funcLog := blaze.NewFuncLog(
		"service.Get",
		log,
		zap.Any("Request", request),
	)
	funcLog.Started()

	filter := getFilter(request.UserID, request.Username, request.Email)

	main, extra, meta, msg := userService.decodeData(
		filter, request.Type.String())

	mainResp, extraResp, metaResp := copyToResponse(main, extra, meta)

	funcLog.Completed(
		zap.Any("Main", main),
		zap.Any("Exta", extra),
		zap.Any("Meta", meta),
		zap.String("Message", msg),
	)

	return &proto.GetResponse{
		Main:    mainResp,
		Extra:   extraResp,
		Meta:    metaResp,
		Message: msg,
	}, nil
}

// Update is used to handle Update function
func (userService UserService) Update(
	ctx context.Context,
	request *proto.UpdateRequest) (
	response *proto.UpdateResponse,
	err error) {
	response = &proto.UpdateResponse{}
	log.Debug("Updating User Main Data")
	var update config.UserMain

	copier.Copy(&update, &request.Update)

	log.Debug(
		"Updating User Main Data",
		zap.String("UserID", request.UserID),
		zap.Any("Update", request.Update),
	)

	userService.user.Update(request.UserID,
		update,
		config.DBConfigUserMainDataCollection)

	log.Info(
		"Updated User Main Data",
		zap.String("UserID", request.UserID),
		zap.Any("Update", request.Update),
	)

	return
}
