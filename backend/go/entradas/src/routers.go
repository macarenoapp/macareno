package entradas

import (
	// "fmt"
	"errors"
	"goApp/common"
	"github.com/gin-gonic/gin"
	"net/http"
	
)

//"fmt"    "encoding/json"  
// "strconv" para los coments
//	"io" "strings"
//	"net/http"
//	"io/ioutil"



func EntradasRegister(router *gin.RouterGroup) {
	router.POST("/", EntradaCreate)
	router.PUT("/:id", EntradaUpdate)
	router.DELETE("/:id", EntradaDelete)
	
	// router.DELETE("/:id/favorite", EntradaUnfavorite)
	// router.POST("/:id/comments", EntradaCommentCreate)
	// router.DELETE("/:id/comments/:id", EntradaCommentDelete)
}

func EntradasAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", EntradaList)
	router.GET("/:id", EntradaById)  
	
}

// router.GET("/:id/comments", EntradaCommentList)

func EntradaCreate(c *gin.Context){
	var entrada Entradas
	c.BindJSON(&entrada);

	err:=CreateEntrada(&entrada)

	if err !=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, gin.H{"entrada":entrada})
		return
	}
}


//List Entradas
func EntradaList(c *gin.Context) {
	var entrada []Entradas

	//Busca las entradas y mete el resultado en la var entrada
	err := GetAllEntradas(&entrada)
	
	if err != nil {
		c.JSON(http.StatusOK, "Not found")
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		// c.JSON(http.StatusOK, entrada)
		// serializer := EntradasSerializer
		// c.JSON(http.StatusOK, gin.H{"entradas":entradaModel})
		c.JSON(http.StatusOK, gin.H{"entradas": entrada})
	}
}

///////// Find ONE
func EntradaById(c *gin.Context) {
	id := c.Params.ByName("id")	

	var entrada Entradas
	err := GetEntradaById(&entrada, id)
	
	if err != nil {
		c.JSON(http.StatusOK, "Entrada Not Found")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"entrada": entrada})
			return
	}
}
////////

//UPDATE entrada

func EntradaUpdate(c *gin.Context){
	var entrada Entradas
	var newEntrada Entradas
	c.BindJSON(&newEntrada);  //Aqui en teoria está la entrada que le hemos pasado por postman

	id := c.Params.ByName("id")
	err := GetEntradaById(&entrada, id) //Este es la entrada que he pillao con ese id, ¿para que? para comprobar que existe ese id

	// entrada.Name = newEntrada.Name
	// entrada.Company = newEntrada.Company
	// entrada.Events = newEntrada.Events


	if err != nil { 
		c.JSON(http.StatusNotFound, "NOT FOUND")
	}else{ 
		c.BindJSON(&entrada)
		err = UpdateEntrada(&entrada)//&entrada  Aqui hay que meterle la entrada nueva, con el c.BingJSON pero no me hace el json de la nueva
		if err != nil {
			c.JSON(http.StatusOK, "Not found")
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"entrada": entrada})
			return
		}
	}

}


//DELETE entrada
func EntradaDelete(c *gin.Context){
	var entrada Entradas
	id := c.Params.ByName("id")

	err := DeleteEntrada(&entrada, id)

	if err != nil{
		c.JSON(http.StatusNotFound, common.NewError("entradas", errors.New("Invalid id")))
		return
	} 

	c.JSON(http.StatusOK, gin.H{"entrada": "Delete Entrada"})

}