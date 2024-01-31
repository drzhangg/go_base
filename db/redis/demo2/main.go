package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

type IpRange struct {
	StartIp uint32
	EndIp   uint32
	Country string
}

func IpToUint(ip string) uint32 {
	parts := strings.Split(ip, ".")
	b1, _ := strconv.Atoi(parts[0])
	b2, _ := strconv.Atoi(parts[1])
	b3, _ := strconv.Atoi(parts[2])
	b4, _ := strconv.Atoi(parts[3])
	return uint32(b1<<24 | b2<<16 | b3<<8 | b4)
}

func LoadToRedis(client *redis.Client, ipRanges []IpRange) error {
	for _, r := range ipRanges {
		key := "ip_ranges"
		score := int(r.EndIp)
		err := client.ZAdd(context.TODO(), key, redis.Z{
			Score:  float64(score),
			Member: r.Country,
		}).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func FindCountryByIp(client *redis.Client, ip string) (string, error) {
	key := "ip_ranges"
	score := string(IpToUint(ip))
	result, err := client.ZRangeByScore(context.TODO(), key, &redis.ZRangeBy{
		Max: score,
	}).Result()
	if err != nil {
		return "", err
	}

	if len(result) == 0 {
		return "", fmt.Errorf("no result found")
	}
	return result[0], nil
}

func LoadIpRanges(filename string) ([]IpRange, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ranges := make([]IpRange, 0)
	for {
		var line string
		_, err := fmt.Scanln(file, &line)
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

func ParseIPRange(line string) (IpRange, error) {
	parts := strings.Split(line, ".")
	if len(parts) != 3 {
		return IpRange{}, fmt.Errorf("Invalid IP range line: %s", line)
	}

	startIp, err := strconv.ParseUint(parts[0], 0, 32)
	//startIp := net.ParseIP(parts[0])
	if err != nil {
		return IpRange{}, fmt.Errorf("Invalid start IP address: %s", parts[0])
	}

	endIp, err := strconv.ParseUint(parts[1], 0, 32)
	//endIp := net.ParseIP(parts[1])
	if err != nil {
		return IpRange{}, fmt.Errorf("Invalid end IP address: %s", parts[1])
	}

	return IpRange{
		StartIp: uint32(startIp),
		EndIp:   uint32(endIp),
		Country: parts[2],
	}, nil
}


func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "150.158.87.137:6379",
		Password: "",
		DB:       0,
	})

	logrus.Infof("")

	ctx := context.TODO()

	cmd := rdb.Set(ctx, "name", "jerry", time.Second*10)
	fmt.Println(cmd.String())
}
