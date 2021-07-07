package model

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// Context : 서버 수행동안 글로벌하게 인스턴스 1개만 유지하는 데이터
type Context struct {
	From      string
	Echo      *echo.Echo
	StartTime time.Time
	Port      int
	Core      int
	Pid       int
	RootDir   string
	RouteDir  string
	PidPath   string
	CacheDir  string
	LogDir    string
	LogDays   int
	LogFiles  map[string]*os.File
	Debug     bool
}
