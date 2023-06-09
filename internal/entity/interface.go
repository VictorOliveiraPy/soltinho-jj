package entity

type UserRepositoryInterface interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindById(id string) (*User, error)
}

type GymRepositoryInterface interface {
	Create(gym *Gym) error
	FindByName(gym_name string) (*Gym, error)
	FindById(id string) (*Gym, error)
}

type StudentRepositoryInterface interface {
	Create(student *Student) error
}
