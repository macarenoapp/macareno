package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"goApp/common"
	"github.com/jinzhu/gorm"
	"goApp/src"
	"github.com/gin-contrib/cors"
	"time"
)

// "goApp/discotecas"
// "goApp/common"

func Migrate(db *gorm.DB) {

	db.AutoMigrate(&discotecas.Discotecas{})
	discotecas.AutoMigrate()

}

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	// MakeRoutes(r)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))


	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:4200"}
	// config.AddAllowHeaders("*")
	// r.Use(cors.New(config))
	// r := MakeRoutes()
	v1 := r.Group("/api")
	
	v1.Use(discotecas.AuthMiddleware(false))
	discotecas.DiscotecasAnonymousRegister(v1.Group("/discotecas"))

	v1.Use(discotecas.AuthMiddleware(true))
	discotecas.DiscotecasRegister(v1.Group("/discotecas"))

	fmt.Printf("0.0.0.0:3000")
	r.Run(":8080")//Cambiar al 8080 para traefik // listen and serve on 0.0.0.0:8080 by default
}

// func MakeRoutes(r *gin.Engine) {

// 	cors := func(c *gin.Context) {

// 		fmt.Println("CONSUlTA")
// 		// c.Request.Method = "POST"
// 		fmt.Println(c.Request.Method)

// 		fmt.Printf("c.Request.Method \n")

// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
// 		c.Writer.Header().Set("Content-Type", "application/json")

// 		if c.Request.Method == "OPTIONS" {
// 			fmt.Println("OPTIONs??????????????")
// 			c.AbortWithStatus(200)
// 		}
// 		c.Next()
// 	}
// 	r.Use(cors)
// }
