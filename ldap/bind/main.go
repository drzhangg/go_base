package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

func main() {
	//l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))

	ldapServer := ""
	ldapPort := 389
	user := "cn=admin,dc=example,dc=com"
	passwd := "123456"

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))

	if err != nil {
		fmt.Println("ldap dial failed,err:",err)
	}
	defer l.Close()
	//err = l.Bind(user, passwd)
	//if err != nil {
	//	// authenticated
	//	fmt.Println("bind err:", err)
	//}

	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username:           user,
		Password:           passwd,
	})
	if err != nil {
		log.Fatalf("Failed to bind: %s\n", err)
	}

	//res,err := l.WhoAmI(nil)
	//if err != nil {
	//	log.Fatalf("Failed to call WhoAmI(): %s\n", err)
	//}
	//fmt.Printf("I am: %s\n", res.AuthzID)

	//l.Compare(user,passwd)
	
	err = l.Add(&ldap.AddRequest{
		DN:         "cn=newuser,ou=users,dc=example,dc=com",
		Attributes: []ldap.Attribute{
			{
				Type: "userPassword",
				Vals: []string{"password123"},
			},
			{
				"objectClass", []string{"inetOrgPerson"},
			},
			{
				"sn",[]string{"newuser"},
			},
		},
		Controls:   nil,
	})
	if err !=nil{
		log.Println("add err:",err)
	}
}
