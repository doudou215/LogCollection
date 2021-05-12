package utils

import (
	"fmt"
	"net"
	"strings"
)

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("GetOutBoundIP error ", err)
		return "", err
	}

	defer conn.Close()
	// 强制类型转换
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Printf("ip %s\n", localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return ip, err
}

func main() {
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println("main: GetOutBoundIP ", err)
	}
	fmt.Println(ip)
}
