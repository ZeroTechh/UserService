package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/utils"
)

func TestMeta(t *testing.T) {
	assert := assert.New(t)

	meta := NewMeta()

	// Testing that meta can create, add and get user meta data
	data, _ := utils.GetMockUserData()
	meta.Create(data.UserID)
	assert.NotZero(meta.Get(data.UserID))

	// Testing that meta can activate user's data
	meta.Activate(data.UserID)
	metaData := meta.Get(data.UserID)
	assert.Equal(accountStatuses.Str("verified"), metaData.AccountStatus)
}
