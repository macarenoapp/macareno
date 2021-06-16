package profile

import (
	"fmt"
	"goApp/common"

)

//Create profile
func CreateProfile(data interface{}) error{
	fmt.Println("CREATEEEEEEEE")
	fmt.Println(data)
	db:=common.GetDB();
	err:=db.Create(data).Error
	return err
}

//Get all Profile
func GetAllProfiles(data interface{}) error{
	db:=common.GetDB();
	err:=db.Find(data).Error
	return err;
}
//GET ONE profile by ID
func GetProfileById(data, id interface{}) error {   //Pillar el profile con la id de usuario que recibimos, no el id del profile
	db := common.GetDB()
	err := db.Where("user = ?", id).First(data).Error
	return err
}

//UPDATE profile
func UpdateProfile(data interface{}) error{
	fmt.Println("RESOLVER UPDATE");
	fmt.Println(data)
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

// func (model *ArticleModel) Update(data interface{}) error {
// 	db := common.GetDB()
// 	err := db.Model(model).Update(data).Error
// 	return err
// }

//DELETE
func DeleteProfile(data, id interface{}) error {
	db := common.GetDB()
	err := db.Where("id = ?", id).Delete(data).Error
	return err
}