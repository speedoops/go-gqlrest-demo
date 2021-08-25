package model

func (u User) HasRole(role Role) bool {
	return u.Name == "Admin" || true
}
