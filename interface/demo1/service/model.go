package service

import "go_base/interface/demo1/db"

type FePermission struct {
	db.Permission
	Level    int             `json:"level"`
	Access   bool            `json:"access"`
	Children []*FePermission `json:"children,omitempty"`
}

//type permissionService int
//
//var PermissionService = new(permissionService)
//
//
//type permissionList interface {
//	Each(func(permission db.Permission))
//}
//
//func (s *permissionService) convertToFePermission(ps permissionList)  {
//	return
//}
