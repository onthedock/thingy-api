package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
)

func main() {

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/thingy", getThingies)
		v1.GET("/thingy/:id", getThingyById)
	}
	r.Run()
}

func getThingies(c *gin.Context) {
	c.JSON(http.StatusOK, thingiesDB)
}

func getThingyById(c *gin.Context) {
	tId := c.Param("id")
	if tId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing thingy id", "data": nil})
		return
	}
	tUlid, err := ulid.Parse(tId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thingy id", "data": nil})
		return
	}
	for _, t := range thingiesDB {
		if tUlid == t.Id {
			c.JSON(http.StatusOK, gin.H{"data": t, "error": nil})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": ""})
}

type Thingy struct {
	Id   ulid.ULID
	Name string
}

var thingiesDB = []Thingy{
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "someThingy"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "otherThingy"},
}
