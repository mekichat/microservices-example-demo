package main

import (
	"encoding/json"
    "net/http"
	"github.com/gin-gonic/gin"
)



func main()  {

	router := gin.Default() // deafult middleware

	router.GET("/process", func(ctx *gin.Context) {

		resp, err := http.Get("http://localhost:8083/data")

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Could not reach service3",		
		    })
		    return
		}
		
		defer resp.Body.Close()
		var data map[string] interface{}

		json.NewDecoder(resp.Body).Decode(&data)
		data["service2"] = "processed by service 2"
		ctx.JSON(200, data)
	})	

	router.Run(":8082")
}