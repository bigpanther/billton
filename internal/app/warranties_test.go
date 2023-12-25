package app_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bigpanther/billton/internal/app"
	"github.com/bigpanther/billton/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetWarrantyNotFound(t *testing.T) {
	router := setupTestRoutes("test-warranties", app.WarrantiesRoutes)
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test-warranties/non-existent-id", nil)
	require.Nil(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"message":"Warranty id not found: non-existent-id"}`, w.Body.String())
}

func TestCreateAndGetWarranty(t *testing.T) {
	router := setupTestRoutes("test-warranties", app.WarrantiesRoutes)

	user := models.User{ID: uuid.UUID{}, Name: "Tim Hortons"}
	db := app.InitDb()
	tx := db.Create(&user)
	require.Nil(t, tx.Error)
	//require.False(t, verrs.HasAny(), verrs.Error())

	warranty := models.Warranty{ID: uuid.UUID{}, BrandName: "Samsung", StoreName: "Costco",
		TransactionTime: time.Now(), ExpiryTime: time.Now().Add(5 * time.Second * 86400),
		Amount: 100000, UserID: user.ID}
	jsonValue, err := json.Marshal(warranty)
	require.Nil(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/test-warranties", bytes.NewBuffer(jsonValue))
	require.Nil(t, err)

	router.ServeHTTP(w, req)
	retW := &models.Warranty{}
	require.Equal(t, 200, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), retW)
	require.Nil(t, err)
	assert.Equal(t, warranty.BrandName, retW.BrandName)
	assert.Equal(t, warranty.UserID, retW.UserID)
	newID := retW.ID
	assert.NotEqual(t, warranty.ID, newID)

	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", fmt.Sprintf("/test-warranties/%s", newID), nil)
	require.Nil(t, err)

	router.ServeHTTP(w, req)
	retW = &models.Warranty{}
	require.Equal(t, 200, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), retW)
	require.Nil(t, err)

	assert.Equal(t, warranty.BrandName, retW.BrandName)
	assert.Equal(t, newID, retW.ID)
}

func TestEditWarranty(t *testing.T) {
	router := setupTestRoutes("test-warranties", app.WarrantiesRoutes)

	user := models.User{ID: uuid.UUID{}, Name: "Sam Smith"}
	db := app.InitDb()
	tx := db.Create(&user)
	require.Nil(t, tx.Error)
	//require.False(t, verrs.HasAny(), verrs.Error())

	warranty := models.Warranty{ID: uuid.UUID{}, BrandName: "LG", StoreName: "Walmart",
		TransactionTime: time.Now(), ExpiryTime: time.Now().Add(5 * time.Second * 86400),
		Amount: 10000, UserID: user.ID}
	jsonValue, err := json.Marshal(warranty)
	require.Nil(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/test-warranties", bytes.NewBuffer(jsonValue))
	require.Nil(t, err)

	router.ServeHTTP(w, req)
	retW := &models.Warranty{}
	require.Equal(t, 200, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), retW)
	require.Nil(t, err)

	assert.Equal(t, warranty.BrandName, retW.BrandName)
	newID := retW.ID
	assert.NotEqual(t, warranty.ID, newID)

	w = httptest.NewRecorder()
	warranty2 := models.Warranty{BrandName: "Samsung", StoreName: "Costco",
		TransactionTime: time.Now(), ExpiryTime: time.Now().Add(5 * time.Second * 86400),
		Amount: 1000, UserID: user.ID}
	jsonValue2, err := json.Marshal(warranty2)
	require.Nil(t, err)

	req, err = http.NewRequest("PUT", fmt.Sprintf("/test-warranties/%s", newID), bytes.NewBuffer(jsonValue2))
	require.Nil(t, err)

	router.ServeHTTP(w, req)
	retW = &models.Warranty{}
	require.Equal(t, 200, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), retW)
	require.Nil(t, err)

	assert.Equal(t, warranty2.BrandName, retW.BrandName)
	assert.Equal(t, newID, retW.ID)
}
func TestDeleteWarranty(t *testing.T) {
	router := setupTestRoutes("test-warranties", app.WarrantiesRoutes)

	user := models.User{ID: uuid.UUID{}, Name: "Sam Smith"}
	db := app.InitDb()
	tx := db.Create(&user)
	require.Nil(t, tx.Error)
	//require.False(t, verrs.HasAny(), verrs.Error())

	warranty := models.Warranty{ID: uuid.UUID{}, BrandName: "LG", StoreName: "Walmart",
		TransactionTime: time.Now(), ExpiryTime: time.Now().Add(5 * time.Second * 86400),
		Amount: 10000, UserID: user.ID}
	jsonValue, err := json.Marshal(warranty)
	require.Nil(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/test-warranties", bytes.NewBuffer(jsonValue))
	require.Nil(t, err)

	router.ServeHTTP(w, req)
	retW := &models.Warranty{}
	err = json.Unmarshal(w.Body.Bytes(), retW)
	require.Nil(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, warranty.BrandName, retW.BrandName)
	newID := retW.ID
	assert.NotEqual(t, warranty.ID, newID)

	w = httptest.NewRecorder()
	req, err = http.NewRequest("DELETE", fmt.Sprintf("/test-warranties/%s", newID), nil)
	require.Nil(t, err)

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	message := "{\"message\":" + fmt.Sprintf("\"Record successfully deleted: %s\"", newID) + "}"
	assert.Equal(t, message, w.Body.String())

}

func TestAddImage(t *testing.T) {
	router := setupTestRoutes("test-warranties", app.WarrantiesRoutes)

	user := models.User{ID: uuid.UUID{}, Name: "Sam Smith"}
	db := app.InitDb()
	tx := db.Create(&user)
	require.Nil(t, tx.Error)
	//require.False(t, verrs.HasAny(), verrs.Error())

	// Create a new Warranty for testing purposes
	warranty := models.Warranty{
		ID:              uuid.UUID{},
		BrandName:       "Samsung",
		StoreName:       "Costco",
		TransactionTime: time.Now(),
		ExpiryTime:      time.Now().Add(5 * time.Second * 86400),
		Amount:          100000,
		UserID:          user.ID,
	}
	// Save the warranty to the database so we can retrieve it later

	tx = db.Create(&warranty)
	require.Nil(t, tx.Error)
	//require.False(t, verrs.HasAny(), verrs.Error())

	// Create a new test file
	fileBuf := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(fileBuf)
	part, err := multipartWriter.CreateFormFile("file", "test.jpg")
	require.Nil(t, err)
	_, err = part.Write([]byte("test image data"))
	require.Nil(t, err)

	multipartWriter.Close()

	// Prepare the HTTP request
	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/test-warranties/%s/upload_image", warranty.ID), fileBuf)
	require.NoError(t, err)
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, fmt.Sprintf("'%s' uploaded!", "test.jpg"), w.Body.String())
}
