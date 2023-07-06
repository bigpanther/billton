package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bigpanther/warrant/warranty"
	"github.com/gin-gonic/gin"
)

// Dummy database
var warranties = map[string]warranty.Warranty{

	"1": {1, time.Now(), time.Now().Add(5 * time.Second * 86400), "Samsung", 10000},
}

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/warranty/:id", warrantyByID)
	r.POST("/warranty", createWarranty)
	return r
}

func warrantyByID(c *gin.Context) {
	id := c.Params.ByName("id")
	w, ok := warranties[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id not found: %s", id),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, w)
}

func createWarranty(c *gin.Context) {
	w := warranty.Warranty{}
	err := c.BindJSON(&w)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error creating warranty",
		})
		return
	}
	// TODO: put it in actual DB
	warranties[fmt.Sprint(w.ID)] = w
	c.IndentedJSON(http.StatusOK, w)
}
