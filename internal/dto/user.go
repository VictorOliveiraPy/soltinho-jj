package dto


type UserDto struct {
	ID  string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	AcademyName string `json:"academy_name"`
	InstructorBelt string `json:"instructor_belt"`
	Password string `json:"-"`
}

type GetJWTInput struct {
	Email      string `json:"email"`
	Password string `json:"-"`

}

type GetJWTOutput struct {
	AccessToken string `json:"acess_token"`
}