package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bigpanther/billton/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WarrantiesRoutes sets up the routes for managing warranties
func WarrantiesRoutes(rg *gin.RouterGroup) {
	rg.POST("", createWarranty)
	rg.GET("/:id", warrantyByID)
	rg.PUT("/:id", editWarranty)
	rg.DELETE("/:id", deleteWarranty)
	rg.POST("/:id/upload_image", addImage)
	rg.GET("/:id/image", getImage)
}

func warrantyByID(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)
	id := c.Params.ByName("id")
	w := &models.Warranty{}
	tx := db.Find(w, "id = ?", id)
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Warranty id not found: %s", id),
		})
		log.Println(tx.Error)
		return
	}
	c.IndentedJSON(http.StatusOK, w)
}

func deleteWarranty(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	id := c.Params.ByName("id")
	w := &models.Warranty{}
	tx := db.Find(w, "id = ?", id)
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id not found: %s", id),
		})
		log.Println(tx.Error)
		return
	}
	tx = db.Delete(w)
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Record cannot be deleted: %s", id),
		})
		log.Println(tx.Error)
		return
	}
	//fmt.Println("Record successfully deleted")

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Record successfully deleted: %s", id),
	})
}

func editWarranty(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	id := c.Params.ByName("id")
	w := &models.Warranty{}
	tx := db.Find(w, "id = ?", id)
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id not found: %s", id),
		})
		log.Println(tx.Error)
		return
	}
	w2 := &models.Warranty{}
	err := c.BindJSON(w2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error updating warranty",
		})
		log.Println(err)
		return
	}
	w.BrandName = w2.BrandName
	w.StoreName = w2.StoreName
	w.Amount = w2.Amount
	w.TransactionTime = w2.TransactionTime
	w.UserID = w2.UserID
	tx = db.Save(w)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error saving warranty",
		})
		log.Println(err)
		return
	}
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": tx.Error.Error(),
		})
		log.Println(tx.Error.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, w)

}

func createWarranty(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

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
	tx := db.Create(w)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error saving warranty",
		})
		log.Println(tx.Error)
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

func getImage(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	w := &models.Warranty{}
	download_path := "warranty_receipts/"
	id := c.Params.ByName("id")
	tx := db.Find(w, "id = ?", id)
	var targetPath string
	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id not found: %s", id),
		})
		log.Println(tx.Error)
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

func addImage(c *gin.Context) {
	db := c.Value("db").(*gorm.DB)

	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	id := c.Params.ByName("id")
	w := &models.Warranty{}
	tx := db.Find(w, "id = ?", id)

	if tx.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id not found: %s", id),
		})
		return
	}
	fileName := string(file.Filename)
	if !isValidFilename(fileName) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error: You can only upload a JPEG or PNG files",
		})
		return
	}
	dst := "warranty_receipts/" + id + filepath.Ext(fileName)

	// Upload the file to specific dst.
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("error saving image: %v", err),
		})
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
