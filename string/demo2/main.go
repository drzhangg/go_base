package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "u001b[3;1H u001b[3;1Hu001b[4;1Hu001b[KI 2.yaml [Modified] 3/3 100%u001b[3;1H"

	//regexp.MustCompile(`u001b[%d;1H`)

	rege := regexp.MustCompile("u001b\\[[0-9];1H")


	//rs := rege.FindStringSubmatch(s)
	//fmt.Println("len::",len(rs))

	rr := rege.FindAllString(s,-1)
	fmt.Println(rr)
	if len(rr) > 0{
		result := rege.ReplaceAllString(rr[len(rr)-1],"")
		fmt.Println("re:",result)

	}

	//rege.

	//if len(rs) > 0 {
	//	result := rege.ReplaceAllString(rs[len(rs)-1],"")
	//	fmt.Println("result:",result)
	//}

	//ns := rege.ReplaceAllString(s,"")
	//fmt.Println(ns)

	//ok,_ := regexp.MatchString("u001b\\[\\d;1H",s)
	//fmt.Println("ok::",ok)

	//ns := strings.ReplaceAll(s,"\\","")
	//fmt.Println(ns)
	//
	//b := strings.ReplaceAll(ns,"u001b","%c")
	//fmt.Println(b)
	//
	//ss := []interface{}{}
	//for i := 0; i < strings.Count(ns,"u001b"); i++ {
	//	ss = append(ss, 0x1B)
	//}
	//fmt.Printf(b,ss...)
}
