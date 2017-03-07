package login

import (
	"fmt"
	"net/http"

	"github.com/josephspurrier/gocleanarchitecture/domain/user"
	"github.com/josephspurrier/gocleanarchitecture/lib/view"
)

// Handler represents the services required for this controller.
type Handler struct {
	UserService user.Service
	ViewService view.Service
}

// Index displays the logon screen.
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.Store(w, r)
		return
	}

	h.ViewService.SetTemplate("login/index").Render(w, r)
}

// Store handles the submission of the login information.
func (h *Handler) Store(w http.ResponseWriter, r *http.Request) {
	// Don't continue if required fields are missing.
	for _, v := range []string{"email", "password"} {
		if len(r.FormValue(v)) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `<html>One or more required fields are missing. `+
				`Click <a href="/">here</a> to try again.</html>`)
			return
		}
	}

	u := new(user.Item)
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password")

	err := h.UserService.Authenticate(u)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `<html>Login failed. `+
			`Click <a href="/">here</a> to try again.</html>`)
		return
	}

	fmt.Fprint(w, "<html>Login successful!</html>")
}
