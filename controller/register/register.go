package register

import (
	"fmt"
	"net/http"

	"gocleanarchitecture/domain/user"
	"gocleanarchitecture/lib/view"
)

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

	h.ViewService.SetTemplate("register/index").Render(w, r)
}

// Store adds a user to the database.
func (h *Handler) Store(w http.ResponseWriter, r *http.Request) {
	// Don't continue if required fields are missing.
	for _, v := range []string{"firstname", "lastname", "email", "password"} {
		if len(r.FormValue(v)) == 0 {
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

	// Add the user to the datbase.
	err := h.UserService.CreateUser(u)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, `<html>User created. `+
		`Click <a href="/">here</a> to login.</html>`)
}
