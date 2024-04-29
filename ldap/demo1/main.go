package main

import (
	"fmt"
	"github.com/go-ldap/ldap"
)

func main() {
	ldapServer := "xxx"
	ldapPort := 389
	user := "cn=admin,dc=example,dc=com"
	passwd := "123456"

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
	// be sure to add error checking!
	defer l.Close()
	err = l.Bind(user, passwd)
	if err != nil {
		// authenticated
		fmt.Println("err:", err)
	}


	// 准备新用户的属性
	addReq := ldap.NewAddRequest("cn=newuser,ou=users,dc=example,dc=com",[]ldap.Control{})
	addReq.Attribute("objectClass", []string{"inetOrgPerson"})
	addReq.Attribute("cn", []string{"newuser"})
	addReq.Attribute("sn", []string{"newuser"})
	addReq.Attribute("mail", []string{"newuser@example.com"})
	addReq.Attribute("userPassword", []string{"password123"}) // 设置用户密码

	// 执行添加用户操作
	err = l.Add(addReq)
	if err != nil {
		fmt.Println("Error adding new user:", err)
		return
	}


	search := &ldap.SearchRequest{
			BaseDN: "dc=example,dc=com",
			Filter: "(objectclass=*)",
	}
	searchResults, err := l.Search(search)

	for _, v := range searchResults.Entries{

		for _,v1 := range v.Attributes{
			fmt.Printf("%#v\n",v1)
		}
	}

	//control := []ldap.Control{}
	////control = append(control, ldap.New)
	//
	//l.Add(ldap.NewAddRequest("",[]ldap.Control{
	//	{}
	//}))

	//searchRequest := ldap.NewSearchRequest(
	//	"dc=glauth,dc=com",
	//	ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
	//	fmt.Sprintf("(uid=%s)", uid),
	//	[]string{"uid", "cn", "mail", "primarygroup", "givenname"},
	//	nil,
	//)
	//
	//sr, err := l.Search(searchRequest)
	//
	//
	//search := &SearchRequest{
	//	BaseDN: "dc=example,dc=com",
	//	Filter: "(objectclass=*)",
	//}
	//searchResults, err := l.Search(search)
	// be sure to add error checking!

}
