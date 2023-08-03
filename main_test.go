package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func initEngine() *gin.Engine {
	db := models.Init()
	router := setupRouter(db)
	return router
}
