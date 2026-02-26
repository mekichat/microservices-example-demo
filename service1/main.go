package main

import (
	"encoding/json"
    "net/http"
	"github.com/gin-gonic/gin"
)


func main()  {
	router := gin.Default() // deafult maluware

	router.GET("/start", func(ctx *gin.Context) {

		resp, err := http.Get("http://localhost:8082/process")

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": "Could not react service2",		
		    })
		    return
		}
		
		defer resp.Body.Close()
		var data map[string] interface{}
		json.NewDecoder(resp.Body).Decode(&data)
		data["service1"] = "Handled by service 1"
		ctx.JSON(200, data)

	})	

	router.Run(":8081")
}