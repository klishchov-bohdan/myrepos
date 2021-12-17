package filesystem

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mysite/repositories/models"
	"os"
)

type UserFileRepository struct {
}

func (ufr *UserFileRepository) GetByEmail(Email string) (user *models.User) {
	var data []byte
	file, err := os.Open("./datastore/files/users/user_1.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for {
		chunk := make([]byte, 64)
		n, err := file.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, chunk[:n]...)
	}
	fmt.Println(data)
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Fatal(err)
	}

	return user
}
