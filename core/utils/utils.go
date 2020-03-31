package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/ZeroTechh/UserService/core/types"
)

func generateRandomString(length int) string {
	charset := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GetMockUserData returns mock user data for testing
func GetMockUserData() (types.Main, types.Extra) {
	mockUsername := strings.ToLower(generateRandomString(10))
	mockPassword := generateRandomString(10)
	mockUserData := types.Main{
		Username: mockUsername,
		Password: mockPassword,
		Email:    mockUsername + "@gmail.com",
	}

	mockUserExtraData := types.Extra{
		FirstName:   mockUsername,
		LastName:    mockUsername,
		Gender:      "male",
		BirthdayUTC: int64(864466669),
	}

	return mockUserData, mockUserExtraData
}
