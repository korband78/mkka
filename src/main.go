package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"src/core/log"
	"src/core/middleware"
	"src/model"
	"src/route"
	"src/utility"
	"strconv"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	fmt.Println("################################################")
	fmt.Printf("%v server starting...\n", time.Now().Format(time.RFC3339Nano))

	// 서버 초기화
	var ctx model.Context
	initializeServer(&ctx)

	/*
		크론탭 : Batch Job
	*/
	crontab := cron.New()
	// 일 단위
	crontab.AddFunc("1 0 * * *", func() {
		// 로그 로테이트, 즉시
		if log.RotateLog(ctx.LogDir, ctx.LogDays) {
			log.Info("rotate log.")
		}
	})

	crontab.Start()
	defer crontab.Stop()
	// 라우팅
	route.Routing(ctx)

	// 서버 정보 출력
	fmt.Printf("# ROOT PATH: %v\n# DEBUG: %v\n# PORT: %v\n# PID: %v\n# CORE: %v\n# LOG SAVE DAYS: %v\n", ctx.RootDir, ctx.Debug, ctx.Port, ctx.Pid, ctx.Core, ctx.LogDays)
	fmt.Printf("%v server startup complete. (%.2fms)\n", time.Now().Format(time.RFC3339Nano), time.Since(ctx.StartTime).Seconds()*1000)
	fmt.Println("################################################")

	// 서버 시작
	if ctx.Port == 443 {
		go func(c *echo.Echo) {
			log.Fatal(c.Start("0.0.0.0:80"))
		}(ctx.Echo)
		log.Fatal(ctx.Echo.StartAutoTLS(fmt.Sprintf("0.0.0.0:%v", ctx.Port)))
	} else {
		log.Fatal(ctx.Echo.Start(fmt.Sprintf("0.0.0.0:%v", ctx.Port)))
	}
}

// 서버 초기화
func initializeServer(ctx *model.Context) {
	/*
		context 초기화
	*/
	// 서버 시작 시간
	ctx.StartTime = time.Now()

	// 멀티 코어 설정
	runtime.GOMAXPROCS(runtime.NumCPU())
	ctx.Core = runtime.GOMAXPROCS(0)

	// 포트
	flag.IntVar(&ctx.Port, "port", 8080, "listen port")

	// 루트 경로
	rdir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	rdir = filepath.Dir(rdir)
	flag.StringVar(&ctx.RootDir, "rdir", rdir, "root dir")

	// 로그 저장 기간
	flag.IntVar(&ctx.LogDays, "ldays", 5, "log save days")

	/*
		디버그 모드 여부
		log.Debug는 디버그 모드 시애만 Stdout으로 출력
		디버그 모드 시 모든 로그는 Stdout로 출력
	*/
	flag.BoolVar(&ctx.Debug, "debug", false, "debug mode on")

	// Flag 파싱
	flag.Parse()

	// 라우팅 경로
	ctx.RouteDir = fmt.Sprintf("%v/src/route", ctx.RootDir)

	// 로그 경로
	ctx.LogDir = fmt.Sprintf("%v/var/log", ctx.RootDir)

	// 캐시 경로
	ctx.CacheDir = fmt.Sprintf("%v/var/.cache", ctx.RootDir)

	// 이메일
	ctx.From = "byssy@naver.com"

	var err error

	// 디렉토리 생성
	for _, targetDir := range []string{ctx.RouteDir, ctx.LogDir, ctx.CacheDir} {
		err = os.MkdirAll(targetDir, os.ModePerm)
		if err != nil {
			log.Fatal(fmt.Sprintf("cannot mkdir %v. (%v)", targetDir, err))
		}
	}

	// 로거 초기화
	log.InitializeLogger(ctx.Debug, ctx.LogDir, ctx.LogDays)

	// PID
	ctx.Pid = os.Getpid()
	// PID 경로
	ctx.PidPath = fmt.Sprintf("%v/var/pid", ctx.RootDir)
	if _, err := os.Stat(ctx.PidPath); err == nil {
		if pidBytes, err := ioutil.ReadFile(ctx.PidPath); err == nil {
			if pid, err := strconv.Atoi(string(pidBytes)); err == nil && utility.IsRunningProcess(pid) {
				log.Info(fmt.Sprintf("there is already a process running. (previous pid:%v, current pid:%v)", pid, ctx.Pid))
			}
		}
	}
	ioutil.WriteFile(ctx.PidPath, []byte(fmt.Sprintf("%v", ctx.Pid)), os.FileMode(0666))

	// Echo 초기화
	ctx.Echo = echo.New()

	/*
		Middleware 설정
	*/
	// AUTO TLS
	if ctx.Port == 443 {
		ctx.Echo.AutoTLSManager.Cache = autocert.DirCache(ctx.CacheDir)
		ctx.Echo.Pre(echoMiddleware.HTTPSRedirect())
	}

	// 요청별 로그 포맷
	ctx.Echo.Use(middleware.RequestLog)

	// CORS
	ctx.Echo.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		ExposeHeaders: []string{"*"},
	}))

	// / 추가
	ctx.Echo.Pre(echoMiddleware.AddTrailingSlashWithConfig(echoMiddleware.TrailingSlashConfig{
		Skipper: func(c echo.Context) bool {
			// 프로메테우스 URI만 스킵
			if c.Request().RequestURI == "/metrics" {
				return true
			}
			return false
		},
	}))

	// 에러 출력
	ctx.Echo.Use(echoMiddleware.Recover())

	// 프로메테우스 Exporter
	prometheus.NewPrometheus("server", nil).Use(ctx.Echo)

}
