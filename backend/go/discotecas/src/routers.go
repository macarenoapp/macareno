package discotecas

import (
	"fmt"
	"errors"
	"goApp/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	// "reflect"
)


//"fmt"    "encoding/json"  
// "strconv" para los coments
//	"io" "strings"
//	"net/http"
//	"io/ioutil"

func DiscotecasRegister(router *gin.RouterGroup) {
	router.POST("/", DiscotecaCreate)
	router.PUT("/:id", DiscotecaUpdate)
	router.DELETE("/:id", DiscotecaDelete)
	router.POST("/:id/favorite", DiscotecaFavorite)
	router.POST("/:id/unfavorite", DiscotecaUnFavorite)
	router.GET("/:id/user", DiscotecasUser)
}

func DiscotecasAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", DiscotecaList)
	router.GET("/:id", DiscotecaById)  
	router.GET("/:id/liked", FavoritesCount)

}	

// router.GET("/:id/comments", DiscotecaCommentList)

func DiscotecaCreate(c *gin.Context){
	var discoteca Discotecas
	c.BindJSON(&discoteca);

	myUserModel := c.MustGet("my_user_model").(User)
	discoteca.User = myUserModel.ID; 
	
	err:=CreateDiscoteca(&discoteca)

	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, gin.H{"discoteca":discoteca})
		return
	}
}


//List Discotecas
func DiscotecaList(c *gin.Context) {
	var discoteca []Discotecas

	//Busca las discotecas y mete el resultado en la var discoteca
	err := GetAllDiscotecas(&discoteca)
	
	if err != nil {
		c.JSON(http.StatusOK, "Not found")
		c.AbortWithStatus(http.StatusNotFound)
	}else{

		c.JSON(http.StatusOK, gin.H{"discotecas": discoteca})
	}
}

///////// Find ONE
func DiscotecaById(c *gin.Context) {
	id := c.Params.ByName("id")	

	var discoteca Discotecas
	err := GetDiscotecaById(&discoteca, id)

	if err != nil {
		c.JSON(http.StatusOK, "Discoteca Not Found")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}else{
		err := VisitDisco(discoteca)
		if err != nil{
			c.JSON(http.StatusOK, "Discoteca Visit Fail")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}else{
			discoteca.Views ++;

			//Pillar los favoritos y si tiene o no por el user aqui
			myUserModel := c.MustGet("my_user_model").(User) //Usuario que da like
			count := favoritesCount(discoteca)
			discoteca.Likes=count;

			var eventsDisco []Events;
			// //Pasamos id a uint 64
			u, err := strconv.ParseUint(id, 10, 64)
			fmt.Println(err)

			eventsDisco = GetEventsDisco(uint(u))
			discoteca.Events = eventsDisco;
			
			if len(myUserModel.Username) > 1{//Hay usuario logueado
				//Ver si le ha dado like a la discoteca
				liked := isFavoriteBy(discoteca,myUserModel)
				discoteca.Liked=liked;

				//Devolvemos la discoteca, el total de likes, si le ha dado like el user y los eventos de esa discoteca
				c.JSON(http.StatusOK, gin.H{"discoteca":discoteca})
				return
			
			}else{//No hay usuario logueado
				c.JSON(http.StatusOK, gin.H{"discoteca": discoteca})
				return
			}
			
		}
		
	}
}
////////

//Get favorites count of a discoteca
func FavoritesCount(c *gin.Context) {  //Podemos pasarle bearer o no, si no lo pasamos es porque no está logueado entonces no marca si 
	var discoteca Discotecas		   //le ha dado like o no. Si le pasamos bearer es que está logueado, entonces nos dirá si le ha dado
	id := c.Params.ByName("id")		   //like o no le ha dado
	

	err := GetDiscotecaById(&discoteca, id)

	if err != nil{//Discoteca no valida
		c.JSON(http.StatusNotFound, "Discoteca Not found")
	 	c.AbortWithStatus(http.StatusNotFound)
	}else{

		myUserModel := c.MustGet("my_user_model").(User) //Usuario que da like
		count := favoritesCount(discoteca)

		if len(myUserModel.Username) > 1{//Hay usuario logueado

			//Ver si le ha dado like a la discoteca
			liked := isFavoriteBy(discoteca,myUserModel)

			//Devolvemos el total de likes y si le ha dado like el user
			c.JSON(http.StatusOK, gin.H{"Total": count, "liked":liked})
		
		}else{//No hay usuario logueado, devolvemos el count de likes
			c.JSON(http.StatusOK, gin.H{"Total": count})
		}
		
	}
}

