package serviceHandler

import (
	"github.com/ZeroTechh/UserService/core/types"
	proto "github.com/ZeroTechh/VelocityCore/proto/UserService"
	"github.com/jinzhu/copier"
)

// gets user data and decodes intro correct type
func (userService UserService) decodeData(
	filter interface{}, dataType string) (
	types.UserMain,
	types.UserExtra,
	types.UserMeta,
	string) {
	var (
		main    types.UserMain
		extra   types.UserExtra
		meta    types.UserMeta
		dataPtr interface{}
	)

	switch dataType {
	case types.Extra:
		dataPtr = &extra
	case types.Meta:
		dataPtr = &meta
	default:
		dataPtr = &main
	}

	collection := collections.Str(dataType)
	msg := userService.user.Get(filter, collection, dataPtr)

	return main, extra, meta, msg
}

func (userService UserService) getUserResponse(
	main types.UserMain,
	extra types.UserExtra,
	meta types.UserExtra) (response *proto.GetResponse) {
	copier.Copy(&response.Main, &main)
	copier.Copy(&response.Extra, &extra)
	copier.Copy(&response.Meta, &meta)
	return response
}
