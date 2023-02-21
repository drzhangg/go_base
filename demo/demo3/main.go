package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type IPRange struct {
	Start net.IP
	End   net.IP
	Name  string
}

func ParseIPRange(line string) (*IPRange, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Invalid IP range line: %s", line)
	}

	startIP := net.ParseIP(parts[0])
	if startIP == nil {
		return nil, fmt.Errorf("Invalid start IP address: %s", parts[0])
	}

	endIP := net.ParseIP(parts[1])
	if endIP == nil {
		return nil, fmt.Errorf("Invalid end IP address: %s", parts[1])
	}

	return &IPRange{
		Start: startIP,
		End:   endIP,
		Name:  parts[2],
	}, nil
}

func FindCountryByIP(ip string, ipRanges []*IPRange) (string, error) {
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return "", fmt.Errorf("Invalid IP address: %s", ip)
	}

	for _, r := range ipRanges {
		if bytesCompare(ipAddr, r.Start) >= 0 && bytesCompare(ipAddr, r.End) <= 0 {
			return r.Name, nil
		}
	}

	return "", fmt.Errorf("No matching IP range found for IP address: %s", ip)
}

func bytesCompare(a, b net.IP) int {
	return bytesToUint32(a.To4()) - bytesToUint32(b.To4())
}

func bytesToUint32(b []byte) int {
	return int(b[3]) | int(b[2])<<8 | int(b[1])<<16 | int(b[0])<<24
}

func LoadIPRangesFile(path string) ([]*IPRange, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ranges := make([]*IPRange, 0)

	for {
		var line string
		_, err := fmt.Fscanln(file, &line)
		if err != nil {
			break
		}

		r, err := ParseIPRange(line)
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, r)
	}

	return ranges, nil
}

func main() {
	ipRanges, err := LoadIPRangesFile("ip_ranges.txt")
	if err != nil {
		panic(err)
	}

	country, err := FindCountryByIP("2.255.255.1", ipRanges)
	if err != nil {
		panic(err)
	}

	fmt.Println("Country: ", country) // CN
}
