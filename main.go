package main

import (
	"net/http"
	"strconv"

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

var thingiesDB = []Thingy{
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-1"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-2"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-3"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-4"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-5"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-6"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-7"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-8"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-9"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-10"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-11"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-12"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-13"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-14"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-15"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-16"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-17"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-18"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-19"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-20"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-21"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-22"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-23"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-24"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-25"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-26"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-27"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-28"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-29"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-30"},
	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-31"},
	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-32"},
}
