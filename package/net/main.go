package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. 解析ipv4地址
	ipStr := "111.229.107.231"
	ip := net.ParseIP(ipStr)
	if ip == nil {
		fmt.Fprintf(os.Stderr, "Err:无效的地址\n")
		return
	}

	// 2. ip默认掩码
	defaultMask := ip.DefaultMask()
	fmt.Printf("ip: %s \n", ip.String())
	fmt.Printf("defaultMask: %s \n", defaultMask.String())
	fmt.Fprintf(os.Stdout, "network: %s \n", ip.Mask(defaultMask)) // 192.168.1.0
	// 3. 解析ipv6地址
	fmt.Println(net.ParseIP("2001:db8::68"))

	// 4. 将域名解析Ip地址
	domain := "www.baidu.com"
	ipAddr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s \n", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "%s IP: %s Network: %s Zone: %s\n", ipAddr.String(), ipAddr.IP, ipAddr.Network(), ipAddr.Zone)

	// 5. 获得域名对应的所有Ip地址
	ns, err := net.LookupHost(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}
	for _, v := range ns {
		fmt.Printf("ip: %s \n", v)
	}

	// 6. 查看主机服务器（service）占用的端口,这些服务，都是tcp或者udp的
	port, err := net.LookupPort("tcp", "telnet")
	if err != nil {
		fmt.Fprintf(os.Stderr, "未找到指定服务")
		return
	}
	fmt.Fprintf(os.Stdout, "telnet port: %d \n", port)

	// 7. 将一个host地址转换为TCPAddr。host=ip:port
	tcpAddr, err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "www.baidu.com:80 IP: %s PORT: %d", tcpAddr.IP.String(), tcpAddr.Port)
}
