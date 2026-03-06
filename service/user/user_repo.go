package user_service

type UserService interface {
	RegisterUser(input UserRegisterRequest) (*User, error)
	LoginUser(input LoginRequest) (string, error)
	DeleteUser(userID uint) error
}

type UserRepo interface {
	RegisterUser(user *User) error
	LoginUser(email, password string) (*User, error)
	DeleteUser(ID uint) error
}
