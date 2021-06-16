package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"goApp/common"
	"github.com/jinzhu/gorm"
	"goApp/src"
	
)

// "goApp/entradas"
// "goApp/common"

func Migrate(db *gorm.DB) {

	db.AutoMigrate(&entradas.Entradas{})
	// entradas.AutoMigrate()

}

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	MakeRoutes(r)
	v1 := r.Group("/api")
	
	v1.Use(entradas.AuthMiddleware(false))
	entradas.EntradasAnonymousRegister(v1.Group("/entradas"))

	v1.Use(entradas.AuthMiddleware(true))
	entradas.EntradasRegister(v1.Group("/entradas"))

	fmt.Printf("0.0.0.0:3000")
	r.Run(":8080")//Cambiar al 8080 para traefik // listen and serve on 0.0.0.0:8080 by default
}

func MakeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		fmt.Printf("c.Request.Method \n")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()
	}
	r.Use(cors)
}
