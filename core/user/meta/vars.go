package meta

import (
	"github.com/ZeroTechh/VelocityCore/logger"
	"github.com/ZeroTechh/hades"
)

var (
	// all the configs
	config          = hades.GetConfig("main.yaml", []string{"config", "../../../config"})
	dbConfig        = config.Map("database")
	metaCollection  = dbConfig.Map("collections").Str("meta")
	debugMode       = config.Map("service").Bool("debug")
	accountStatuses = config.Map("accountStatuses")

	log = logger.GetLogger(
		config.Map("service").Str("lowLevelLogFile"),
		config.Map("service").Bool("debug"),
	)
)
