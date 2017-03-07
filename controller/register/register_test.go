package register_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/controller/register"
	"github.com/josephspurrier/gocleanarchitecture/domain/user"
	"github.com/josephspurrier/gocleanarchitecture/lib/view"
)

// AssetEqual throws an error if the two values are not equal.
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

// TestStoreMissingRequiredField ensures required fields should be entered.
func TestStoreMissingRequiredFields(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(register.Handler)
	h.UserService = new(user.MockService)
	h.ViewService = view.New("../../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusBadRequest)
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
	h.UserService = new(user.MockService)
	h.ViewService = view.New("../../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusCreated)
}

// TestStoreAuthenticateFail ensures register can fail.
func TestStoreAuthenticateFail(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the request body.
	val := url.Values{}
	r.Form = val
	r.Form.Add("email", "jdoe2@example.com")
	r.Form.Add("password", "BadPa$$w0rd")

	// Call the handler.
	h := new(register.Handler)
	h.UserService = new(user.MockService)
	h.ViewService = view.New("../../view", "tmpl")

	// Create a new user.
	u := new(user.Item)
	u.Email = "jdoe2@example.com"
	u.Password = "Pa$$w0rd"
	h.UserService.CreateUser(u)

	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusBadRequest)
}
