package view_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/lib/view"
)

// AssertEqual throws an error if the two values are not equal.
func AssertEqual(t *testing.T, actualValue interface{}, expectedValue interface{}) {
	if actualValue != expectedValue {
		t.Errorf("\n got: %v\nwant: %v", actualValue, expectedValue)
	}
}

// TestVar ensures the var functions work properly.
func TestVar(t *testing.T) {
	// Test adding and retrieving a variable.
	v := view.New("", "")
	v.AddVar("foo", "bar")
	AssertEqual(t, v.GetVar("foo"), "bar")

	// Test deleting a variable.
	v.DelVar("foo")
	AssertEqual(t, v.GetVar("foo"), nil)
}

// TestRenderFail ensures render fails properly.
func TestRenderFail(t *testing.T) {
	// Test adding and retrieving a variable.
	v := view.New("", "")

	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Fail on template parse error.
	v.Render(w, r)
	AssertEqual(t, w.Code, http.StatusInternalServerError)
}
