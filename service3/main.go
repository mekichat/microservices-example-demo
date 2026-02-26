package main

import (
		
	"github.com/gin-gonic/gin"
)



func main()  {

	router := gin.Default() // deafult maluware

	router.GET("/data", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"service": "service3",
			"message": "Hello from service 3",
		})

	})

	router.Run(":8083")

}