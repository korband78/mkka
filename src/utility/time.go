package utility

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TimeToInt64 : 파라미터 시간 => int64
func TimeToInt64(t time.Time) int64 {
	timeStr := t.Format("20060102150405.999")
	parsed := strings.Split(timeStr, ".")
	if len(parsed) < 2 {
		result, _ := strconv.ParseInt(fmt.Sprintf("%v000", parsed[0]), 10, 64)
		return result
	}
	length := len(parsed[1])
	for i := 0; i < 3-length; i++ {
		parsed[1] += "0"
	}
	result, _ := strconv.ParseInt(fmt.Sprintf("%v%v", parsed[0], parsed[1]), 10, 64)
	return result
}

// NowToInt64 : 현재시간 => int64
func NowToInt64() int64 {
	return TimeToInt64(time.Now())
}

// Int64ToTime : int64 => 시간
func Int64ToTime(t int64) (time.Time, error) {
	timeStr := fmt.Sprintf("%v", t)
	if len(timeStr) >= 17 {
		return time.Parse("20060102150405.999MST", fmt.Sprintf("%v.%vKST", timeStr[0:14], timeStr[14:17]))
	}
	return time.Now(), errors.New("타임 포맷이 바르지 않습니다")
}
