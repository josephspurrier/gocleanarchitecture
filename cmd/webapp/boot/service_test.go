package boot_test

import (
	"os"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/cmd/webapp/boot"
	"github.com/josephspurrier/gocleanarchitecture/domain"
)

// TestRegisterServices ensures each of the services is set up properly.
func TestRegisterServices(t *testing.T) {
	// Register the services.
	s := boot.RegisterServices("../html")

	// Test the user service.
	_, err := s.User.ByEmail("notexist")
	AssertEqual(t, err, domain.ErrUserNotFound)

	// Test the view service.
	s.View.AddVar("foo", "bar")
	v := s.View.GetVar("foo")
	AssertEqual(t, v, "bar")

	// Cleanup
	_ = os.Remove("db.json")
}
