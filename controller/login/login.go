package login

import (
	"fmt"
	"net/http"

	"gocleanarchitecture/domain"
	"gocleanarchitecture/domain/user"
	"gocleanarchitecture/lib/view"
)

type Handler struct {
	ContextService domain.ContextService
	UserService    user.Service
	ViewService    view.Service
}

// Index displays the logon screen.
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.Store(w, r)
		return
	}

	h.ViewService.Template("login/index").Render(w, r)
}

// Store handles the submission of the login information.
func (h *Handler) Store(w http.ResponseWriter, r *http.Request) {
	// Don't continue if required fields are missing.
	for _, v := range []string{"email", "password"} {
		if len(r.FormValue(v)) == 0 {
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
		fmt.Fprint(w, `<html>Login failed. `+
			`Click <a href="/">here</a> to try again.</html>`)
		return
	}

	fmt.Fprint(w, "<html>Login successful!</html>")
	return
}
