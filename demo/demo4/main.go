package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ipToUint(ip string) uint {
	var result uint

	ips := strings.Split(ip,".")
	for i := 0; i < 4; i++ {
		oc,_ := strconv.Atoi(ips[i])
		result += uint(oc) << uint(8*(3-i))
	}
	return result
}

func ipToint(ip string) int {
	var result int
	ips := strings.Split(ip, ".")
	ip3, _ := strconv.Atoi(ips[3])
	ip2, _ := strconv.Atoi(ips[2])
	ip1, _ := strconv.Atoi(ips[1])
	ip0, _ := strconv.Atoi(ips[0])
	result += ip3
	result += ip2 << 8
	result += ip1 << 16
	result += ip0 << 24
	return ip3 | ip2 << 8 | ip1 << 16 | ip0 << 24
}

func ipToInt(ip string) uint32 {
	var result uint32
	octets := strings.Split(ip, ".")
	for i := 0; i < 4; i++ {
		octet, _ := strconv.Atoi(octets[i])
		result += uint32(octet) << uint(8*(3-i))
	}
	return result
}

func main() {
	ip := "192.168.1.1"
	fmt.Println(ipToInt(ip))
	fmt.Println(ipToint(ip))
	fmt.Println(ipToUint(ip))
}
