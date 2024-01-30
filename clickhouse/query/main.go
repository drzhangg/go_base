package main

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"log"
	"time"
)

func main() {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"xxx:8123"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout:      time.Second * 30,
		MaxOpenConns:     5,
		MaxIdleConns:     5,
		ConnMaxLifetime:  time.Duration(10) * time.Minute,
		ConnOpenStrategy: clickhouse.ConnOpenInOrder,
		BlockBufferSize: 10,
		MaxCompressionBuffer: 10240,
		Protocol: clickhouse.HTTP,
	})
	if err != nil {
		log.Println("conn err:",err)
	}
	err = conn.Ping(context.Background())

	if err != nil {
		log.Println("Ping err:",err)
	}

}
