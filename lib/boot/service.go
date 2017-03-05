package boot

import (
	"gocleanarchitecture/database"
	"gocleanarchitecture/domain/user"
	"gocleanarchitecture/lib/view"
)

// Service represents all the services that the application uses.
type Service struct {
	UserService user.Service
	ViewService view.Service
}

// RegisterServices sets up each service and returns the container for all
// the services.
func RegisterServices() *Service {
	s := new(Service)

	// Initialize the clients.
	dbClient := database.NewClient("db.json")

	// Store all the services for the application.
	s.UserService = dbClient.UserService()
	s.ViewService = view.New("../../view", "tmpl")

	return s
}
