package main

import (
	"fmt"
	"net"
)

func main() {
	ips,_ := Ips()
	fmt.Println(ips)
}

func Ips() (map[string]string, error) {

	ips :=  make(map[string]string)

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range interfaces {
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			return nil, err
		}
		addresses, err := byName.Addrs()
		for _, v := range addresses {
			ips[byName.Name] = v.String()
		}
	}
	return ips, nil
}
