package repositories

import (
	"github.com/daddydemir/kirtasiye-projesi/config"
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/security"
)

func StationeryGetAll() []models.Stationery {
	var stationery []models.Stationery
	config.DB.Find(&stationery)
	return stationery
}

func StationeryGetById(stationeryId string) models.Stationery {
	var stationery models.Stationery
	config.DB.Find(&stationery, "id = ?", stationeryId)
	return stationery
}

func StationeryAdd(stationery models.Stationery) {
	stationery.Password = security.HashPassword(stationery.Password)
	config.DB.Create(&stationery)
}

func StationeryUpdate(stationery models.Stationery) {
	config.DB.Save(&stationery)
}

func StationeryDelete(stationeryId string) {
	config.DB.Delete(models.Stationery{}, "id = ?", stationeryId)
}

func UpdateSImage(url string, stationeryId string) {
	var stationery models.Stationery = StationeryGetById(stationeryId)
	stationery.ImagePath = url
	config.DB.Save(&stationery)
}

func StationeryByName(companyName string) (models.Stationery, bool) {
	var stationery models.Stationery
	config.DB.Find(&stationery, "company_name = ?", companyName)
	if stationery.CompanyName == "" {
		return models.Stationery{}, false
	} else {
		return stationery, true
	}

}
