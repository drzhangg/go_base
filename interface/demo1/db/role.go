package db

type RolePermission struct {
	RoleId       string
	PermissionId string
	Permission   *Permission
}

type Permission struct {
	Age  int
	Name string
}

type RolePermissionList []*RolePermission

func (s RolePermissionList) Each(fn func(*Permission)) {
	for _, v := range s {
		fn(v.Permission)
	}
}
