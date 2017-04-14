package boot_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/cmd/webapp/boot"
)

// TestRoutes ensures each of the routes is set up properly.
func TestRoutes(t *testing.T) {
	// Register the services and load the routes.
	h := boot.ServicesAndRoutes("../html")

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
	AssertEqual(t, w.Code, http.StatusOK)

	// Test a 404.
	w = httptest.NewRecorder()
	r, err = http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatal(err)
	}
	h.ServeHTTP(w, r)
	AssertEqual(t, w.Code, http.StatusNotFound)

	// Test the register page.
	w = httptest.NewRecorder()
	r, err = http.NewRequest("GET", "/register", nil)
	if err != nil {
		t.Fatal(err)
	}
	h.ServeHTTP(w, r)
	AssertEqual(t, w.Code, http.StatusOK)
}
