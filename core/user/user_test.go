package user

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/ZeroTechh/UserService/core/utils"
)

func TestUser(t *testing.T) {
	assert := assert.New(t)
	user := NewUser()

	// Testing User.Add
	// Testing that invalid user data returns message
	userID, msg := user.Add(types.UserMain{}, types.UserExtra{})
	assert.NotZero(msg)
	assert.Zero(userID)

	// Testing that valid data does not return message
	main, extra := utils.GetMockUserData()
	userID, msg = user.Add(main, extra)
	assert.Zero(msg)
	assert.NotZero(userID)
	main.UserID = userID

	// Testing that already existing username gives a message
	userID, msg = user.Add(main, extra)
	assert.NotZero(msg)
	assert.Zero(userID)

	// Testing that already existing email gives a message
	newMain, newExtra := utils.GetMockUserData()
	newMain.Email = main.Email
	userID, msg = user.Add(newMain, newExtra)
	assert.NotZero(msg)
	assert.Zero(userID)

	// Testing User.Get
	// Testing that for non existing user, message is returned
	var data types.UserMain
	msg = user.Get(types.UserMain{UserID: "Invalid"}, mainColl, &data)
	assert.NotZero(msg)
	assert.Zero(data)

	// Testing that user data is returned for valid filter
	msg = user.Get(types.UserMain{Username: main.Username}, mainColl, &data)
	assert.Zero(msg)
	assert.Equal(main, data)

	// Testing User.Update
	// Testing that for non existing user, message is returned
	msg = user.Update(
		types.UserMain{UserID: "Invalid"},
		types.UserMain{UserID: "Invalid"},
		types.Main,
	)
	assert.NotZero(msg)

	// Testing that update is successful for valid data
	var extraData types.UserExtra
	update := types.UserExtra{FirstName: "newname"}
	filter := types.UserMain{UserID: main.UserID}
	msg = user.Update(
		filter,
		update,
		types.Extra,
	)
	assert.Zero(msg)
	user.Get(filter, extraColl, &extraData)
	assert.Equal("newname", extraData.FirstName)

	// Testing user.Auth
	// Testing that valid is true for valid auth data
	valid, userID := user.Auth(types.UserMain{Username: main.Username}, main.Password)
	assert.True(valid)
	assert.NotZero(userID)

	// Testing that valid is false for invalid password
	valid, userID = user.Auth(types.UserMain{Username: main.Username}, "")
	assert.False(valid)
	assert.Zero(userID)

	// Testing that valid is false for invalid filter
	valid, userID = user.Auth(types.UserMain{Username: "ss"}, main.Password)
	assert.False(valid)
	assert.Zero(userID)

	// Testing user.Activate
	msg = user.Activate(main.Email)
	assert.Zero(msg)
}

func TestUtils(t *testing.T) {
	assert := assert.New(t)
	uuid := generateUUID()
	assert.True(govalidator.IsUUIDv4(uuid))
}
