package utility

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

// IsRunningProcess : 프로세스 실행 중 여부
func IsRunningProcess(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	if err == nil {
		return true
	}
	return false
}

// IsRunningPort : 포트 사용 여부
func IsRunningPort(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return true
	}
	defer l.Close()
	return true
}
