package filesystem

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"mysite/internal/models"
	"os"
)

type OrderFileRepository struct {
}

func (ofr *OrderFileRepository) GetByID(id uint64) (order *models.Order, err error) {
	orderRepo, err := ioutil.ReadDir("./internal/datastore/files/orders/")
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range orderRepo {
		file, err := os.Open("./internal/datastore/files/orders/" + fileInfo.Name())
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
		err = json.Unmarshal(data, &order)
		if err != nil {
			return nil, err
		}
		if order.ID == id {
			return order, nil
		}
	}
	return nil, nil
}