//UPDATE discoteca

func DiscotecaUpdate(c *gin.Context){
	var discoteca Discotecas
	var newDiscoteca Discotecas
	c.BindJSON(&newDiscoteca);  //Aqui en teoria está la discoteca que le hemos pasado por postman

	id := c.Params.ByName("id")
	err := GetDiscotecaById(&discoteca, id) 

	discoteca.Name = newDiscoteca.Name


	if err != nil { 
		c.JSON(http.StatusNotFound, "NOT FOUND")
	}else{ 
		c.BindJSON(&discoteca)
		err = UpdateDiscoteca(&discoteca)//&discoteca  Aqui hay que meterle la discoteca nueva, con el c.BingJSON pero no me hace el json de la nueva
		if err != nil {
			c.JSON(http.StatusOK, "Not found")
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"discoteca": discoteca})
			return
		}
	}

}


//DELETE discoteca
func DiscotecaDelete(c *gin.Context){
	var discoteca Discotecas
	id := c.Params.ByName("id")

	err := DeleteDiscoteca(&discoteca, id)

	if err != nil{
		c.JSON(http.StatusNotFound, common.NewError("discotecas", errors.New("Invalid id")))
		return
	}
	
	//Borrar discoreca de redis
	client := common.NewClient()
	err2 :=  common.DeleteDiscoRedis(id,client)

	if err2 != nil{
		fmt.Println("Error deleting redis");
	}


	c.JSON(http.StatusOK, gin.H{"discoteca": "Delete Discoteca"})

}


//////Favorite
func DiscotecaFavorite(c *gin.Context){
	id := c.Params.ByName("id")
	var discoteca Discotecas

	err := GetDiscotecaById(&discoteca, id)//Discoteca a la que le damos like

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("discotecas", errors.New("Invalid id")))
		return
	}

	myUserModel := c.MustGet("my_user_model").(User) //Usuario que da like


	err2 := favoriteBy(myUserModel,discoteca)
	if err2 != nil{
		fmt.Println("ERROR like: ",err2)
	}

	//Enviar a redis los datos del user y la discoteca que da like
	
	client := common.NewClient()
	err3 := common.SaveUserLike(myUserModel.ID, strconv.FormatUint(uint64(discoteca.Id), 10),client)

	if err3 != nil{
		fmt.Println("ERRor redis userlike ",err3)
	}

	c.JSON(http.StatusOK, gin.H{ "User":myUserModel,"Disco":discoteca})
}

//UNFAVORITE
func DiscotecaUnFavorite(c *gin.Context){
	id := c.Params.ByName("id")
	var discoteca Discotecas

	err := GetDiscotecaById(&discoteca, id)//Discoteca a la que le damos like

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("discotecas", errors.New("Invalid id")))
		return
	}

	myUserModel := c.MustGet("my_user_model").(User) //Usuario que da like

	err2 := unFavoriteBy(myUserModel,discoteca)
	if err2 != nil{
		fmt.Println("ERROR like: ",err2)
	}

	c.JSON(http.StatusOK, gin.H{ "User":myUserModel,"Disco":discoteca})
}

//Get de las discotecas creadas por un usuario
func DiscotecasUser(c *gin.Context){
	id := c.Params.ByName("id") //Id del usuario que queremos coger sus discotecas
	var discotecas []Discotecas;

	user_id, err := strconv.ParseUint(id, 10, 32)

	if err != nil{
		fmt.Println("Error parse uint64 ", err)
	}

	myUserModel := c.MustGet("my_user_model").(User)
	my_User_id := uint64(myUserModel.ID)

	if user_id == my_User_id {

		err2 := GetDiscotecaByUser(user_id, &discotecas) //Discotecas del user

		if err2 != nil{
			fmt.Println("ERROR get user discotecas ", err2)
		}else{
			c.JSON(http.StatusOK, gin.H{ "discotecas":discotecas})
		}

	}else{
		fmt.Println("Los ids no son iguales, OK")
		fmt.Println("ID que le paso en la url: ",user_id)
		fmt.Println("ID del myusermodel: ",my_User_id)
		c.JSON(http.StatusNotFound,  "Error id user")

	}
}