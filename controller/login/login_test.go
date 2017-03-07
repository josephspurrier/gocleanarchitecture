package login_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/controller/login"
	"github.com/josephspurrier/gocleanarchitecture/lib/view"
)

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
	h := new(login.Handler)
	h.ViewService = view.New("../../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, 200)
}
