package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/thingy", getThingies)
		v1.GET("/thingy/:id", getThingyById)
		v1.PUT("/thingy", putThingy)
		v1.POST("/thingy/name/:name", newThingy)
		v1.DELETE("/thingy/:id", deleteThingy)
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}

func getThingies(c *gin.Context) {
	queryOffset := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(queryOffset)
	if err != nil {
		offset = 0
	}
	if offset < 0 || offset >= len(thingiesDB) {
		offset = 0
	}
	if offset+thingiesPerPage > len(thingiesDB) {
		c.JSON(http.StatusOK, thingiesDB[offset:])
		return
	}

	c.JSON(http.StatusOK, thingiesDB[offset:offset+thingiesPerPage])
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

func newThingy(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thingy", "data": nil})
		return
	}

	t := Thingy{Id: ulid.Make(), Name: name}
	thingiesDB = append(thingiesDB, t)
	c.JSON(http.StatusAccepted, gin.H{"err": nil, "data": t})
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

const thingiesPerPage int = 10

var thingiesDB = addThingiesToDB(34)

func addThingiesToDB(n int) []Thingy {
	var thingies = []Thingy{}
	if n <= 0 {
		n = 5
	}
	for i := 0; i < n; i++ { // TODO: Update to Go 1.22 to rante over integer
		thingies = append(thingies, Thingy{Id: ulid.Make(), Name: fmt.Sprintf("Thingy-%d", i)})
	}
	return thingies
}
