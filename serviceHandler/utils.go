package serviceHandler

import (
	"github.com/ZeroTechh/UserService/core/types"
	proto "github.com/ZeroTechh/VelocityCore/proto/UserService"
	"github.com/jinzhu/copier"
)

func getFilter(userID, username, email string) types.UserMain {
	filter := types.UserMain{UserID: userID}
	if username != "" {
		filter = types.UserMain{Username: username}
	} else if email != "" {
		filter = types.UserMain{Email: email}
	}
	return filter
}

// copies core/types structs to proto structs
func copyToResponse(
	main types.UserMain,
	extra types.UserExtra,
	meta types.UserMeta) (
	mainResponse *proto.UserMain,
	extraResponse *proto.UserExtra,
	metaResponse *proto.UserMeta) {
	copier.Copy(&mainResponse, &main)
	copier.Copy(&extraResponse, &extra)
	copier.Copy(&metaResponse, &meta)
	return
}

// proto structs to copies core/types structs
func copyFromRequest(
	mainRequest *proto.UserMain,
	extraRequest *proto.UserExtra,
	metaRequest *proto.UserMeta) (
	main types.UserMain,
	extra types.UserExtra,
	meta types.UserMeta) {
	copier.Copy(&main, &mainRequest)
	copier.Copy(&extra, &extraRequest)
	copier.Copy(&meta, &metaRequest)
	return
}
