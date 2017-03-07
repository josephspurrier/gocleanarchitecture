package register_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/controller/register"
	"github.com/josephspurrier/gocleanarchitecture/database"
	"github.com/josephspurrier/gocleanarchitecture/lib/view"
)

// AssertEqual throws an error if the two values are not equal.
func AssertEqual(t *testing.T, actualValue interface{}, expectedValue interface{}) {
	if actualValue != expectedValue {
		t.Errorf("\n got: %v\nwant: %v", actualValue, expectedValue)
	}
}

// TestIndex ensures the index function returns a 200 code.
func TestIndex(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(register.Handler)
	h.ViewService = view.New("../../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusOK)
}

// TestStoreCreateOK ensures register can be successful.
func TestStoreCreateOK(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the request body.
	val := url.Values{}
	r.Form = val
	r.Form.Add("firstname", "John")
	r.Form.Add("lastname", "Doe")
	r.Form.Add("email", "jdoe@example.com")
	r.Form.Add("password", "Pa$$w0rd")

	// Call the handler.
	h := new(register.Handler)
	db := new(database.MockService)
	h.UserService = database.NewUserService(db)
	h.ViewService = view.New("../../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusCreated)

	// Fail on duplicate creation.
	w = httptest.NewRecorder()
	h.Index(w, r)
	AssertEqual(t, w.Code, http.StatusInternalServerError)
}

// TestStoreCreateNoFieldFail ensures register can fail with no fields.
func TestStoreCreateNoFieldFail(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(register.Handler)
	db := new(database.MockService)
	h.UserService = database.NewUserService(db)
	h.ViewService = view.New("../../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusBadRequest)
}

// TestStoreCreateOneMissingFieldFail ensures register can fail with one missing
// field.
func TestStoreCreateOneMissingFieldFail(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the request body.
	val := url.Values{}
	r.Form = val
	r.Form.Add("firstname", "John")
	//r.Form.Add("lastname", "Doe")
	r.Form.Add("email", "jdoe@example.com")
	r.Form.Add("password", "Pa$$w0rd")

	// Call the handler.
	h := new(register.Handler)
	db := new(database.MockService)
	h.UserService = database.NewUserService(db)
	h.ViewService = view.New("../../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusBadRequest)
}
