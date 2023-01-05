package main

import (
	"go_base/interface/demo1/db"
	"go_base/interface/demo1/service"
)

type permissionService int

var PermissionService = new(permissionService)


type permissionList interface {
	Each(func(permission *db.Permission))
}

func (s *permissionService) convertToFePermission(ps permissionList) []*service.FePermission {

	r := make([]*service.FePermission, 0)

	ps.Each(func(v *db.Permission) {
		if v == nil {
			return
		}

		r = append(r, &service.FePermission{
			Permission: *v,
			Access:     false,
			Children:   nil,
		})
	})

	return r
}
