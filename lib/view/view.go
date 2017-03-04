package view

import (
	"html/template"
	"net/http"
	"path"
)

type Template struct {
	Extension string
	Folder    string
	Vars      map[string]interface{}

	rootTemplate string
	baseTemplate string
}

type Service interface {
	BaseTemplate(relativeFilePath string) *Template
	Render(w http.ResponseWriter, r *http.Request) error
	Template(relativeFilePath string) *Template
}

func New() *Template {
	t := new(Template)

	t.Folder = "../../view"
	t.rootTemplate = path.Join(t.Folder, "default")
	t.baseTemplate = path.Join(t.Folder, "base")
	t.Extension = "tmpl"

	return t
}

// BaseTemplate sets the base template to render.
func (t *Template) BaseTemplate(s string) *Template {
	t.baseTemplate = path.Join(t.Folder, s)
	return t
}

// Template sets the template to render.
func (t *Template) Template(s string) *Template {
	t.rootTemplate = path.Join(t.Folder, s)
	return t
}

// Render outputs the template to the ResponseWriter.
func (t *Template) Render(w http.ResponseWriter, r *http.Request) error {
	key := t.rootTemplate + "." + t.Extension
	base := t.baseTemplate + "." + t.Extension

	// Determine if there is an error in the template syntax.
	tc, err := template.ParseFiles(
		base,
		key)
	if err != nil {
		http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
		return err
	}

	// Display the content to the screen.
	err = tc.Execute(w, t.Vars)
	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}
