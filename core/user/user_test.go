package user

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/ZeroTechh/UserService/core/utils"
)

func TestUser(t *testing.T) {
	assert := assert.New(t)

	user := New()

	// Testing Add
	main, extra := utils.GetMockUserData()
	id, msg := user.Add(main, extra)
	assert.NotZero(id)
	assert.Zero(msg)

	// Testing Auth
	assert.True(user.Auth(main.Username, "", main.Password))
	assert.False(user.Auth(main.Username, "", "INVALIDPASSWORD"))

	// Testing GetUserID
	// main, extra = utils.GetMockUserData()
	userID := user.GetUserID(main.Username, "")
	fmt.Println(userID)
	assert.Equal(id, userID)
	userID = user.GetUserID("", main.Email)
	assert.Equal(id, userID)

	// Testing Add returns invalid data for invalid main and extra data
	main, extra = utils.GetMockUserData()
	id, msg = user.Add(types.Main{}, extra)
	assert.Zero(id)
	assert.NotZero(msg)

	id, msg = user.Add(main, types.Extra{})
	assert.Zero(id)
	assert.NotZero(msg)

	// Testing Get
	returnedMain, _, _ := user.GetFromID(userID, "main")
	assert.NotZero(returnedMain)
	_, returnedExtra, _ := user.GetFromID(userID, "extra")
	assert.NotZero(returnedExtra)
	_, _, returnedMeta := user.GetFromID(userID, "meta")
	assert.NotZero(returnedMeta)

	// Testing Update
	update := types.Main{Username: "test"}
	msg = user.UpdateMain(userID, "", "", update)
	assert.Zero(msg)
	main, _, _ = user.GetFromID(userID, "main")
	assert.Equal("test", main.Username)

	update2 := types.Extra{FirstName: "test"}
	msg = user.UpdateExtra(userID, update2)
	assert.Zero(msg)
	_, extra, _ = user.GetFromID(userID, "extra")
	assert.Equal("test", extra.FirstName)
}
