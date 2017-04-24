package handler_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/adapter"
	"github.com/josephspurrier/gocleanarchitecture/adapter/jsonrepo"
	"github.com/josephspurrier/gocleanarchitecture/cmd/webapp/handler"
	"github.com/josephspurrier/gocleanarchitecture/domain"
	"github.com/josephspurrier/gocleanarchitecture/lib/jsondb"
	"github.com/josephspurrier/gocleanarchitecture/lib/view"

	"github.com/stretchr/testify/assert"
)

// TestLoginIndex ensures the index function returns a 200 code.
func TestLoginIndex(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(handler.Login)
	h.View = view.New("../html", "tmpl")
	h.Index(w, r)

	// Check the output.
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestLoginStoreMissingRequiredField ensures required fields should be entered.
func TestLoginStoreMissingRequiredFields(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(handler.Login)
	h.User = domain.NewUserService(
		jsonrepo.NewUserRepo(new(jsondb.MockService)),
		new(adapter.Passhash))
	h.View = view.New("../html", "tmpl")
	h.Store(w, r)

	// Check the output.
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestLoginStoreAuthenticateOK ensures login can be successful.
func TestLoginStoreAuthenticateOK(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	email := "jdoe@example.com"
	password := "Pa$$w0rd"

	// Set the request body.
	r.PostForm = url.Values{}
	r.PostForm.Add("email", email)
	r.PostForm.Add("password", password)

	// Call the handler.
	h := new(handler.Login)
	h.User = domain.NewUserService(
		jsonrepo.NewUserRepo(new(jsondb.MockService)),
		new(adapter.Passhash))
	h.View = view.New("../html", "tmpl")

	// Create a new user.
	err = h.User.Create("first", "last", email, password)
	if err != nil {
		t.Error(err)
	}

	h.Store(w, r)

	// Check the output.
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestLoginStoreAuthenticateFail ensures login can fail.
func TestLoginStoreAuthenticateFail(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	email := "jdoe2@example.com"
	password := "Pa$$w0rd"

	// Set the request body.
	r.Form = url.Values{}
	r.Form.Add("email", email)
	r.Form.Add("password", password)

	// Call the handler.
	h := new(handler.Login)
	h.User = domain.NewUserService(
		jsonrepo.NewUserRepo(new(jsondb.MockService)),
		new(adapter.Passhash))
	h.View = view.New("../html", "tmpl")

	// Create a new user.
	err = h.User.Create("first", "last", email, password+"bad")
	if err != nil {
		t.Error(err)
	}

	h.Store(w, r)

	// Check the output.
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
