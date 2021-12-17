package filesystem

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"mysite/repositories/models"
	"os"
)

type UserFileRepository struct {
}

func (ufr *UserFileRepository) GetByEmail(Email string) (user *models.User, err error) {
	userRepo, err := ioutil.ReadDir("./datastore/files/users/")
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range userRepo {
		file, err := os.Open("./datastore/files/users/" + fileInfo.Name())
		if err != nil {
			return nil, err
		}
		defer file.Close()
		user := &models.User{}
		var data []byte
		for {
			chunk := make([]byte, 64)
			n, err := file.Read(chunk)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			data = append(data, chunk[:n]...)
		}
		err = json.Unmarshal(data, &user)
		if err != nil {
			return nil, err
		}
		if user.Email == Email {
			return user, nil
		}
	}
	return nil, nil
}
