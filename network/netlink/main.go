package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {
	list,err := netlink.LinkList()
	if err != nil {
		fmt.Println("err1:",err)
		return
	}


	for _,link := range list{
		if link.Type() != "device"{
			continue
		}

		addr,err := netlink.AddrList(link,netlink.FAMILY_V4)
		if err != nil {
			fmt.Println("err2:",err)
			continue
		}
		fmt.Println("addr:",addr)
	}
}
