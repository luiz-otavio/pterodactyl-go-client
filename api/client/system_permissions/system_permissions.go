package system_permissions

type PermissionInfo struct {
	Description string            `json:"description"`
	Keys        map[string]string `json:"keys"`
}

type ClientPermissions []PermissionInfo

type ClientPermission interface {
	GetPermissions() (ClientPermissions, error)
}
