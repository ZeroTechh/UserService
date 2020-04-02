package serviceHandler

import (
	"github.com/ZeroTechh/UserService/core/userExtra"
	"github.com/ZeroTechh/UserService/core/userMain"
	"github.com/ZeroTechh/UserService/core/userMeta"
	"github.com/ZeroTechh/VelocityCore/logger"
	"github.com/ZeroTechh/hades"
)

var (
	config = hades.GetConfig("main.yaml", []string{"config", "../config"})
	log    = logger.GetLogger(
		config.Map("service").Str("logFile"),
		config.Map("service").Bool("debug"),
	)
)

// Handler is used to handle all user service functions
type Handler struct {
	main  userMain.Main
	extra userExtra.Extra
	meta  userExtra.Extra
}

// Init is used to initialize
func (handler *Handler) Init() {
	handler.main = userMain.New()
	handler.extra = userExtra.New()
	handler.meta = userMeta.New()
}
