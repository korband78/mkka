package middleware

import (
	"fmt"
	"src/core/log"
	"src/utility"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// RequestLog : 요청별 로깅
func RequestLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		start := time.Now()
		// reqID 부여
		reqID := utility.UUID(true)
		res.Header().Set("X-Request-Id", reqID)
		if err = next(c); err != nil {
			c.Error(err)
		}

		// 로깅 정보
		latency := time.Since(start).Seconds() * 1000
		ip := c.RealIP()
		host := req.Host
		uri := req.RequestURI
		method := req.Method
		reqSize := req.Header.Get(echo.HeaderContentLength)
		if reqSize == "" {
			reqSize = "0"
		}
		resSize := strconv.FormatInt(res.Size, 10)
		status := res.Status

		// 시간, 응답시간(ms) 메소드, 요청URL, 요청 크기(bytes), 응답 크기(bytes), 요청 IP
		log.Access(fmt.Sprintf("%v %v %v%v %.2f %v %v %v %v", status, method, host, uri, latency, reqSize, resSize, ip, reqID))
		return
	}
}
