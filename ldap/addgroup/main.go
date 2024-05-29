package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

func main() {

	// cn=admin,dc=zhang,dc=com
	// Wlkj131q
	addr := ""
	port := "31327"

	username := "cn=admin,dc=zhang,dc=com"

	//l,err := ldap.DialURL(fmt.Sprintf("ldap://%s:%s", addr, port))
	//if err !=nil{
	//	log.Fatal(err)
	//}
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%s", addr, port))
	//l,err := ldap.DialURL("ldap://1.116.145.159:389")
	if err != nil {
		fmt.Println("ldap dial failed,err:",err)
	}
	defer l.Close()

	err = l.Bind(username,"Wlkj131q")
	if err != nil {
		log.Fatal(err)
	}


	//searchRequest1 := ldap.NewSearchRequest(
	//	"dc=zhang,dc=com",
	//	ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,0,0,false,
	//	"givenName=zhangsan",
	//	[]string{"uid","cn","mail","primarygroup","givenName"},
	//	nil,
	//	)

	result,err := l.Search(ldap.NewSearchRequest("dc=zhang,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,0,0,false,
		fmt.Sprintf("(givenName=%s)", "zhangsan"),
		[]string{"uid","cn","mail","givenName"},
		nil))
	if err != nil {
		fmt.Println("failed to query LDAP: %w", err)
		return
	}

	log.Println("Got", len(result.Entries), "search results")

	for _,val := range result.Entries{
		fmt.Println("val::",val)
	}

	fmt.Println(result.Entries[0].GetAttributeValue("uid"), result.Entries[0].GetAttributeValue("givenName"), result.Entries[0].GetAttributeValue("mail"))


	//l.Search(&ldap.SearchRequest{
	//	BaseDN:           "dc=eryajf,dc=net",
	//	Scope:            0,
	//	DerefAliases:     0,
	//	SizeLimit:        0,
	//	TimeLimit:        0,
	//	TypesOnly:        false,
	//	Filter:           "",
	//	Attributes:       nil,
	//	Controls:         nil,
	//	EnforceSizeLimit: false,
	//})
}
