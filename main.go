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
}
