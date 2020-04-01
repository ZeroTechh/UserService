package validate

import (
	"reflect"

	"github.com/ZeroTechh/VelocityCore/logger"
	"github.com/ZeroTechh/blaze"
	"github.com/ZeroTechh/hades"
	"github.com/ZeroTechh/sentinal/v2"
	"go.uber.org/zap"
)

var (
	config = hades.GetConfig("main.yaml", []string{"config", "../config", "../../config", "../../../config"})
	log    = logger.GetLogger(
		config.Map("service").Str("validateLogFile"),
		config.Map("service").Bool("debug"),
	)
	schemaPaths = []string{"schemas", "validate/schemas", "../validate/schemas", "core/validate/schemas", "../../validate/schemas"}
)

// IsValid checks if data is valid
func IsValid(data interface{}, dataType string, isUpdate bool) bool {
	var (
		valid bool
		msg   map[string][]string
		err   error
	)
	schemaFile := dataType + ".yaml"

	funcLog := blaze.NewFuncLog(
		"UserService.Validator",
		log,
		zap.String("Data Type", dataType),
		zap.Bool("Is Update", isUpdate),
		zap.Any("Data", data),
	)
	funcLog.Started()

	if !isUpdate {
		valid, msg, err = sentinal.ValidateWithYAML(
			data,
			schemaFile,
			schemaPaths,
			customFuncs,
		)
	} else {
		valid, msg, err = sentinal.ValidateFieldsWithYAML(
			data,
			schemaFile,
			schemaPaths,
			customFuncs,
		)

		// checking if the update data is not trying to update UserID
		if reflect.ValueOf(data).FieldByName("UserID").String() != "" {
			valid = false
			msg = map[string][]string{
				"UserID": []string{"Tring to update UserID"},
			}
		}
	}

	if err != nil {
		funcLog.Error(err)
		valid = false
	}

	funcLog.Completed(
		zap.Bool("Valid", valid),
		zap.Any("Message", msg),
	)

	return valid
}
