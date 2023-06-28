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

func verifyPassword(
	comparePassword func(password string, hash string) bool,
	attainHash func(username string) (string, error),
	username string,
	password string,
) (bool, error) {
	hash, err := attainHash(username)
	if err != nil {
		return false, err
	}

	return comparePassword(password, hash), nil
}
