package utility

import "fmt"

// GetFloat64 : float64 로 변환
func GetFloat64(value interface{}) (float64, error) {
	var y float64
	switch v := value.(type) {
	case int:
		y = float64(v)
	case int8:
		y = float64(v)
	case int16:
		y = float64(v)
	case int32:
		y = float64(v)
	case int64:
		y = float64(v)
	case uint:
		y = float64(v)
	case uint8:
		y = float64(v)
	case uint16:
		y = float64(v)
	case uint32:
		y = float64(v)
	case uint64:
		y = float64(v)
	case float32:
		y = float64(v)
	case float64:
		y = v
	default:
		return y, fmt.Errorf("값을 변환할 수 없습니다")
	}
	return y, nil
}

// GetInt64 : int64 로 변환
func GetInt64(value interface{}) (int64, error) {
	var y int64
	switch v := value.(type) {
	case int:
		y = int64(v)
	case int8:
		y = int64(v)
	case int16:
		y = int64(v)
	case int32:
		y = int64(v)
	case int64:
		y = v
	case uint:
		y = int64(v)
	case uint8:
		y = int64(v)
	case uint16:
		y = int64(v)
	case uint32:
		y = int64(v)
	case uint64:
		y = int64(v)
	case float32:
		y = int64(v)
	case float64:
		y = int64(v)
	default:
		return y, fmt.Errorf("값을 변환할 수 없습니다")
	}
	return y, nil
}

// GetInt : int 로 변환
func GetInt(value interface{}) (int, error) {
	var y int
	switch v := value.(type) {
	case int:
		y = v
	case int8:
		y = int(v)
	case int16:
		y = int(v)
	case int32:
		y = int(v)
	case int64:
		y = int(v)
	case uint:
		y = int(v)
	case uint8:
		y = int(v)
	case uint16:
		y = int(v)
	case uint32:
		y = int(v)
	case uint64:
		y = int(v)
	case float32:
		y = int(v)
	case float64:
		y = int(v)
	default:
		return y, fmt.Errorf("값을 변환할 수 없습니다")
	}
	return y, nil
}
