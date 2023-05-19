package usecase

import (
	"golang-pkg/internal/marketplace/repository"
)

func CreateNewItem(body *marketplace.Items) error {
	if err := marketplace.ValidCreation(body); err != nil {
		return err
	}
	repository.Create(body)
	return nil
}

func GetTypes() []marketplace.ProductType {
	var body []marketplace.ProductType
	body = repository.GetTypes()
	return body
}

func CreateType(newType string) {
	repository.CreateType(newType)
}

func DeleteItem(id int) {
	repository.Del(id)
	return
}

func UpdateItem(body marketplace.Items) marketplace.Items {
	var newBody marketplace.Items
	newBody = repository.UpdateById(body)
	return newBody
}

func GetItem(id int) marketplace.Items {
	var body marketplace.Items
	body = repository.GetItem(id)
	return body
}

func GetItems(startId int, productType int) []marketplace.Items {
	var body []marketplace.Items
	body = repository.GetItems(startId, productType)
	return body
}
