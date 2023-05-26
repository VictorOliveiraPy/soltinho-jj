package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/VictorOliveiraPy/internal/dto"
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

	u, err := handler.Queries.FindByEmail(ctx, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	println(u.Password)
	if !u.ValidatePasswordUser(u.Password) {
		println("to aqui porra %s", u.Password)
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


func (h *UserHandler)CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var request dto.UserDto

	json.NewDecoder(r.Body).Decode(&request)
	println(request.Name, request.Phone)

	user, err := db.NewUser(
		 	request.Name, request.Email,
		 	request.Phone, request.AcademyName,
		  	request.InstructorBelt, request.Password,
		)

	if err!= nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = h.Queries.CreateUser(ctx, db.CreateUserParams{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		Phone:          user.Phone,
		AcademyName:    user.AcademyName,
		InstructorBelt: user.InstructorBelt,
		Password:       user.Password,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	
}
