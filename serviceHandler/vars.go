package serviceHandler

import (
	"github.com/ZeroTechh/VelocityCore/logger"
	"github.com/ZeroTechh/hades"
)

var (
	config = hades.GetConfig(
		"main.yaml", []string{"config", "../config"})
	log = logger.GetLogger(
		config.Map("service").Str("lowLevelLogFile"),
		config.Map("service").Bool("debug"),
	)
	collections = config.Map("database").Map("collections")
	mainColl    = collections.Str("main")
	extraColl   = collections.Str("extra")
)
