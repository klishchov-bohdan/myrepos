package services

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/klishchov-bohdan/myrepos/internal/middleware"
	"github.com/klishchov-bohdan/myrepos/internal/models"
	"github.com/klishchov-bohdan/myrepos/internal/models/login"
	"github.com/klishchov-bohdan/myrepos/internal/models/registration"
	"github.com/klishchov-bohdan/myrepos/internal/repositories"
	"github.com/klishchov-bohdan/myrepos/pkg/token"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
)

type Service struct {
	User repositories.UserRepositories
}

func New(UFRepo repositories.UserRepositories) *Service {
	return &Service{
		User: UFRepo,
	}
}

func (service *Service) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := &login.Request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := service.User.GetByEmail(req.Email)
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

func (service *Service) Registration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := &registration.Request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := service.User.GetByEmail(req.Email)
		fmt.Println(err)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		if user != nil {
			http.Error(w, "email already exists", http.StatusConflict)
			return
		}
		created, err := service.User.Create(models.NewUser(req.Name, req.Email, req.Password))
		if err != nil {
			http.Error(w, "cant create user", http.StatusInternalServerError)
			return
		}
		// created user

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

		tokenString, err := token.GenerateToken(created.ID, AccessTokenLifeTime, AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := token.GenerateToken(created.ID, RefreshTokenLifeTime, RefreshSecret)
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

func (service *Service) Profile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		accessString := token.GetTokenFromBearerString(r.Header.Get("Authorization"))
		AccessSecret := os.Getenv("AccessSecret")
		claims, err := middleware.CheckAccessToken(accessString, AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		user, err := service.User.GetByID(claims.ID)
		if err != nil {
			http.Error(w, "user not found", http.StatusUnauthorized)
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

func (service *Service) Refresh(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		err := godotenv.Load("config/.env")
		if err != nil {
			http.Error(w, "Cant load .env file", http.StatusInternalServerError)
		}
		accessTokenLifeTime, err := strconv.Atoi(os.Getenv("AccessTokenLifeTime"))
		if err != nil {
			http.Error(w, "Cant convert AccessTokenLifeTime", http.StatusInternalServerError)
		}
		accessSecret := os.Getenv("AccessSecret")
		refreshTokenLifeTime, err := strconv.Atoi(os.Getenv("RefreshTokenLifeTime"))
		if err != nil {
			http.Error(w, "Cant convert RefreshTokenLifeTime", http.StatusInternalServerError)
		}
		refreshSecret := os.Getenv("RefreshSecret")

		refreshString := token.GetTokenFromBearerString(r.Header.Get("Authorization"))
		claims, err := token.ValidateToken(refreshString, refreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := service.User.GetByID(claims.ID)
		if err != nil {
			http.Error(w, "cant get user", http.StatusUnauthorized)
			return
		}
		if user == nil {
			http.Error(w, "user not found", http.StatusUnauthorized)
			return
		}

		newAccessString, err := token.GenerateToken(user.ID, accessTokenLifeTime, accessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		newRefreshString, err := token.GenerateToken(user.ID, refreshTokenLifeTime, refreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := &login.Response{
			AccessToken:  newAccessString,
			RefreshToken: newRefreshString,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
	}

}
