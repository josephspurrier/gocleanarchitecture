package boot

import (
	"net/http"
)

// ServicesAndRoutes returns an HTTP handler after registering the services
// and loading the routes.
func ServicesAndRoutes() http.Handler {
	return RegisterServices().LoadRoutes()
}
