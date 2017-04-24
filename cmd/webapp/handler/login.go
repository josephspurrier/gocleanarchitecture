package handler

import (
	"fmt"
	"net/http"

	"github.com/josephspurrier/gocleanarchitecture/cmd/webapp/adapter"
	"github.com/josephspurrier/gocleanarchitecture/domain"
)

// Login represents the services required for this controller.
type Login struct {
	User domain.IUserService
	View adapter.IViewService
}

// Index displays the logon screen.
func (h *Login) Index(w http.ResponseWriter, r *http.Request) {
	h.View.SetTemplate("login/index")
	h.View.Render(w, r)
}

// Store handles the submission of the login information.
func (h *Login) Store(w http.ResponseWriter, r *http.Request) {
	// Don't continue if required fields are missing.
	for _, v := range []string{"email", "password"} {
		if len(r.PostFormValue(v)) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `<html>One or more required fields are missing. `+
				`Click <a href="/">here</a> to try again.</html>`)
			return
		}
	}

	err := h.User.Authenticate(r.PostFormValue("email"),
		r.PostFormValue("password"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `<html>Login failed. `+
			`Click <a href="/">here</a> to try again.</html>`)
		return
	}

	fmt.Fprint(w, "<html>Login successful!</html>")
}
