package filesystem

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mysite/internal/models"
	"os"
)

type UserFileRepository struct {
}

func (ufr *UserFileRepository) GetByEmail(email string) (user *models.User, err error) {
	userRepo, err := ioutil.ReadDir("./internal/datastore/files/users/")
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range userRepo {
		file, err := os.Open("./internal/datastore/files/users/" + fileInfo.Name())
		if err != nil {
			return nil, err
		}
		defer file.Close()
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
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

func (ufr *UserFileRepository) Create(user *models.User) (createdUser *models.User, err error) {
	userRepo, err := ioutil.ReadDir("./internal/datastore/files/users/")
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range userRepo {
		file, err := os.Open("./internal/datastore/files/users/" + fileInfo.Name())
		if err != nil {
			return nil, err
		}
		defer file.Close()
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
		checkUser := &models.User{}
		err = json.Unmarshal(data, &checkUser)
		if err != nil {
			return nil, err
		}
		if checkUser.Email == user.Email {
			return nil, errors.New("user is already exists in " + fileInfo.Name())
		}
	}
	fileName := "user_" + fmt.Sprint(len(userRepo)+1) + ".json"
	file, err := os.Create("./internal/datastore/files/users/" + fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	_, err = file.WriteString(string(userJSON))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(userJSON, &createdUser)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
