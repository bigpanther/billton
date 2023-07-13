package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bigpanther/warrant/context"
	"github.com/bigpanther/warrant/models"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop/v6"
)

func main() {
	db := models.Init()
	r := setupRouter(db)
	r.Run()
}

func setupRouter(db *pop.Connection) *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.GET("/warranties/:id", withDb(db, warrantyByID))
	r.POST("/warranties", withDb(db, createWarranty))
	r.POST("/warranty/:id/upload", withDb(db, addImage))

	r.POST("/users", withDb(db, createUser))
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

func createWarranty(c *context.AppContext) {
	w := &models.Warranty{}
	err := c.BindJSON(w)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error creating warranty",
		})
		log.Println(err)
		return
	}
	verrs, err := c.DB.ValidateAndCreate(w)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
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
	verrs, err := c.DB.ValidateAndCreate(u)
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

func addImage(c *context.AppContext) {

	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	id := c.Params.ByName("id")
	w := &warranty.Warranty{}
	err := c.DB.Find(w, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id not found: %s", id),
		})
		return
	}
	fileName := string(file.Filename)
	if fileName[len(fileName)-3:] == "jpg" || fileName[len(fileName)-4:] == "jpeg" || fileName[len(fileName)-3:] == "png" {
		dst := "warranty_receipts/" + string(id) + ".jpg"

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Error: You can only upload a JPEG or PNG files"),
		})
		return
	}

}
