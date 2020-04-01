package userExtra

import (
	"testing"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/utils"
)

func TestMeta(t *testing.T) {
	assert := assert.New(t)

	extra := NewExtra()

	// Testing that extra can add and get user extra data
	_, data := utils.GetMockUserData()
	msg := extra.Create(data)
	assert.Zero(msg)
	assert.Equal(data, extra.Get(data.UserID))

	// Testing that it returns invalid data message for invalid data
	msg = extra.Create(types.Extra{})
	assert.NotZero(msg)

	// Testing that extra cann update user extra data
	name := "name"
	update := types.Extra{FirstName: name}
	msg = extra.Update(data.UserID, update)
	assert.Zero(msg)
	assert.Equal(name, extra.Get(data.UserID).FirstName)

	// Testing that it returns invalid data message for invalid data
	msg = extra.Update(data.UserID, types.Extra{FirstName: "CC"}) // should fail as there are capital letters
	assert.NotZero(msg)
}
