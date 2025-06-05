package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/joaolima7/api1_goexpert/internal/dto"
	"github.com/joaolima7/api1_goexpert/internal/entity"
	"github.com/joaolima7/api1_goexpert/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB        database.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExperiesIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExperiesIn int) *UserHandler {
	return &UserHandler{UserDB: userDB, Jwt: jwt, JwtExperiesIn: jwtExperiesIn}
}

// GetJWT godoc
// @Summary      Get JWT token
// @Description  Get JWT token for user authentication
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.GetJWTInput  true  "User credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404 {object}  Error
// @Failure      400 {object}  Error
// @Router       /users/generate-token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "Invalid JSON: " + err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: "User not found: " + err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !u.ValidatePasword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExperiesIn)).Unix(),
	})

	accessToken := dto.GetJWTOutput{AccessToken: token}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Create user
// @Description  Create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserInput  true  "User data"
// @Success      201
// @Failure      500  {object}  Error
// @Router	   /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: "Invalid user data: " + err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(*u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: "Database error: " + err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
