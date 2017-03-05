package boot

import (
	"net/http"

	"github.com/josephspurrier/gocleanarchitecture/controller/login"
	"github.com/josephspurrier/gocleanarchitecture/controller/register"
)

// LoadRoutes returns a handler with all the routes.
func (s *Service) LoadRoutes() http.Handler {
	// Create the mux.
	h := http.NewServeMux()

	// Register the pages.
	s.AddLogin(h)
	s.AddRegister(h)

	// Return the handler.
	return h
}

// AddLogin registers the login handlers.
func (s *Service) AddLogin(mux *http.ServeMux) {
	// Create handler.
	h := new(login.Handler)

	// Assign services.
	h.UserService = s.UserService
	h.ViewService = s.ViewService

	// Load routes.
	mux.HandleFunc("/", h.Index)
}

// AddRegister registers the register handlers.
func (s *Service) AddRegister(mux *http.ServeMux) {
	// Create handler.
	h := new(register.Handler)

	// Assign services.
	h.UserService = s.UserService
	h.ViewService = s.ViewService

	// Load routes.
	mux.HandleFunc("/register", h.Index)
}
