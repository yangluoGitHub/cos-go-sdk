package coscloud

import (
	"bytes"
	"crypto/x509"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

// Config cos configure
type Config struct {
	Endpoint           string         // cos地址
	Appid              uint           // appid
	SecretId           string         // SecretId
	SecretKey          string         // SecretKey
	SignExpiredSeconds uint           // 多次有效签名过期秒数
	RetryTimes         uint           // 失败重试次数，默认3
	UserAgent          string         // SDK名称/版本/系统信息
	DefaultSliceSize   int            // 默认分片大小，3MB
	Timeout            time.Duration  // 超时时间，默认60s
	Pool               *x509.CertPool // https CA证书 默认为nil
}

// 获取默认配置
func getDefaultCosConfig() *Config {
	config := Config{}

	config.Endpoint = "https://web.file.myqcloud.com/files/v1/"
	config.Appid = 0
	config.SecretId = ""
	config.SecretKey = ""
	config.SignExpiredSeconds = 200
	config.RetryTimes = 3
	config.UserAgent = userAgent
	config.DefaultSliceSize = 3 * 1024 * 1024
	config.Timeout = time.Second * 60 // seconds
	config.Pool = nil

	return &config
}

// Get User Agent
// Go sdk相关信息，包括sdk版本，操作系统类型，GO版本
var userAgent = func() string {
	sys := getSysInfo()
	return fmt.Sprintf("cos-go-sdk/%s (%s/%s/%s;%s)", SDK_VERSION, sys.name,
		sys.release, sys.machine, runtime.Version())
}()

type sysInfo struct {
	name    string // 操作系统名称windows/Linux
	release string // 操作系统版本 2.6.32-220.23.2.ali1089.el5.x86_64等
	machine string // 机器类型amd64/x86_64
}

// Get　system info
// 获取操作系统信息、机器类型
func getSysInfo() sysInfo {
	name := runtime.GOOS
	release := "-"
	machine := runtime.GOARCH
	if out, err := exec.Command("uname", "-s").CombinedOutput(); err == nil {
		name = string(bytes.TrimSpace(out))
	}
	if out, err := exec.Command("uname", "-r").CombinedOutput(); err == nil {
		release = string(bytes.TrimSpace(out))
	}
	if out, err := exec.Command("uname", "-m").CombinedOutput(); err == nil {
		machine = string(bytes.TrimSpace(out))
	}
	return sysInfo{name: name, release: release, machine: machine}
}
