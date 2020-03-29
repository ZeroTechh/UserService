package validate

import (
	"reflect"
	"strconv"
	"time"
)

var customFuncs = map[string]func(reflect.Value, string) (bool, string, error){
	"isGender": isGender,
	"age":      age,
}

// Checks if the gender is valid
func isGender(value reflect.Value, _ string) (bool, string, error) {
	validGenders := config.Map("userExtraData").Map("genders").Data
	usersGender := value.String()
	for _, validGender := range validGenders {
		if validGender == usersGender {
			return true, "", nil
		}
	}
	return false, "Invalid gender", nil
}

// checks if age is greater than required age
func age(value reflect.Value, ageRequiredStr string) (valid bool, msg string, err error) {
	usersAgeStamp := time.Unix(value.Int(), 0)
	ageRequired, _ := strconv.ParseInt(ageRequiredStr, 10, 64)
	valid = int64(time.Now().Year()-usersAgeStamp.Year()) >= ageRequired
	if !valid {
		msg = "Age is less than required"
	}
	return
}
