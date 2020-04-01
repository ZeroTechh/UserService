package userMain

import (
	"testing"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/utils"
)

func TestMain(t *testing.T) {
	assert := assert.New(t)

	Main := NewMain()

	// Testing that Main can add and get user Main data
	data, _ := utils.GetMockUserData()
	filter := types.Main{UserID: data.UserID}
	msg := Main.Create(data)
	assert.Zero(msg)
	assert.Equal(data, Main.Get(filter))

	// Testing that it returns invalid data message for invalid data
	msg = Main.Create(types.Main{})
	assert.NotZero(msg)

	// Testing that Main cann update user Main data
	data, _ = utils.GetMockUserData() // generating new data to get new username
	name := data.Username
	update := types.Main{Username: name}
	msg = Main.Update(filter, update)
	assert.Zero(msg)
	assert.Equal(name, Main.Get(filter).Username)

	// Testing that it returns invalid data message for invalid data
	msg = Main.Update(filter, types.Main{Username: "CC"}) // should fail as there are capital letters
	assert.NotZero(msg)
}
