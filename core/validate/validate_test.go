package validate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ZeroTechh/UserService/types"
	"github.com/ZeroTechh/UserService/utils"
)

type test struct {
	FirstName int
}

func TestValidators(t *testing.T) {
	assert := assert.New(t)

	// Testing that IsDataValid is valid for valid data
	mainData, extraData := utils.GetMockUserData()
	valid := IsDataValid(mainData, extraData)
	assert.True(valid)

	// Testing that IsDataValid returns false for invalid age
	extraData.BirthdayUTC = time.Now().Unix()
	valid = IsDataValid(mainData, extraData)
	assert.False(valid)

	// Testing that IsDataValid returns false for invalid data
	valid = IsDataValid(types.UserMain{}, types.UserExtra{})
	assert.False(valid)

	// Testing that IsUpdateValid returns true for valid data
	valid = IsUpdateValid(
		types.UserExtra{FirstName: "test"},
		config.Map("database").Map("collections").Str("extra"))
	assert.True(valid)
	valid = IsUpdateValid(
		types.UserMain{Username: "test123"},
		config.Map("database").Map("collections").Str("main"))
	assert.True(valid)

	// Testing that IsUpdateValid returns false if user id is not nil
	valid = IsUpdateValid(
		types.UserMain{UserID: "test123"},
		config.Map("database").Map("collections").Str("main"))
	assert.False(valid)

	// Testing that IsUpdateValid returns false if data is invalid
	valid = validateData(
		test{123}, "extra", false)
	assert.False(valid)

	collType := getDataType("")
	assert.NotZero(collType)
}
