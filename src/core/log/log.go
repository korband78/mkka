package log

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
)

// logTimeLayout : 로그 시간 레이아웃 포맷 (2006-01-02T15:04:05MST)
const logTimeLayout string = "2006-01-02MST"

// logFileNames : access - 요청별 로그, info - 기록용 로그, error - 에러
var logFileNames = []string{"debug", "access", "info", "error"}

// loggers : 로거
var loggers map[string]*log.Logger

// logLatestTime : 최근 로그 로테이트한 시간
var logLatestTime time.Time

var debug bool

// Debug : debug 출력
func Debug(i ...interface{}) {
	if debug {
		_, file, line, _ := runtime.Caller(1)
		loggers["debug"].Debug(fmt.Sprintf("%v:%v %v", file, line, fmt.Sprint(i...)))
	}
}

// Access : access 로그
func Access(i ...interface{}) {
	loggers["access"].Info(i...)
	if debug {
		loggers["debug"].Debug(i...)
	}
}

// Info : info 로그
func Info(i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	message := fmt.Sprintf("%v:%v %v", file, line, fmt.Sprint(i...))
	loggers["info"].Info(message)
	if debug {
		loggers["debug"].Debug(message)
	}
}

// Error : error 로그
func Error(i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	message := fmt.Sprintf("%v:%v %v", file, line, fmt.Sprint(i...))
	loggers["error"].Error(message)
	if debug {
		loggers["debug"].Debug(message)
	}
}

// Fatal : error 로그 & 종료
func Fatal(i ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	message := fmt.Sprintf("%v:%v %v", file, line, fmt.Sprint(i...))
	loggers["error"].Error(message)
	loggers["debug"].Debug(message)
	os.Exit(1)
}

// InitializeLogger : Loggers 초기화
func InitializeLogger(setDebug bool, logDir string, logDays int) {
	debug = setDebug
	loggers = make(map[string]*log.Logger)
	for _, logFileName := range logFileNames {
		loggers[logFileName] = log.New(logFileName)
		loggers[logFileName].SetHeader(`${time_rfc3339_nano}`)
	}
	RotateLog(logDir, logDays)
}

// RotateLog : 로그 로테이트 적용
func RotateLog(logDir string, logDays int) bool {
	now := time.Now()
	if now.Format(logTimeLayout) == logLatestTime.Format(logTimeLayout) {
		return false
	}
	logLatestTime = now
	for _, logFileName := range logFileNames {
		loggers[logFileName].SetLevel(log.DEBUG)
		if logFileName == "debug" {
			continue
		}

		// 새로운 로그 파일 생성
		os.Chdir(logDir)
		originPath := fmt.Sprintf("%v.%v", logFileName, now.Format(logTimeLayout))
		logFile, err := os.OpenFile(
			originPath,
			os.O_CREATE|os.O_RDWR|os.O_APPEND,
			os.FileMode(0666),
		)
		if err != nil {
			Error(err)
			continue
		}

		// 심볼릭 링크 연결
		symbolicPath := fmt.Sprintf("%v", logFileName)
		if _, err := os.Lstat(symbolicPath); err == nil {
			os.Remove(symbolicPath)
		}
		os.Symlink(originPath, symbolicPath)

		// 연결
		prevLogFile := loggers[logFileName].Output()
		loggers[logFileName].SetOutput(logFile)
		if prevLogFile != os.Stdout {
			if w, ok := prevLogFile.(*os.File); !ok {
				w.Close()
			}
		}
	}

	// 이전 로그 삭제
	files, err := ioutil.ReadDir(logDir)
	if err != nil {
		Error(err)
		return false
	}
	for _, file := range files {
		parsed := strings.Split(file.Name(), ".")
		if len(parsed) != 2 {
			continue
		}
		targetTime, err := time.Parse(logTimeLayout, parsed[1])
		if err != nil {
			continue
		}
		if int(time.Since(targetTime).Hours()/24) > logDays {
			err := os.Remove(fmt.Sprintf("%v/%v", logDir, file.Name()))
			if err != nil {
				Error(err)
			}
		}
	}
	return true
}
