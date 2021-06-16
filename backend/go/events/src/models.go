package events
import (
	// "fmt"
	// "goApp/common"
	// "github.com/jinzhu/gorm"
	
)

// "github.com/jinzhu/gorm"
// "goApp/common"
//"fmt"   "errors"  "github.com/gin-gonic/gin"   "net/http"


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

type Discotecas struct {
	Id          uint
	Name        string   `json:"name"`
	Description string   `json:description`
	PostalCode  uint     `json:postal`
	Location    string   `json:location`
	User        uint     `json:user`
	Views		int      `json:"views"`
}

// type DiscoEventModel struct {
// 	gorm.Model
// 	Discoteca     Discotecas
// 	DiscotecaID   uint
// 	Event         Events
// 	EventID       uint
// }


type User struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email; unique_index"`
	Bio          string  `gorm:"column:bio; size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password; not null"`
	Type		 string	 `gorm:"column:type;" default:'client'`
}

// func AutoMigrate() {
// 	db := common.GetDB()
// 	db.AutoMigrate(&DiscoEventModel{})
// }

// func CreateEventDisco(event Events, discotecaId uint) error {
// 	// fmt.Println("FAvorite de: ",user," A la discoteca: ", discoteca)
// 	fmt.Println("Relacionamos el evento: ",event," a la discoteca: ",discotecaId)
// 	db := common.GetDB()
// 	var discoEvent DiscoEventModel
// 	err := db.FirstOrCreate(&discoEvent, &DiscoEventModel{
// 		DiscotecaID:   discotecaId,
// 		EventID:  	   event.Id,
// 	}).Error
// 	return err
// }