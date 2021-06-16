package events

import (
	// "fmt"
	"errors"
	"goApp/common"
	"github.com/gin-gonic/gin"
	"net/http"
	// "strconv"
)

//"fmt"    "encoding/json"  
// para los coments
//	"io" "strings"
//	"net/http"
//	"io/ioutil"



func EventsRegister(router *gin.RouterGroup) {
	router.POST("/", EventCreate)
	router.PUT("/:id", EventUpdate)
	router.DELETE("/:id", EventDelete)
	
	// router.DELETE("/:id/favorite", EventUnfavorite)
	// router.POST("/:id/comments", EventCommentCreate)
	// router.DELETE("/:id/comments/:id", EventCommentDelete)
}

func EventsAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", EventList)
	router.GET("/:id", EventById)  
	
}

// router.GET("/:id/comments", EventCommentList)

func EventCreate(c *gin.Context){
	var event Events //Event que hemos creado
	c.BindJSON(&event);

	//Pillar el id de la discoteca

	err1 := CreateEvent(&event)

	if err1 != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, gin.H{"event":event})
		return
	}
	
}


//List Events
func EventList(c *gin.Context) {
	var event []Events

	//Busca las events y mete el resultado en la var event
	err := GetAllEvents(&event)
	
	if err != nil {
		c.JSON(http.StatusOK, "Not found")
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, gin.H{"events": event})
	}
}

///////// Find ONE
func EventById(c *gin.Context) {
	id := c.Params.ByName("id")	

	var event Events
	err := GetEventById(&event, id)
	
	if err != nil {
		c.JSON(http.StatusOK, "Event Not Found")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"event": event})
			return
	}
}
////////

//UPDATE event

func EventUpdate(c *gin.Context){
	var event Events
	var newEvent Events
	c.BindJSON(&newEvent);  //Aqui en teoria está la event que le hemos pasado por postman

	id := c.Params.ByName("id")
	err := GetEventById(&event, id) //Este es la event que he pillao con ese id, ¿para que? para comprobar que existe ese id

	// event.Name = newEvent.Name
	// event.Company = newEvent.Company
	// event.Events = newEvent.Events


	if err != nil { 
		c.JSON(http.StatusNotFound, "NOT FOUND")
	}else{ 
		c.BindJSON(&event)
		err = UpdateEvent(&event)//&event  Aqui hay que meterle la event nueva, con el c.BingJSON pero no me hace el json de la nueva
		if err != nil {
			c.JSON(http.StatusOK, "Not found")
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
			return
		}
	}

}


//DELETE event
func EventDelete(c *gin.Context){
	var event Events
	id := c.Params.ByName("id")

	err := DeleteEvent(&event, id)

	if err != nil{
		c.JSON(http.StatusNotFound, common.NewError("events", errors.New("Invalid id")))
		return
	} 

	c.JSON(http.StatusOK, gin.H{"event": "Delete Event"})
}
