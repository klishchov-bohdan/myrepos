package filesystem

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"mysite/repositories/models"
	"os"
)

type ProductFileRepository struct {
}

func (pfr *ProductFileRepository) GetByCategory(category string) (products []*models.Product, err error) {
	productRepo, err := ioutil.ReadDir("./datastore/files/products/")
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range productRepo {
		file, err := os.Open("./datastore/files/products/" + fileInfo.Name())
		if err != nil {
			return nil, err
		}
		defer file.Close()
		product := &models.Product{}
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
		err = json.Unmarshal(data, &product)
		if err != nil {
			return nil, err
		}
		if product.Category == category {
			products = append(products, product)
		}
	}
	return products, nil
}
