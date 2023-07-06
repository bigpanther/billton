package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bigpanther/warrant/warranty"
	"gotest.tools/v3/assert"
)

func TestGetWarrantyNotFound(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/warranty/2", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"message":"id not found: 2"}`, w.Body.String())
}

func TestGetWarranty(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/warranty/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, strings.Contains(w.Body.String(), "Samsung"), true)
}

func TestCreateWarranty(t *testing.T) {
	router := setupRouter()
	warranty := warranty.Warranty{ID: 2, Brand: "NotOppo"}
	jsonValue, _ := json.Marshal(warranty)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/warranty", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
