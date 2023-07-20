package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bigpanther/warrant/context"
	"github.com/bigpanther/warrant/models"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

func main() {
	db := models.Init()
	r := setupRouter(db)
	r.Run()
}

func setupRouter(db *pop.Connection) *gin.Engine {
	r := gin.Default()
	r.GET("/warranties/:id", withDb(db, warrantyByID))
	r.POST("/warranties", withDb(db, createWarranty))
	r.POST("/users", withDb(db, createUser))
	r.PUT("/warranties/:id", withDb(db, editWarranty))
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
	w := &models.Warranty{}
	err := c.DB.Find(w, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id not found: %s", id),
		})
		log.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, w)
}

func editWarranty(c *context.AppContext) {
	id := c.Params.ByName("id")
	w := &models.Warranty{}
	err := c.DB.Find(w, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id not found: %s", id),
		})
		log.Println(err)
		return
	}
	err2 := c.BindJSON(w)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error updating warranty",
		})
		log.Println(err)
		return
	}
	verrs, err := c.DB.ValidateAndUpdate(w)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error saving warranty",
		})
		log.Println(err)
		return
	}
	if verrs.HasAny() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": verrs,
		})
		log.Println(verrs)
		return
	}
	c.IndentedJSON(http.StatusOK, w)

}

func createWarranty(c *context.AppContext) {
	w := &models.Warranty{}
	err := c.BindJSON(w)
	w.ID = uuid.UUID{}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error creating warranty",
		})
		log.Println(err)
		return
	}
	verrs, err := c.DB.ValidateAndCreate(w, "ID")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error saving warranty",
		})
		log.Println(err)
		return
	}
	if verrs.HasAny() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": verrs,
		})
		log.Println(verrs)
		return
	}
	c.IndentedJSON(http.StatusOK, w)
}

func createUser(c *context.AppContext) {
	u := &models.User{}
	err := c.BindJSON(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error creating user",
		})
		log.Println(err)
		return
	}
	verrs, err := c.DB.ValidateAndCreate(u, "ID")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error saving user",
		})
		log.Println(err)

		return
	}
	if verrs.HasAny() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": verrs,
		})
		log.Println(verrs)
		return
	}
	c.IndentedJSON(http.StatusOK, u)
}
