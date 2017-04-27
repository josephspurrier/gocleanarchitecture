package boot

import (
	"fmt"
	"net/http"

	"github.com/josephspurrier/gocleanarchitecture/adapter"
	"github.com/josephspurrier/gocleanarchitecture/cmd/webapp/handler"
	"github.com/josephspurrier/gocleanarchitecture/lib/router"
)

// LoadRoutes returns a handler with all the routes.
func (s *Service) LoadRoutes() http.Handler {
	// Create the router.
	h := router.New()

	// Set the 404 page.
	h.SetNotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Page Not Found")
	})

	// Register the pages.
	s.addLogin(h)
	s.addRegister(h)

	// Return the handler.
	return h.Router()
}

// addLogin registers the login handlers.
func (s *Service) addLogin(r adapter.IRouterService) {
	// Create handler.
	h := new(handler.Login)

	// Assign services.
	h.User = s.User
	h.View = s.View

	// Load routes.
	r.Get("/", h.Index)
	r.Post("/", h.Store)
}

// addRegister registers the register handlers.
func (s *Service) addRegister(r adapter.IRouterService) {
	// Create handler.
	h := new(handler.Register)

	// Assign services.
	h.User = s.User
	h.View = s.View

	// Load routes.
	r.Get("/register", h.Index)
	r.Post("/register", h.Store)
}
