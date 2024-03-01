package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/thingy", getThingies)
	}
	r.Run()
}

func getThingies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
