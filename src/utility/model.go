package utility

import (
	"fmt"
	"reflect"
	"regexp"
	"src/core/log"
	"unicode/utf8"

	"go.mongodb.org/mongo-driver/bson"
)

// INFINITY : 무한대
var INFINITY int64 = 99999999999999999

func _validate(row bson.M, value interface{}) (returnErr error) {
	defer func() {
		if v := recover(); v != nil {
			msg := fmt.Sprintf("오류가 발생했습니다(%v)", v)
			log.Error(msg)
			returnErr = fmt.Errorf(msg)
		}
	}()

	_name := row["name"]
	if _name == nil {
		return returnErr
	}
	name := fmt.Sprintf("%v", _name)

	valueStr := fmt.Sprintf("%v", value)
	_enableBlank := row["enable_blank"]
	if _enableBlank != nil {
		enableBlank, err := GetBool(_enableBlank)
		if err == nil && enableBlank && valueStr == "" {
			return nil
		}
	}

	_minLength := row["min_length"]
	if _minLength != nil {
		minLength, err := GetInt(_minLength)
		if err == nil && utf8.RuneCountInString(valueStr) < minLength {
			returnErr = fmt.Errorf("%v: 길이가 %v이상이어야 합니다", name, minLength)
			return returnErr
		}
	}

	_maxLength := row["max_length"]
	if _maxLength != nil {
		maxLength, err := GetInt(_maxLength)
		if err == nil && utf8.RuneCountInString(valueStr) > maxLength {
			returnErr = fmt.Errorf("%v: 길이가 %v이하여야 합니다", name, maxLength)
			return returnErr
		}
	}

	valueInt64, err := GetInt64(value)
	if err == nil {
		_minValue := row["min_value"]
		if _minValue != nil {
			minValue, err := GetInt64(_minValue)
			if err == nil && valueInt64 < minValue {
				returnErr = fmt.Errorf("%v: 값이 %v이상이어야 합니다", name, minValue)
				return returnErr
			}
		}

		_maxValue := row["max_value"]
		if _maxValue != nil {
			maxValue, err := GetInt64(_maxValue)
			if err == nil && valueInt64 > maxValue {
				returnErr = fmt.Errorf("%v: 값이 %v이하여야 합니다", name, maxValue)
				return returnErr
			}
		}
	}

	_regexStr := row["regex"]
	if _regexStr != nil {
		regexStr := row["regex"].(string)
		regex := regexp.MustCompile(regexStr)
		if !regex.MatchString(valueStr) {
			returnErr = fmt.Errorf("%v: 유효하지 않는 형식입니다", name)
			return returnErr
		}
	}

	_inArray := row["in"]
	if _inArray != nil {
		isMatch := false
		for _, item := range _inArray.([]interface{}) {
			if fmt.Sprintf("%v", item) == valueStr {
				isMatch = true
				break
			}
		}
		if !isMatch {
			returnErr = fmt.Errorf("%v: 허용되지 않은 값입니다", name)
			return returnErr
		}
	}

	return nil
}

// Validate : 유효성 검사
func Validate(validateObj bson.M, field string, pValue interface{}) (returnErr error) {
	defer func() {
		if v := recover(); v != nil {
			msg := fmt.Sprintf("오류가 발생했습니다(%v)", v)
			log.Error(msg)
			returnErr = fmt.Errorf(msg)
		}
	}()

	returnErr = nil

	_row, ok := validateObj[field]
	if _row == nil || !ok {
		return returnErr
	}
	row := _row.(bson.M)

	switch reflect.TypeOf(pValue).Kind() {
	case reflect.Slice:
		values := reflect.ValueOf(pValue)
		for i := 0; i < values.Len(); i++ {
			returnErr = _validate(row, values.Index(i))
			if returnErr != nil {
				return returnErr
			}
		}
	case reflect.Array:
		values := reflect.ValueOf(pValue)
		for i := 0; i < values.Len(); i++ {
			returnErr = _validate(row, values.Index(i))
			if returnErr != nil {
				return returnErr
			}
		}
	default:
		returnErr = _validate(row, pValue)
	}

	return returnErr
}
