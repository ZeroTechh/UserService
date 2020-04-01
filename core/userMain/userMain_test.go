package userMain

import (
	"testing"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/utils"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)

	main := NewMain()

	// Testing that Main can add and get user Main data
	data, _ := utils.GetMockUserData()
	filter := types.Main{UserID: data.UserID}
	msg := main.Create(data)
	assert.Zero(msg)
	assert.Equal(data, main.Get(filter))

	// Testing that it returns invalid data message for invalid data
	msg = main.Create(types.Main{})
	assert.NotZero(msg)

	// Testing that Main can update user Main data
	data, _ = utils.GetMockUserData() // generating new data to get new username
	name := data.Username
	update := types.Main{Username: name}
	msg = main.Update(filter, update)
	assert.Zero(msg)
	assert.Equal(name, main.Get(filter).Username)

	// Testing that it returns invalid data message for invalid data
	msg = main.Update(filter, types.Main{Username: "CC"}) // should fail as there are capital letters
	assert.NotZero(msg)

	// Testing Auth
	assert.True(main.Auth(data.Username, "", data.Password))
	assert.False(main.Auth("invalidUsername", "", data.Password))
}
