package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := "GET\n/iaas/\naccess_key_id=GBPFITMISNISAOHQBYGC&action=DescribeClusters&apps.n%5B%5D=app-n9ro0xcp&owner=usr-zWBOsvlF&signature_method=HmacSHA256&signature_version=1&status.n=active&time_stamp=2024-02-27T11%3A09%3A10Z&version=1&zone=xining"

	re := regexp.MustCompile(`=([^&\s]+)`)
	output := re.ReplaceAllStringFunc(input, func(match string) string {
		value := strings.TrimPrefix(match, "=")
		return "=placeholder-" + strconv.Itoa(len(value))
	})

	fmt.Println(output)
}
