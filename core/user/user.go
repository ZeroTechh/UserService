package user

import (
	"github.com/ZeroTechh/UserService/core/types"

	"github.com/ZeroTechh/UserService/core/user/userExtra"
	"github.com/ZeroTechh/UserService/core/user/userMain"
	"github.com/ZeroTechh/UserService/core/user/userMeta"
)

// User handles user data
type User struct {
	main  *userMain.Main
	extra *userExtra.Extra
	meta  *userMeta.Meta
}

func (user User) init() {
	user.main = userMain.New()
	user.extra = userExtra.New()
	user.meta = userMeta.New()
}

// Add adds user to database and returns userID
func (user User) Add(main types.Main, extra types.Extra) (userID string, msg string) {
	userID = user.main.GenerateID()
	main.UserID = userID
	extra.UserID = userID

	msg = user.main.Create(main)
	if msg != "" {
		return
	}

	msg = user.extra.Create(extra)
	if msg != "" {
		return
	}

	user.meta.Create(userID)
	return
}

// Get gets the user data (either main, meta or extra) based on userID then returns it
func (user User) GetFromID(userID, dataType string) (main types.Main, extra types.Extra, meta types.Meta) {
	switch dataType {
	case "extra":
		extra = user.extra.Get(userID)
	case "meta":
		meta = user.meta.Get(userID)
	default:
		main = user.main.Get(types.Main{UserID: userID})
	}
	return
}

// GetUserID returns user's id based on email or username
func (user User) GetUserID(username, email string) string {
	filter := types.Main{Username: username, Email: email}
	return user.main.Get(filter).UserID
}
