package routes

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/klishchov-bohdan/myrepos/internal/models/login"
	"github.com/klishchov-bohdan/myrepos/internal/services"
	"github.com/klishchov-bohdan/myrepos/pkg/token"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
)

type Routes struct {
	service *services.Service
}

func New(service *services.Service) *Routes {
	return &Routes{
		service: service,
	}
}

func (rts *Routes) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := &login.Request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := rts.service.User.GetByEmail(req.Email)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		// authenticated

		err = godotenv.Load("config/.env")
		if err != nil {
			http.Error(w, "Cant load .env file", http.StatusInternalServerError)
		}
		AccessTokenLifeTime, err := strconv.Atoi(os.Getenv("AccessTokenLifeTime"))
		if err != nil {
			http.Error(w, "Cant convert AccessTokenLifeTime", http.StatusInternalServerError)
		}
		AccessSecret := os.Getenv("AccessSecret")
		RefreshTokenLifeTime, err := strconv.Atoi(os.Getenv("RefreshTokenLifeTime"))
		if err != nil {
			http.Error(w, "Cant convert RefreshTokenLifeTime", http.StatusInternalServerError)
		}
		RefreshSecret := os.Getenv("RefreshSecret")

		tokenString, err := token.GenerateToken(user.ID, AccessTokenLifeTime, AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := token.GenerateToken(user.ID, RefreshTokenLifeTime, RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := &login.Response{
			AccessToken:  tokenString,
			RefreshToken: refreshString,
		}
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Only Post method", http.StatusMethodNotAllowed)
	}

}

func (rts *Routes) Profile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tokenString := token.GetTokenFromBearerString(r.Header.Get("Authorization"))
		AccessSecret := os.Getenv("AccessSecret")
		claims, err := token.ValidateToken(tokenString, AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		user, err := rts.service.User.GetByID(claims.ID)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		resp := login.UserResponseProfile{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
	}

}

//func (rts *Routes) Refresh(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "POST":
//
//		w.WriteHeader(http.StatusOK)
//		json.NewEncoder(w).Encode(resp)
//
//	default:
//		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
//	}
//
//}
