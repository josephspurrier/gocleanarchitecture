package view

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// BadResponseWriter represents a http.ResponseWriter that always fails.
type BadResponseWriter struct {
	Failed bool
}

// Header returns an empty header.
func (w *BadResponseWriter) Header() http.Header {
	return make(http.Header)
}

// Write always returns 0 and an error.
func (w *BadResponseWriter) Write(p []byte) (int, error) {
	w.Failed = true
	return 0, errors.New("writer failure")
}

// WriteHeader does nothing.
func (w *BadResponseWriter) WriteHeader(i int) {
}

// TestVar ensures the var functions work properly.
func TestVar(t *testing.T) {
	// Test adding and retrieving a variable.
	v := New("", "")
	v.AddVar("foo", "bar")
	AssertEqual(t, v.GetVar("foo"), "bar")

	// Test deleting a variable.
	v.DelVar("foo")
	AssertEqual(t, v.GetVar("foo"), nil)
}

// TestRenderFail ensures render fails properly.
func TestRenderFail(t *testing.T) {
	// Test adding and retrieving a variable.
	v := New("", "")

	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Fail on template parse error.
	err = v.Render(w, r)
	AssertNotNil(t, err)
	AssertEqual(t, w.Code, http.StatusInternalServerError)
}

// TestRenderExecuteFail ensures render fails properly.
func TestRenderExecuteFail(t *testing.T) {
	// Test adding and retrieving a variable.
	v := New("../../view", "tmpl")

	// Set up the request.
	br := new(BadResponseWriter)
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Fail on file parse error.
	err = v.Render(br, r)
	AssertNotNil(t, err)
	AssertEqual(t, br.Failed, true)
}
