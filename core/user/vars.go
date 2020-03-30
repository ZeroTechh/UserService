package user

import (
	"github.com/ZeroTechh/VelocityCore/logger"
	"github.com/ZeroTechh/hades"
)

var (
	// all the configs
	config          = hades.GetConfig("main.yaml", []string{"config", "../config", "../../config"})
	dbConfig        = config.Map("database")
	extraDataConfig = config.Map("userExtraData")

	log = logger.GetLogger(
		config.Map("service").Str("lowLevelLogFile"),
		config.Map("service").Bool("debug"),
	)

	// all the collections of db
	collections = dbConfig.Map("collections")
	mainColl    = collections.Str("main")
	extraColl   = collections.Str("extra")
	metaColl    = collections.Str("meta")

	accountStatuses = extraDataConfig.Map("accountStatuses")
	messages        = config.Map("messages")
	debugMode       = config.Map("service").Bool("debug")
)
