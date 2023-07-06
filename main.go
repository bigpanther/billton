package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bigpanther/warrant/context"
	"github.com/bigpanther/warrant/models"
	warranty "github.com/bigpanther/warrant/models"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

// Dummy database
var warranties = map[string]warranty.Warranty{

	"1": {ID: uuid.UUID{}, TransactionTime: time.Now(), ExpiryTime: time.Now().Add(5 * time.Second * 86400), BrandName: "Samsung", StoreName: "Costco", Amount: 10000},
}

func main() {
	db := models.Init()
	r := setupRouter(db)
	r.Run()
}

func setupRouter(db *pop.Connection) *gin.Engine {
	r := gin.Default()
	r.GET("/warranty/:id", withDb(db, warrantyByID))
	r.POST("/warranty", withDb(db, createWarranty))
	return r
}

func withDb(db *pop.Connection, f func(c *context.AppContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := &context.AppContext{Context: c, DB: db}
		f(a)
	}
}

func warrantyByID(c *context.AppContext) {
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

func createWarranty(c *context.AppContext) {
	w := &warranty.Warranty{}
	err := c.BindJSON(w)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error creating warranty",
		})
		return
	}
	verrs, err := c.DB.ValidateAndCreate(w)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error saving warranty",
		})
		return
	}
	if verrs.HasAny() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": verrs,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, w)
}
