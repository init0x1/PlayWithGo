package users

type User struct {
	UserID   string
	UserName string
	Password string
}

func (user User) CreateUser() string {
	return "User created: " + user.UserName
}
