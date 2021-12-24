package main

import (
	"fmt"
	"github.com/klishchov-bohdan/myrepos/internal/models"
	"github.com/klishchov-bohdan/myrepos/internal/repositories/filesystem"
	"log"
)

func main() {
	usr := &models.User{
		Email:        "user1432@gmail.com",
		PasswordHash: "pass2335",
		CreatedAt:    "2021-12-15 17:39:22",
	}
	sr := &filesystem.SupplierFileRepository{}
	suppliers, err := sr.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, supplier := range suppliers {
		fmt.Println(supplier)
	}

	ur := &filesystem.UserFileRepository{}
	user, err := ur.GetByEmail("example@test.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	createdUser, err := ur.Create(usr)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(createdUser)

	pr := &filesystem.ProductFileRepository{}
	products, err := pr.GetByCategory("Pizza")
	if err != nil {
		log.Fatal(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}

	or := &filesystem.OrderFileRepository{}
	order, err := or.GetByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(order)
}
