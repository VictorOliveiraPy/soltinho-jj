package usecase

import (
	"time"

	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/errors"
	"github.com/go-chi/jwtauth"
)

type GetJWTInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWTOutput struct {
	AccessToken string `json:"acess_token"`
}

type GetTokenUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

func GetTokenUserUseCase(
	UserRepository entity.UserRepositoryInterface,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: UserRepository,
	}
}

func (u *GetTokenUseCase) checkEmailExists(email string) (*entity.User, error) {
	existingUser, _ := u.UserRepository.FindByEmail(email)
	if existingUser == nil {
		return nil, errors.EmailNotFound{Email: email}
	}
	return existingUser, nil
}

func (u *GetTokenUseCase) GetUserToken(jwt *jwtauth.JWTAuth, jwtExpiresIn int, input GetJWTInput) (GetJWTOutput, error) {
	user, err := u.checkEmailExists(input.Email)
	if err != nil {
		return GetJWTOutput{}, err
	}

	if !user.ValidatePassword(input.Password) {
		return GetJWTOutput{}, errors.PasswordInvalid{Password: "Senha inválida"}
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID,
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	return GetJWTOutput{AccessToken: tokenString}, nil

}
