package routes

import (
	"github.com/klishchov-bohdan/myrepos/internal/services"
	"net/http"
)

func SetRoutes(service *services.Service) {
	http.HandleFunc("/login", service.Login)
	http.HandleFunc("/registration", service.Registration)
	http.HandleFunc("/refresh", service.Refresh)
	http.HandleFunc("/profile", service.Profile)
}
