package main

import (
	"github.com/klishchov-bohdan/myrepos/internal/models"
	"github.com/klishchov-bohdan/myrepos/internal/repositories/filesystem"
	"github.com/klishchov-bohdan/myrepos/internal/routes"
	"github.com/klishchov-bohdan/myrepos/internal/services"
	"log"
	"net/http"
)

func main() {
	user := models.NewUser("John", "mail@mail.com", "12345")
	repo := &filesystem.UserFileRepository{}
	_, _ = repo.Create(user)
	service := services.New(repo)
	routes.SetRoutes(service)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
