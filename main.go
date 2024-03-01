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
		v1.PUT("/thingy", putThingy)
		v1.DELETE("/thingy/:id", deleteThingy)
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
	c.JSON(http.StatusNotFound, gin.H{"data": nil, "error": "thingy not found"})
}

func putThingy(c *gin.Context) {
	var t Thingy
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid thingy", "data": nil})
		return
	}
	tId, err := addThingyToDB(&t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil, "err": "error processing thingy"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"data": tId, "err": nil})
}

func addThingyToDB(t *Thingy) (string, error) {
	thingiesDB = append(thingiesDB, Thingy{Id: t.Id, Name: t.Name})
	return t.Id.String(), nil
}

func deleteThingy(c *gin.Context) {
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
	for i, t := range thingiesDB {
		if tUlid == t.Id {
			thingiesDB[i] = thingiesDB[len(thingiesDB)-1]
			thingiesDB = thingiesDB[:len(thingiesDB)-1]
			c.JSON(http.StatusAccepted, gin.H{"data": t.Id, "error": nil})
			return
		}
	}
	c.JSON(http.StatusGone, gin.H{"data": nil, "err": "thingy not found"})
}

type Thingy struct {
	Id   ulid.ULID
	Name string
}

var thingiesDB = []Thingy{
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "someThingy"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "otherThingy"},
}
