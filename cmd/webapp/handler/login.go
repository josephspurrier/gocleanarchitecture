package handler

import (
	"fmt"
	"net/http"

	"github.com/josephspurrier/gocleanarchitecture/domain"
)

// Login represents the services required for this controller.
type Login struct {
	User domain.IUserService
	View domain.IViewService
}

// Index displays the logon screen.
func (h *Login) Index(w http.ResponseWriter, r *http.Request) {
	// Handle 404.
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 Page Not Found")
		return
	}

	if r.Method == "POST" {
		h.Store(w, r)
		return
	}

	h.View.SetTemplate("login/index")
	h.View.Render(w, r)
}

// Store handles the submission of the login information.
func (h *Login) Store(w http.ResponseWriter, r *http.Request) {
	// Don't continue if required fields are missing.
	for _, v := range []string{"email", "password"} {
		if len(r.FormValue(v)) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `<html>One or more required fields are missing. `+
				`Click <a href="/">here</a> to try again.</html>`)
			return
		}
	}

	u := new(domain.User)
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password")

	err := h.User.Authenticate(u)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `<html>Login failed. `+
			`Click <a href="/">here</a> to try again.</html>`)
		return
	}

	fmt.Fprint(w, "<html>Login successful!</html>")
}
