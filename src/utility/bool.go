package utility

import "fmt"

// GetBool : bool 로 변환
func GetBool(value interface{}) (bool, error) {
	var y bool
	switch v := value.(type) {
	case bool:
		y = bool(v)
	default:
		return y, fmt.Errorf("값을 변환할 수 없습니다")
	}
	return y, nil
}
