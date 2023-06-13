package auth

type Role struct {
	roleName string
}

var (
	Unknown = Role{""}
	Guest   = Role{"guest"}
	User    = Role{"user"}
	Admin   = Role{"admin"}
)
