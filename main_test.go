package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"

	"github.com/bigpanther/warrant/models"
	"github.com/stretchr/testify/assert"
)

func TestGetWarrantyNotFound(t *testing.T) {
	router := initEngine()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/warranties/non-existent-id", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"message":"id not found: non-existent-id"}`, w.Body.String())
}

func TestCreateAndGetWarranty(t *testing.T) {
	router := initEngine()
	warranty := models.Warranty{ID: uuid.UUID{}, BrandName: "Samsung", StoreName: "Costco",
		TransactionTime: time.Now(), ExpiryTime: time.Now().Add(5 * time.Second * 86400),
		Amount: 100000}
	jsonValue, _ := json.Marshal(warranty)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/warranties", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	retW := &models.Warranty{}
	json.Unmarshal(w.Body.Bytes(), retW)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, warranty.BrandName, retW.BrandName)
	newID := retW.ID
	assert.NotEqual(t, warranty.ID, newID)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", fmt.Sprintf("/warranties/%s", newID), nil)
	router.ServeHTTP(w, req)
	retW = &models.Warranty{}
	json.Unmarshal(w.Body.Bytes(), retW)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, warranty.BrandName, retW.BrandName)
	assert.Equal(t, newID, retW.ID)
}

func TestAddImage(t *testing.T) {
	router := initEngine()

	// Create a new Warranty for testing purposes
	warranty := models.Warranty{
		ID:              uuid.UUID{},
		BrandName:       "Samsung",
		StoreName:       "Costco",
		TransactionTime: time.Now(),
		ExpiryTime:      time.Now().Add(5 * time.Second * 86400),
		Amount:          100000,
	}
	// Save the warranty to the database so we can retrieve it later
	db := models.Init()
	verrs, err := db.ValidateAndCreate(&warranty)
	assert.NoError(t, err)
	assert.False(t, verrs.HasAny())

	// Create a new test file
	fileBuf := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(fileBuf)
	part, err := multipartWriter.CreateFormFile("file", "test.jpg")
	assert.NoError(t, err)
	part.Write([]byte("test image data"))
	multipartWriter.Close()

	// Prepare the HTTP request
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/warranties/%s/upload", warranty.ID), fileBuf)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, fmt.Sprintf("'%s' uploaded!", "test.jpg"), w.Body.String())
}

func TestEditWarranty(t *testing.T) {
	router := initEngine()
	warranty := models.Warranty{ID: uuid.UUID{}, BrandName: "LG", StoreName: "Walmart",
		TransactionTime: time.Now(), ExpiryTime: time.Now().Add(5 * time.Second * 86400),
		Amount: 10000}
	jsonValue, _ := json.Marshal(warranty)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/warranties", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	retW := &models.Warranty{}
	json.Unmarshal(w.Body.Bytes(), retW)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, warranty.BrandName, retW.BrandName)
	newID := retW.ID
	assert.NotEqual(t, warranty.ID, newID)

	w = httptest.NewRecorder()
	warranty2 := models.Warranty{BrandName: "Samsung", StoreName: "Costco",
		TransactionTime: time.Now(), ExpiryTime: time.Now().Add(5 * time.Second * 86400),
		Amount: 1000}
	jsonValue2, _ := json.Marshal(warranty2)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/warranties/%s", newID), bytes.NewBuffer(jsonValue2))
	router.ServeHTTP(w, req)
	retW = &models.Warranty{}
	json.Unmarshal(w.Body.Bytes(), retW)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, warranty2.BrandName, retW.BrandName)
	assert.Equal(t, newID, retW.ID)
}
func TestDeleteWarranty(t *testing.T) {
	router := initEngine()
	warranty := models.Warranty{ID: uuid.UUID{}, BrandName: "LG", StoreName: "Walmart",
		TransactionTime: time.Now(), ExpiryTime: time.Now().Add(5 * time.Second * 86400),
		Amount: 10000}
	jsonValue, _ := json.Marshal(warranty)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/warranties", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	retW := &models.Warranty{}
	json.Unmarshal(w.Body.Bytes(), retW)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, warranty.BrandName, retW.BrandName)
	newID := retW.ID
	assert.NotEqual(t, warranty.ID, newID)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/warranties/%s", newID), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	message := "{\"message\":" + fmt.Sprintf("\"Record successfully deleted: %s\"", newID) + "}"
	assert.Equal(t, message, w.Body.String())

}

func initEngine() *gin.Engine {
	db := models.Init()
	router := setupRouter(db)
	return router
}
