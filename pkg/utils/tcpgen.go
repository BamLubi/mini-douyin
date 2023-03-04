package utils

import (
	"net"
	"strings"
)

func GetFreePort() (uint64, error) {
	// 解析地址
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, nil
	}
	// 利用 ListenTCP 方法的如下特性
	// 如果 addr 的端口字段为0，函数将选择一个当前可用的端口
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, nil
	}
	// 关闭资源
	defer listen.Close()
	// 为了拿到具体的端口值，我们转换成 *net.TCPAddr类型获取其Port
	port := listen.Addr().(*net.TCPAddr).Port
	return uint64(port), nil
}

func GetIp() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]
	return ip, nil
}