package userMain

import (
	"testing"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/utils"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)

	main := New()

	// Testing Create and Get function
	data, _ := utils.GetMockUserData()
	filter := types.Main{UserID: data.UserID}
	msg := main.Create(data)
	assert.Zero(msg)
	assert.Equal(data, main.Get(filter))

	// Testing Create returns invalid data message
	msg = main.Create(types.Main{})
	assert.NotZero(msg)

	// Testing that add user returns username exists message
	data2, _ := utils.GetMockUserData()
	data2.Username = data.Username
	msg = main.Create(data2)
	assert.Equal(messages.Str("usernameExists"), msg)

	// Testing that add user returns email exists message
	data2, _ = utils.GetMockUserData()
	data2.Email = data.Email
	msg = main.Create(data2)
	assert.Equal(messages.Str("emailExists"), msg)

	// Testing update function
	data2, _ = utils.GetMockUserData()
	msg = main.Update(
		types.Main{UserID: data.UserID},
		types.Main{Username: data.Username},
	)
	assert.Zero(msg)

	// Testing update returns invalid data
	msg = main.Update(
		types.Main{UserID: data.UserID},
		types.Main{Username: "INVALIDUSERNAME"},
	)
	assert.NotZero(msg)

	// Testing generation of id
	id := main.GenerateID()
	assert.NotZero(id)
}
