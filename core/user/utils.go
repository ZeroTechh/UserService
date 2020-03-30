package user

import (
	"time"

	"github.com/google/uuid"

	"github.com/ZeroTechh/UserService/core/types"
)

// generates basic user meta data
func generateMeta() types.UserMeta {
	status := accountStatuses.Str("unverified")
	// automatically mark user as verified if debug is true
	if debugMode {
		status = accountStatuses.Str("verified")
	}
	return types.UserMeta{
		AccountStatus:      status,
		AccountCreationUTC: time.Now().Unix(),
	}
}

// Generates A Version 4 UUID
func generateUUID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic("Error While Generating UUID")
	}
	return id.String()
}
