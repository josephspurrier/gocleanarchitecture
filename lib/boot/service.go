package boot

import (
	"gocleanarchitecture/database"
	"gocleanarchitecture/domain"
	"gocleanarchitecture/domain/user"
	"gocleanarchitecture/lib/view"
)

// Service represents all the services that the application uses.
type Service struct {
	ContextService domain.ContextService
	UserService    user.Service
	ViewService    view.Service
}

// RegisterServices sets up each service and returns the container for all
// the services.
func RegisterServices() *Service {
	s := new(Service)

	dbClient := database.NewClient("db.json")
	viewService := view.New()

	s.UserService = dbClient.UserService()
	s.ViewService = viewService

	return s
}
