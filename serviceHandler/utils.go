package serviceHandler

import (
	"github.com/ZeroTechh/UserService/core/types"
	proto "github.com/ZeroTechh/VelocityCore/proto/UserService"
	"github.com/jinzhu/copier"
)

// copies core/types structs to proto get response
func copyToGetResponse(
	main types.Main,
	extra types.Extra,
	meta types.Meta) (response *proto.GetResponse) {
	copier.Copy(&response.Main, &main)
	copier.Copy(&response.Extra, &extra)
	copier.Copy(&response.Meta, &meta)
	return
}

// proto structs to copies core/types structs
func copyFromRequest(
	mainRequest *proto.UserMain,
	extraRequest *proto.UserExtra,
	metaRequest *proto.UserMeta) (
	main types.Main,
	extra types.Extra,
	meta types.Meta) {
	copier.Copy(&main, &mainRequest)
	copier.Copy(&extra, &extraRequest)
	copier.Copy(&meta, &metaRequest)
	return
}
