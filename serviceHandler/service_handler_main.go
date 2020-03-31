package serviceHandler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserService"
	"github.com/ZeroTechh/blaze"
	"go.uber.org/zap"

	"github.com/ZeroTechh/UserService/core/user"
)

// UserService is to used to handle user service
type UserService struct {
	user *user.User
}

// init is used to initialize
func (userService *UserService) init() {
	log.Debug("Service Initializing")
	userService.user = user.NewUser()
}

// Auth is used to handle Auth function
func (userService UserService) Auth(
	ctx context.Context,
	request *proto.AuthRequest) (*proto.AuthResponse, error) {
	funcLog := blaze.NewFuncLog(
		"service.Auth",
		log,
		zap.String("Email", request.Email),
		zap.String("Username", request.Username),
	)
	funcLog.Started()

	filter := getFilter("", request.Username, request.Email)
	valid, userID := userService.user.Auth(filter, request.Password)

	funcLog.Completed(
		zap.Bool("Valid", valid),
		zap.String("UserID", userID),
	)

	return &proto.AuthResponse{
		UserID: userID,
		Valid:  valid,
	}, nil
}

// Activate is used handle Activate function
func (userService UserService) Activate(
	ctx context.Context,
	request *proto.ActivateRequest) (*proto.ActivateResponse, error) {
	funcLog := blaze.NewFuncLog(
		"service.Activate",
		log,
		zap.String("Email", request.Email),
	)
	funcLog.Started()

	msg := userService.user.Activate(request.Email)
	funcLog.Completed(zap.String("Message", msg))
	return &proto.ActivateResponse{Message: msg}, nil
}
