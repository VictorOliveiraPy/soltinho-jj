package web

import (
	//"context"
	//"database/sql"
	"encoding/json"

	//"fmt"
	"net/http"
	//"strings"
	//"time"

	//"github.com/VictorOliveiraPy/internal/dto"
	"github.com/VictorOliveiraPy/internal/service"

	//"github.com/VictorOliveiraPy/internal/infra/logger"
	//"github.com/VictorOliveiraPy/internal/infra/web/handlers"
	"github.com/VictorOliveiraPy/internal/usecase"
	//"go.uber.org/zap"
	//db "github.com/VictorOliveiraPy/internal/infra/database"
	//"github.com/go-chi/chi/v5"
	//"github.com/go-chi/jwtauth"
)




type WebUserHandler struct {
	UserService   *service.UserService
}

func NewWebUserHandler(userService *service.UserService) *WebUserHandler {
	return &WebUserHandler{
		UserService: userService,
	}
}


// type EntityHandler struct {
// 	gymDB *sql.DB
// 	*db.Queries
// }

// func NewEntityHandler(dbConn *sql.DB) *EntityHandler {
// 	return &EntityHandler{
// 		gymDB:   dbConn,
// 		Queries: db.New(dbConn),
// 	}
// }


// // GetJWT godoc
// // @Summary      Get a user JWT
// // @Description  Get a user JWT
// // @Tags         users
// // @Accept       json
// // @Produce      json
// // @Param        request   body     dto.GetJWTInput  true  "user credentials"
// // @Success      200  {object}  dto.GetJWTOutput
// // @Failure      404  {object}  Error
// // @Failure      500  {object}  Error
// // @Router       /users/generate_token [post]
// func (handler *EntityHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
// 	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
// 	jwtExpiresIn := r.Context().Value("JwtExperesIn").(int)
// 	var user dto.GetJWTInput
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		handlers.BadRequestHandler(w, err)
// 		return
// 	}

// 	u, err := handler.Queries.GetUserByEmail(context.Background(), user.Email)
// 	if err != nil {
// 		handlers.NotFoundHandler(w, err)
// 		return
// 	}

// 	if !u.ValidatePassword(user.Password) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	_, tokenString, _ := jwt.Encode(map[string]interface{}{
// 		"sub": u.ID,
// 		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
// 	})
// 	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(accessToken)
// }

// // Create user godoc
// // @Summary      Create user
// // @Description  Create user
// // @Tags         users
// // @Accept       json
// // @Produce      json
// // @Param        request     body      dto.User  true  "user request"
// // @Success      201
// // @Failure      500         {object}  Error
// // @Router       /users [post]
// func (h *EntityHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	logger.Info("New user creation initiated",
// 		zap.String("route", "/create-user"),
// 	)
// 	ctx := context.Background()
// 	var request dto.User

// 	json.NewDecoder(r.Body).Decode(&request)

// 	email, _ := h.Queries.GetUserByEmail(ctx, request.Email)

// 	if email.Email != "" {
// 		err := fmt.Errorf(strings.TrimRight("O email fornecido já está em uso.", "\n.,;:!?"))
// 		logger.Error("Erro durante a criação do usuário", err,
// 			zap.String("email", email.Email),
// 		)
// 		w.WriteHeader(http.StatusConflict)
// 		return
// 	}

// 	user, err := entity.NewUser(request.Username, request.Password, request.Email, request.RoleID)
// 	if err != nil {
// 		logger.Error("Erro ao criar objetos de um novo usuário", err)
// 		handlers.BadRequestHandler(w, err)
// 		return
// 	}

// 	err = h.Queries.CreateUser(ctx, db.CreateUserParams{
// 		ID:       user.ID,
// 		Email:    user.Email,
// 		Username: user.UserName,
// 		Password: user.Password,
// 		RoleID:   user.RoleID,
// 		Active:   user.Active,
// 	})

// 	if err != nil {
// 		logger.Error("Erro ao criar um novo usuário", err)
// 		handlers.BadRequestHandler(w, err)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	logger.Info("New user created",
// 		zap.String("route", "/create-user"),
// 		zap.String("user_id", user.ID),
// 	)

// }

// // Get user godoc
// // @Summary      Get user
// // @Description  Get user full profile
// // @Tags         users
// // @Accept       json
// // @Produce      json
// // @Success      200
// // @Failure      500         {object}  Error
// // @Router       /users [get]
// func (h *EntityHandler) GetUserFullProfile(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	json.NewDecoder(r.Body).Decode(&id)

// 	if id == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	row, err := h.Queries.GetUserCompleteProfile(context.Background(), id)
// 	if err != nil {
// 		err := fmt.Errorf(strings.TrimRight("ID fornecido Nao existe.", "\n.,;:!?"))
// 		logger.Error("ID fornecido Nao existe.", err,
// 			zap.String("ID", id),
// 		)
// 		handlers.NotFoundHandler(w, err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(row)

// }

func handleConflictError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusConflict)
	w.Write([]byte(err.Error()))
}


func (h *WebUserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.UserInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UserService.CreateUser(r.Context(), dto)
	if err != nil {
		if e, ok := err.(service.EmailAlreadyExistsError); ok {
			handleConflictError(w, e)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Usuário criado com sucesso"))
}
