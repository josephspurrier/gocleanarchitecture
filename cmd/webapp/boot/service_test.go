package boot_test

import (
	"os"
	"testing"

	. "github.com/josephspurrier/gocleanarchitecture/cmd/webapp/boot"
	"github.com/josephspurrier/gocleanarchitecture/domain"

	"github.com/stretchr/testify/assert"
)

// TestRegisterServices ensures each of the services is set up properly.
func TestRegisterServices(t *testing.T) {
	// Register the services.
	s := RegisterServices("../html")

	// Test the user service.
	_, err := s.User.ByEmail("notexist")
	assert.Equal(t, domain.ErrUserNotFound, err)

	// Test the view service.
	s.View.AddVar("foo", "bar")
	v := s.View.GetVar("foo")
	assert.Equal(t, "bar", v)

	// Cleanup
	_ = os.Remove("db.json")
}
