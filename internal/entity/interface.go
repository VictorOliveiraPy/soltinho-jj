package entity

type UserRepositoryInterface interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}
