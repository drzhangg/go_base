package main

import (
	"fmt"
	"go_base/interface/demo1/db"
)

func main() {
	rs := []*db.RolePermission{
		{
			RoleId:       "123",
			PermissionId: "abc",
			Permission: &db.Permission{
				Age:  20,
				Name: "jerry",
			},
		},
		{
			RoleId:       "456",
			PermissionId: "def",
			Permission: &db.Permission{
				Age:  26,
				Name: "zhang",
			},
		},
	}

	s := PermissionService.convertToFePermission(db.RolePermissionList(rs))
	//fmt.Printf("s:%#v", s)
	for _, v := range s {
		fmt.Println(v)
	}


}
