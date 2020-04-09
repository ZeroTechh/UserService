package user

import (
	"github.com/ZeroTechh/UserService/core/types"

	"github.com/ZeroTechh/UserService/core/user/userExtra"
	"github.com/ZeroTechh/UserService/core/user/userMain"
	"github.com/ZeroTechh/UserService/core/user/userMeta"
)

// New returns new user data manager
func New() *User {
	user := User{}
	user.init()
	return &user
}

// User handles user data
type User struct {
	main  *userMain.Main
	extra *userExtra.Extra
	meta  *userMeta.Meta
}

func (user *User) init() {
	user.main = userMain.New()
	user.extra = userExtra.New()
	user.meta = userMeta.New()
}

// Add adds user to database and returns userID
func (user User) Add(main types.Main, extra types.Extra) (string, string) {
	userID := user.main.GenerateID()
	main.UserID = userID
	extra.UserID = userID

	msg := user.main.Create(main)
	if msg != "" {
		return "", msg
	}

	msg = user.extra.Create(extra)
	if msg != "" {
		return "", msg
	}

	user.meta.Create(userID)
	return userID, ""
}

// GetFromID gets the user data (either main, meta or extra) based on userID then returns it
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

// UpdateMain updates user's main data
func (user User) UpdateMain(userID, username, email string, update types.Main) string {
	filter := types.Main{UserID: userID, Username: username, Email: email}
	return user.main.Update(filter, update)
}

// UpdateExtra updates user's extra data
func (user User) UpdateExtra(userID string, update types.Extra) string {
	return user.extra.Update(userID, update)
}

// Auth is used to authenticate username, email or password
func (user User) Auth(username, email, password string) bool {
	// TODO add password hashing
	filter := types.Main{Username: username, Email: email}
	data := user.main.Get(filter)
	return data.Password == password && data != types.Main{}
}
