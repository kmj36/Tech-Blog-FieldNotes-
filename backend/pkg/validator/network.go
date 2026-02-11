package validator

import (
	"fmt"
	"net"
	"strconv"
)

// host:port 문자열이 유효한 IP 주소인지 검사하는 함수
func ValidateAddr(addr string) error {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}

	// IP 검사
	if ip := net.ParseIP(host); ip == nil {
		return fmt.Errorf("invalid ip:%s", host)
	}

	// Port 검사
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid port:%s", portStr)
	}

	return nil
}