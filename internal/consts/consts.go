package consts

const (
	TIME_OTP_REGISTERED int = 1
	TIME_ACCESSS_TOKEN int = 60 * 30
	TIME_REFRESH_TOKEN int64 = 3600 * 24 * 2
)

const (
	ADMIN   = "admin"
	MANAGER = "manager"
	EDITOR  = "editor"
	VIEWER  = "viewer"
	SUPPORT = "support"
)


var RoleDesc = map[string]string{
	ADMIN: "Has full access to all resources and system settings.",
	MANAGER: "Can manage user accounts, roles, and view reports.",
	EDITOR: "Can create, edit, and delete content but cannot manage users.",
	VIEWER: "Can only view content and reports without any editing rights.",
	SUPPORT: "Handles user queries and provides technical support.",
}
