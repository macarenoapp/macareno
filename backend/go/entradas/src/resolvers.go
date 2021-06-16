package entradas

import (
	// "fmt"
	"goApp/common"

)

//Create entrada
func CreateEntrada(data interface{}) error{
	db:=common.GetDB();
	err:=db.Create(data).Error
	return err
}

//Get all Entradas
func GetAllEntradas(data interface{}) error{
	db:=common.GetDB();
	err:=db.Find(data).Error
	return err;
}
//GET ONE entrada by ID
func GetEntradaById(data, id interface{}) error {
	db := common.GetDB()
	err := db.Where("id = ?", id).First(data).Error
	return err
}

//UPDATE entrada
func UpdateEntrada(data interface{}) error{
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

//DELETE
func DeleteEntrada(data, id interface{}) error {
	db := common.GetDB()
	err := db.Where("id = ?", id).Delete(data).Error
	return err
}

