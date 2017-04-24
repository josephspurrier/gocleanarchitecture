package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRoutes ensures each of the routes is set up properly.
func TestRoutes(t *testing.T) {
	// Register the services and load the routes.
	h := RegisterServices("html").LoadRoutes()

	var err error
	var w *httptest.ResponseRecorder
	var r *http.Request

	// Test the login page.
	w = httptest.NewRecorder()
	r, err = http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	h.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test a 404.
	w = httptest.NewRecorder()
	r, err = http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatal(err)
	}
	h.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Test the register page.
	w = httptest.NewRecorder()
	r, err = http.NewRequest("GET", "/register", nil)
	if err != nil {
		t.Fatal(err)
	}
	h.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
}
