package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	r.POST("/warranties/:id/upload", withDb(db, addImage))
	r.GET("/warranties/:id/download", withDb(db, getImage))

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
func isValidFilename(name string) bool {
	if strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".jpeg") || strings.HasSuffix(name, ".png") {
		return true
	}
	return false
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

func getImage(c *context.AppContext) {
	w := &models.Warranty{}
	download_path := "warranty_receipts/"
	id := c.Params.ByName("id")
	err := c.DB.Find(w, id)
	var targetPath string
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id not found: %s", id),
		})
		log.Println(err)
		return
	}
	_, error := os.Stat(string(download_path + id + ".jpg"))
	_, error1 := os.Stat(string(download_path + id + ".jpeg"))
	_, error2 := os.Stat(string(download_path + id + ".png"))
	// check if error is "file not exists"
	if os.IsNotExist(error) && os.IsNotExist(error1) && os.IsNotExist(error2) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("file does not exist for id: %s", id),
		})
		return
	}
	if !os.IsNotExist(error) {
		targetPath = download_path + id + ".jpg"
	}
	if !os.IsNotExist(error1) {
		targetPath = download_path + id + ".jpeg"
	}
	if !os.IsNotExist(error2) {
		targetPath = download_path + id + ".png"
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+id)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)

}
func addImage(c *context.AppContext) {

	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	id := c.Params.ByName("id")
	w := &models.Warranty{}
	err := c.DB.Find(w, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id not found: %s", id),
		})
		return
	}
	fileName := string(file.Filename)
	if isValidFilename(fileName) {
		dst := "warranty_receipts/" + id + filepath.Ext(fileName)

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
