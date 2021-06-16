package entradas
import (

	// "goApp/common"
	
)
// "fmt"
// "github.com/jinzhu/gorm"
// "goApp/common"
//"fmt"   "errors"  "github.com/gin-gonic/gin"   "net/http"


type Entradas struct {
	Id          uint	`gorm:"primary_key"`
	Event       uint   	`json:"event"`
	Price       int     `json:"price"`
}


type User struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email; unique_index"`
	Bio          string  `gorm:"column:bio; size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password; not null"`
	Type		 string	 `gorm:"column:type;" default:'client'`
}



//De momento no hace falta esta funcion porque pillo el usuario de myUser
// func FindOneUser(condition interface{}) (User, error) {
// 	db := common.GetDB()
// 	var model User
// 	err := db.Where(condition).First(&model).Error
// 	return model, err
// }