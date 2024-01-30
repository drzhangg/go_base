package main

import (
	"context"
	"database/sql"

	"github.com/mailru/go-clickhouse/v2"
	"log"
	"time"
)

func main() {
	connect, err := sql.Open("chhttp", "http://xxx:8123/default")
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS example (
			country_code FixedString(2),
			os_id        UInt8,
			browser_id   UInt8,
			categories   Array(Int16),
			action_day   Date,
			action_time  DateTime
		) engine=Memory
	`)

	if err != nil {
		log.Fatal(err)
	}

	tx, err := connect.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(`
		INSERT INTO example (
			country_code,
			os_id,
			browser_id,
			categories,
			action_day,
			action_time
		) VALUES (
			?, ?, ?, ?, ?, ?
		)`)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		if _, err := stmt.Exec(
			"RU",
			10+i,
			100+i,
			clickhouse.Array([]int16{1, 2, 3}),
			clickhouse.Date(time.Now()),
			time.Now(),
		); err != nil {
			log.Fatal(err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	rows, err := connect.Query(`
		SELECT
			country_code,
			os_id,
			browser_id,
			categories,
			action_day,
			action_time
		FROM
			example`)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			country               string
			os, browser           uint8
			categories            []int16
			actionDay, actionTime time.Time
		)
		if err := rows.Scan(
			&country,
			&os,
			&browser,
			&categories,
			&actionDay,
			&actionTime,
		); err != nil {
			log.Fatal(err)
		}
		log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s",
			country, os, browser, categories, actionDay, actionTime,
		)
	}

	ctx := context.Background()
	rows, err = connect.QueryContext(context.WithValue(ctx, clickhouse.QueryID, "dummy-query-id"), `
		SELECT
			country_code,
			os_id,
			browser_id,
			categories,
			action_day,
			action_time
		FROM
			example`)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			country               string
			os, browser           uint8
			categories            []int16
			actionDay, actionTime time.Time
		)
		if err := rows.Scan(
			&country,
			&os,
			&browser,
			&categories,
			&actionDay,
			&actionTime,
		); err != nil {
			log.Fatal(err)
		}
		log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s",
			country, os, browser, categories, actionDay, actionTime,
		)
	}
}
