package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/core/types"
	"github.com/ZeroTechh/UserService/core/utils"
)

func TestValidators(t *testing.T) {
	assert := assert.New(t)

	// Checking if validator returns true for valid data
	main, extra := utils.GetMockUserData()
	assert.True(IsValid(main, types.Main, false))
	assert.True(IsValid(extra, types.Extra, false))

	// Checking if validator returns false for invalid data
	extra.FirstName = "Test test"
	assert.False(IsValid(extra, types.Extra, false))
}
