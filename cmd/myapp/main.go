package main

import (
	"github.com/klishchov-bohdan/myrepos/internal/models"
	"github.com/klishchov-bohdan/myrepos/internal/repositories/filesystem"
	rts "github.com/klishchov-bohdan/myrepos/internal/routes"
	"github.com/klishchov-bohdan/myrepos/internal/services"
	"log"
	"net/http"
)

func main() {
	user := models.NewUser("John", "mail@mail.com", "12345")
	repo := &filesystem.UserFileRepository{}
	_, _ = repo.Create(user)
	service := services.New(repo)
	routes := rts.New(service)
	http.HandleFunc("/login", routes.Login)
	http.HandleFunc("/registration", routes.Registration)
	http.HandleFunc("/refresh", routes.Refresh)
	http.HandleFunc("/profile", routes.Profile)
	//http.HandleFunc("/refresh", routes.Refresh)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
