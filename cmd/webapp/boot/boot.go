package boot

import (
	"net/http"
)

// ServicesAndRoutes returns an HTTP handler after registering the services
// and loading the routes.
func ServicesAndRoutes(templateFolder string) http.Handler {
	return RegisterServices(templateFolder).LoadRoutes()
}
