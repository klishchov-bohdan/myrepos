package main

import (
	"fmt"
	"log"
	"mysite/repositories/filesystem"
)

func main() {
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

	pr := &filesystem.ProductFileRepository{}
	products, err := pr.GetByCategory("Pizza")
	if err != nil {
		log.Fatal(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}
}
