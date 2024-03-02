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

// []Thingy{
// 	{Id: ulid.MustParse("01HQXNGNKG5XZ5WCGAC2D318BQ"), Name: "Thingy-1"},
// 	{Id: ulid.MustParse("01HQXNJ69R615NEH3ZVWN522M9"), Name: "Thingy-2"},
// 	{Id: ulid.MustParse("01HQXXHJXPT6AW4GCCV20VBJN4"), Name: "Thingy-3"},
// 	{Id: ulid.MustParse("01HQXXHXEKEJ9SNH2JZ20KZRKP"), Name: "Thingy-4"},
// 	{Id: ulid.MustParse("01HQXXHSDQCWDZME3J67T4BD68"), Name: "Thingy-5"},
// 	{Id: ulid.MustParse("01HQXXK79CK7KK7M4K3MM9HY0Z"), Name: "Thingy-6"},
// 	{Id: ulid.MustParse("01HQXXKDEQTXPRQDR84SWVSDFD"), Name: "Thingy-7"},
// 	{Id: ulid.MustParse("01HQXXKM1GN1XZYKRX08CBA4QP"), Name: "Thingy-8"},
// 	{Id: ulid.MustParse("01HQXXKSWGG4SFGGMFKF2SHW2E"), Name: "Thingy-9"},
// 	{Id: ulid.MustParse("01HQXXKYTRKSDEA5RS983CV0WQ"), Name: "Thingy-10"},
// 	{Id: ulid.MustParse("01HQXXM33V8D11N2WESF6HFYJS"), Name: "Thingy-11"},
// 	{Id: ulid.MustParse("01HQXXM7ZQAQBPCY7PNNZ4J1J7"), Name: "Thingy-12"},
// 	{Id: ulid.MustParse("01HQXXMCXWWA9ZYBPB0XJPHTQY"), Name: "Thingy-13"},
// 	{Id: ulid.MustParse("01HQXXMHZE2HEPX9HCE2A15CDG"), Name: "Thingy-14"},
// 	{Id: ulid.MustParse("01HQXXMR6W3MYTGCYGJQGG186P"), Name: "Thingy-15"},
// 	{Id: ulid.MustParse("01HQXXMW4VBSRR2N2466DPH26W"), Name: "Thingy-16"},
// 	{Id: ulid.MustParse("01HQXXN0BH0MKYCNCP56YE9DGV"), Name: "Thingy-17"},
// 	{Id: ulid.MustParse("01HQXXN4V7MH31N21V9ANH1DMA"), Name: "Thingy-18"},
// 	{Id: ulid.MustParse("01HQXXN9PA8G8MGF5RB0W55NSP"), Name: "Thingy-19"},
// 	{Id: ulid.MustParse("01HQXXNDMMCFJN4VVT2B89T9XY"), Name: "Thingy-20"},
// 	{Id: ulid.MustParse("01HQXXNHJFXNB2CXZF8TMVF1GH"), Name: "Thingy-21"},
// 	{Id: ulid.MustParse("01HQXXNPP2V6EPBP2GYYAP0D85"), Name: "Thingy-22"},
// 	{Id: ulid.MustParse("01HQXXNV2ZHH1GAPEXGFYP1WZ1"), Name: "Thingy-23"},
// 	{Id: ulid.MustParse("01HQXXNZ4VEEEQ62THVEQFP1GS"), Name: "Thingy-24"},
// 	{Id: ulid.MustParse("01HQXXP72Q66NRSH2NC6R851PV"), Name: "Thingy-25"},
// 	{Id: ulid.MustParse("01HQXXPMVZ71Y26P70EG1GED1A"), Name: "Thingy-26"},
// 	{Id: ulid.MustParse("01HQXXPT4Z3WST85VJM5F1VAXJ"), Name: "Thingy-27"},
// 	{Id: ulid.MustParse("01HQXXQ0AAPNXRT3TGTWFH89VK"), Name: "Thingy-28"},
// 	{Id: ulid.MustParse("01HQXXQ4HCPJCKP2X5KHCT26XC"), Name: "Thingy-29"},
// 	{Id: ulid.MustParse("01HQXXQ9EZ7K3HY00SE4AWDEC3"), Name: "Thingy-30"},
// 	{Id: ulid.MustParse("01HQXXQFVMV73JG88PRRA7X147"), Name: "Thingy-31"},
// 	{Id: ulid.MustParse("01HQXXQMZ0M4BPEHHJRYXRS1E9"), Name: "Thingy-32"},
// }

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
