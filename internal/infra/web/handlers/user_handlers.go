package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/VictorOliveiraPy/internal/dto"
	"github.com/VictorOliveiraPy/internal/entity"

	db "github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	userDB *sql.DB
	*db.Queries
	Jwt *jwtauth.JWTAuth
}

type Error struct {
	Message string `json:"message"`
}

func NewUserHandler(dbConn *sql.DB) *UserHandler {
	return &UserHandler{
		userDB:  dbConn,
		Queries: db.New(dbConn),
	}
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
func (handler *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExperesIn").(int)

	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := handler.Queries.GetUserByEmail(ctx, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID,
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.UserDto  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var request dto.User

	json.NewDecoder(r.Body).Decode(&request)

	user, err := entity.NewUser(request.Username, request.Password, request.Email, request.RoleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.Queries.CreateUser(ctx, db.CreateUserParams{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.UserName,
		Password: user.Password,
		RoleID:   user.RoleID,
		Active:   user.Active,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

}
