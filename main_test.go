package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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
	req, _ := http.NewRequest("GET", "/warranty/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"message":"id not found: 2"}`, w.Body.String())
}

func TestGetWarranty(t *testing.T) {
	router := initEngine()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/warranty/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, strings.Contains(w.Body.String(), "Samsung"), true)
}

func TestCreateWarranty(t *testing.T) {
	router := initEngine()
	warranty := models.Warranty{ID: uuid.UUID{}, BrandName: "Samsung", StoreName: "Costco",
		TransactionTime: time.Now(), ExpiryTime: time.Now().Add(5 * time.Second * 86400),
		Amount: 100000}
	jsonValue, _ := json.Marshal(warranty)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/warranty", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	retW := &models.Warranty{}
	json.Unmarshal(w.Body.Bytes(), retW)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Samsung", retW.BrandName)
	assert.NotEqual(t, warranty.ID, retW.ID)
}

func initEngine() *gin.Engine {
	db := models.Init()
	router := setupRouter(db)
	return router
}
