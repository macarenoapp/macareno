package discotecas
import (
	"github.com/jinzhu/gorm"
	"goApp/common"
	"fmt"
)
// "fmt"
// "github.com/jinzhu/gorm"
// "goApp/common"
//"fmt"   "errors"  "github.com/gin-gonic/gin"   "net/http"


type Discotecas struct {
	Id          uint
	Name        string   `json:"name"`
	Description string   `json:description`
	PostalCode  uint     `json:postal`
	Location    string   `json:location`
	User        uint     `json:user`
	Views		int      `json:"views"`
	Likes       uint
	Liked       bool
	Events		[]Events
}


type FavoriteModel struct {
	gorm.Model
	Favorite     Discotecas
	FavoriteID   uint
	FavoriteBy   User
	FavoriteByID uint
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

type Events struct {
	Id            uint
	Name          string   `json:"name"`
	Description   string   `json:description`
	Start_date    string   `json:start_date`
	Start_time    string   `json:start_time`
	Entradas      uint     `json:entradas`
	Entradas_sold uint 	   `json:entradas_sold`
	Price 		  uint     `json:price`
	Consumicion   bool     `json:consumicion`
	Discoteca     uint     `json:discoteca`
}


// type DiscoEventModel struct {
// 	gorm.Model
// 	Discoteca     Discotecas
// 	DiscotecaID   uint
// 	Event         Events
// 	EventID       uint
// }

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&FavoriteModel{})
}



//De momento no hace falta esta funcion porque pillo el usuario de myUserModel
func FindOneUser(condition interface{}) (User, error) {
	db := common.GetDB()
	var model User
	err := db.Where(condition).First(&model).Error
	return model, err
}


//LIKE
func favoriteBy(user User, discoteca Discotecas) error {
	db := common.GetDB()
	var favorite FavoriteModel
	err := db.FirstOrCreate(&favorite, &FavoriteModel{
		FavoriteID:   discoteca.Id,
		FavoriteByID: user.ID,
	}).Error
	return err
}

//UNLIKE
func unFavoriteBy(user User, discoteca Discotecas) error {
	var unfavorite FavoriteModel
	db := common.GetDB()
	db.Where("favorite_id = ? AND favorite_by_id = ?", discoteca.Id, user.ID).First(&unfavorite)
	err := db.Delete(unfavorite).Error
	return err
}

//Get los eventos de una discoteca
func GetEventsDisco(id uint) []Events{
	var events []Events
	var eventModel Events
	db := common.GetDB()
	db.Model(eventModel).Where("discoteca = ?", id).Find(&events)

	return events
	
}

//Is favorited
func isFavoriteBy(discoteca Discotecas,user User) bool {
	db := common.GetDB()
	var favorite FavoriteModel
	err := db.Where(FavoriteModel{
		FavoriteID:   discoteca.Id,
		FavoriteByID: user.ID,
		// deleted_at:  null,
	}).First(&favorite).Error

	if err!=nil {
		fmt.Println("Error isfavorited ",err)
	} else {
		fmt.Println("IS FAVORITE--",favorite)
	}
	return favorite.ID != 0
}


//Add visit
func VisitDisco(discoteca Discotecas) error{
	db := common.GetDB()
	discoteca.Views ++;
	err := db.Save(discoteca).Error

	return err
}
