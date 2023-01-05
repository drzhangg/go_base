package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

func main() {
	val,ok := os.LookupEnv("LOG_LEVEL")
	fmt.Println(val,ok)

	uid := uuid.New()
	fmt.Println(uid.String())
}
