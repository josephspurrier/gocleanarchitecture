package boot

import (
	"github.com/josephspurrier/gocleanarchitecture/adapter/passhash"
	"github.com/josephspurrier/gocleanarchitecture/adapter/repo"
	"github.com/josephspurrier/gocleanarchitecture/domain"
	"github.com/josephspurrier/gocleanarchitecture/lib/jsondb"
	"github.com/josephspurrier/gocleanarchitecture/lib/view"
)

// Service represents all the services that the application uses.
type Service struct {
	User domain.IUserService
	View domain.IViewService
}

// RegisterServices sets up each service and returns the container for all
// the services.
func RegisterServices(templateFolder string) *Service {
	s := new(Service)

	// Initialize the clients.
	db := jsondb.NewClient("db.json")

	// Store all the services for the application.
	s.User = domain.NewUserService(
		repo.NewUserRepo(db),
		new(passhash.Item))
	s.View = view.New(templateFolder, "tmpl")

	return s
}
