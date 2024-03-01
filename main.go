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
	}
	r.Run()
}

func getThingies(c *gin.Context) {
	c.JSON(http.StatusOK, thingiesDB)
}

type Thingy struct {
	Id   ulid.ULID
	Name string
}

var thingiesDB = []Thingy{
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "someThingy"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "otherThingy"},
}
