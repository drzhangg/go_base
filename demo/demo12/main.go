package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
)

func main() {
	//stringToSign := "GET\n/iaas/\naccess_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek3a"
	stringToSign := "GET\n/iaas/\naccess_key_id=GBPFITMISNISAOHQBYGC&action=DescribeClusters&apps.n%5B%5D=app-n9ro0xcp&owner=usr-zWBOsvlF&signature_method=HmacSHA256&signature_version=1&status.n=active&time_stamp=2024-02-27T11%3A09%3A10Z&version=1&zone=xining"
	secretAccessKey := "QxarSIh5sdB25RkjrvdHc0mcDS01Klrm3exJKD0I"

	h := hmac.New(sha256.New, []byte(secretAccessKey))
	h.Write([]byte(stringToSign))
	sign := h.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(sign)
	encodedSignature := url.QueryEscape(signature)

	fmt.Println(encodedSignature)
}
