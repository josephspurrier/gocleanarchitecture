package boot_test

import (
	"os"
	"testing"

	"github.com/josephspurrier/gocleanarchitecture/domain"
	"github.com/josephspurrier/gocleanarchitecture/lib/boot"
)

// TestRegisterServices ensures each of the services is set up properly.
func TestRegisterServices(t *testing.T) {
	// Register the services.
	s := boot.RegisterServices()

	// Test the user service.
	_, err := s.UserService.User("notexist")
	AssertEqual(t, err, domain.ErrUserNotFound)

	// Test the view service.
	s.ViewService.AddVar("foo", "bar")
	v := s.ViewService.GetVar("foo")
	AssertEqual(t, v, "bar")

	// Cleanup
	os.Remove("db.json")
}
