package main

import (
	"github.com/josephspurrier/gocleanarchitecture/adapter"
	"github.com/josephspurrier/gocleanarchitecture/adapter/jsonrepo"
	appadapter "github.com/josephspurrier/gocleanarchitecture/cmd/webapp/adapter"
	"github.com/josephspurrier/gocleanarchitecture/domain"
	"github.com/josephspurrier/gocleanarchitecture/lib/jsondb"
	"github.com/josephspurrier/gocleanarchitecture/lib/view"
)

// Service represents all the services that the application uses.
type Service struct {
	User domain.IUserService
	View appadapter.IViewService
}

// RegisterServices sets up each service and returns the container for all
// the services.
func RegisterServices(templateFolder string) *Service {
	s := new(Service)

	// Initialize the clients.
	db := jsondb.New("jsondb")

	// Store all the services for the application.
	s.User = domain.NewUserService(
		jsonrepo.NewUserRepo(db),
		new(adapter.Passhash))
	s.View = view.New(templateFolder, "tmpl")

	return s
}
