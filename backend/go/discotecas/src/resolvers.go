package discotecas

import (
	// "fmt"
	"goApp/common"

)

//Create discoteca
func CreateDiscoteca(data interface{}) error{
	db:=common.GetDB();
	err:=db.Create(data).Error
	return err
}

//Get all Discotecas
func GetAllDiscotecas(data interface{}) error{
	db:=common.GetDB();
	err := db.Order("views desc").Find(data).Error
	return err;
}
//GET ONE discoteca by ID
func GetDiscotecaById(data, id interface{}) error {
	db := common.GetDB()
	err := db.Where("id = ?", id).First(data).Error
	return err
}

//Get discotecas byUser
func GetDiscotecaByUser(user uint64, data interface{}) error{
	db := common.GetDB();
	err := db.Where("User = ?", user).Find(data).Error

	return err
}




//Get Discotecas By likes
func GetDiscotecasLiked(data interface{}) error{
	db:=common.GetDB();
	err := db.Order("views desc").Find(data).Error
	return err;
}
//Favorites count
func favoritesCount(discoteca Discotecas) uint {
	db := common.GetDB()
	var count uint
	db.Model(&FavoriteModel{}).Where(FavoriteModel{
		FavoriteID: discoteca.Id,
	}).Count(&count)
	return count
}

//UPDATE discoteca
func UpdateDiscoteca(data interface{}) error{
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

//DELETE
func DeleteDiscoteca(data, id interface{}) error {
	db := common.GetDB()
	err := db.Where("id = ?", id).Delete(data).Error
	return err
}

