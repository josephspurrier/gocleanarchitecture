package boot

import (
	"github.com/josephspurrier/gocleanarchitecture/domain"
	"github.com/josephspurrier/gocleanarchitecture/lib/passhash"
	"github.com/josephspurrier/gocleanarchitecture/lib/view"
	"github.com/josephspurrier/gocleanarchitecture/repository"
	"github.com/josephspurrier/gocleanarchitecture/usecase"
)

// Service represents all the services that the application uses.
type Service struct {
	UserService domain.UserCase
	ViewService view.Service
}

// RegisterServices sets up each service and returns the container for all
// the services.
func RegisterServices() *Service {
	s := new(Service)

	// Initialize the clients.
	db := repository.NewClient("db.json")

	// Store all the services for the application.
	s.UserService = usecase.NewUserCase(
		repository.NewUserRepo(db),
		new(passhash.Item))
	s.ViewService = view.New("../../view", "tmpl")

	return s
}
