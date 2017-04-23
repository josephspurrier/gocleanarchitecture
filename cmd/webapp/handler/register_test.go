package handler_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/adapter/jsonrepo"
	"github.com/josephspurrier/gocleanarchitecture/adapter/passhash"
	"github.com/josephspurrier/gocleanarchitecture/cmd/webapp/handler"
	"github.com/josephspurrier/gocleanarchitecture/domain"
	"github.com/josephspurrier/gocleanarchitecture/lib/jsondb"
	"github.com/josephspurrier/gocleanarchitecture/lib/view"

	"github.com/stretchr/testify/assert"
)

// TestRegisterIndex ensures the index function returns a 200 code.
func TestRegisterIndex(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(handler.Register)
	h.View = view.New("../html", "tmpl")
	h.Index(w, r)

	// Check the output.
	assert.Equal(t, w.Code, http.StatusOK)
}

// TestRegisterStoreCreateOK ensures register can be successful.
func TestRegisterStoreCreateOK(t *testing.T) {
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
	h := new(handler.Register)
	h.User = domain.NewUserService(
		jsonrepo.NewUserRepo(new(jsondb.MockService)),
		new(passhash.Item))
	h.View = view.New("../view", "tmpl")
	h.Store(w, r)

	// Check the output.
	assert.Equal(t, w.Code, http.StatusCreated)

	// Fail on duplicate creation.
	w = httptest.NewRecorder()
	h.Store(w, r)
	assert.Equal(t, w.Code, http.StatusInternalServerError)
}

// TestRegisterStoreCreateNoFieldFail ensures register can fail with no fields.
func TestRegisterStoreCreateNoFieldFail(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(handler.Register)
	h.User = domain.NewUserService(
		jsonrepo.NewUserRepo(new(jsondb.MockService)),
		new(passhash.Item))
	h.View = view.New("../view", "tmpl")
	h.Store(w, r)

	// Check the output.
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

// TestRegisterStoreCreateOneMissingFieldFail ensures register can fail with one missing
// field.
func TestRegisterStoreCreateOneMissingFieldFail(t *testing.T) {
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
	h := new(handler.Register)
	h.User = domain.NewUserService(
		jsonrepo.NewUserRepo(new(jsondb.MockService)),
		new(passhash.Item))
	h.View = view.New("../view", "tmpl")
	h.Store(w, r)

	// Check the output.
	assert.Equal(t, w.Code, http.StatusBadRequest)
}
