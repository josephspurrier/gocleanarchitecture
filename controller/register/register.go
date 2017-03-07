package register

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

// Index displays the register screen.
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.Store(w, r)
		return
	}

	h.ViewService.SetTemplate("register/index").Render(w, r)
}

// Store adds a user to the database.
func (h *Handler) Store(w http.ResponseWriter, r *http.Request) {
	// Don't continue if required fields are missing.
	for _, v := range []string{"firstname", "lastname", "email", "password"} {
		if len(r.FormValue(v)) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, `<html>One or more required fields are missing. `+
				`Click <a href="/register">here</a> to try again.</html>`)
			return
		}
	}

	// Build the user from the form values.
	u := new(user.Item)
	u.FirstName = r.FormValue("firstname")
	u.LastName = r.FormValue("lastname")
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password")

	// Add the user to the database.
	err := h.UserService.CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, `<html>User created. `+
		`Click <a href="/">here</a> to login.</html>`)
}
