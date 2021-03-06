package auth

import (
	"github.com/daddydemir/kirtasiye-projesi/models"
	"github.com/daddydemir/kirtasiye-projesi/repositories"
)

func StationeryAuth(tokenString string, stationery models.Stationery) (bool, map[string]string) {
	tokenStationery, status := repositories.StationeryByName(TokenParser(tokenString))
	if status {
		if tokenStationery.Id == stationery.Id {
			return true, map[string]string{"message": "Kullanıcı gereken yetkilere sahip."}
		} else {
			return false, map[string]string{"message": "Yetkisiz kullanıcı."}
		}
	} else {
		return false, map[string]string{"message": "Yetkisiz kullanıcı."}
	}
}
