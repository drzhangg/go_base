package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
)

func main() {

	// access_key_id=GBPFITMISNISAOHQBYGC&action=DescribeClusters&clusters.n=%5Bcl-sz2chfvp%5D&owner=usr-zWBOsvlF&signature_method=HmacSHA256&signature_version=1&time_stamp=2024-02-28T17%3A17%3A37Z&version=1&zone=xn1a
	//stringToSign := "GET\n/iaas/\naccess_key_id=GBPFITMISNISAOHQBYGC&action=DescribeClusters&apps.n%5B%5D=app-n9ro0xcp&owner=usr-zWBOsvlF&signature_method=HmacSHA256&signature_version=1&status.n=active&time_stamp=2024-02-27T11%3A09%3A10Z&version=1&zone=xining"
	stringToSign := "GET\n/iaas/\naccess_key_id=GBPFITMISNISAOHQBYGC&action=DescribeClusters&clusters.n%5B%5D=cl-sz2chfvp&owner=usr-zWBOsvlF&signature_method=HmacSHA256&signature_version=1&time_stamp=2024-02-28T17%3A17%3A37Z&verbose=1&version=1&zone=xining"
	secretAccessKey := "QxarSIh5sdB25RkjrvdHc0mcDS01Klrm3exJKD0I"

	h := hmac.New(sha256.New, []byte(secretAccessKey))
	h.Write([]byte(stringToSign))
	sign := h.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(sign)
	encodedSignature := url.QueryEscape(signature)

	fmt.Println(encodedSignature)
	//m7w9chHrgyeN6h9P9Qm0b9gG8gbJ7uDP6OQaEJWItYw%3D
}
