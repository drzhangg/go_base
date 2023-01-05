package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"strings"
)

func InetNtoA(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

func iptoint(ipaddr string) uint32 {
	var ret uint32

	ip := net.ParseIP(ipaddr)
	if ip == nil {
		return 0
	}

	if err := binary.Read(bytes.NewBuffer(ip.To4()), binary.BigEndian, &ret); err != nil {
		return 0
	}

	return ret
}

func main() {
	// 3232235776
	ip := "3.2.255.255"
	fmt.Println(iptoint(ip))

	fmt.Println(net.ParseIP(ip).To4())


	ips := strings.Split(ip,".")
	n1,_ := strconv.Atoi(ips[0])
	n2,_ := strconv.Atoi(ips[1])
	n3,_ := strconv.Atoi(ips[2])
	n4,_ := strconv.Atoi(ips[3])

	result := n1 << 24 | n2 << 16 | n3 << 8 | n4

	print(result)

	//17106697   50266407   uS
	//50266408   84082687   CN


}
