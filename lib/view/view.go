package view

import (
	"html/template"
	"net/http"
	"path"
)

// Item represents a view template.
type Item struct {
	Extension string
	Folder    string
	Vars      map[string]interface{}

	baseTemplate string
	template     string
}

// Service represents a service for managing a template.
type Service interface {
	Render(w http.ResponseWriter, r *http.Request) error
	SetBaseTemplate(relativeFilePath string) *Item
	SetTemplate(relativeFilePath string) *Item
}

// New returns a new template.
func New(folder string, extension string) *Item {
	t := new(Item)

	// Set the initial values.
	t.Folder = folder
	t.Extension = extension
	t.SetTemplate("default")
	t.SetBaseTemplate("base")

	return t
}

// SetBaseTemplate sets the base template to render.
func (t *Item) SetBaseTemplate(s string) *Item {
	t.baseTemplate = path.Join(t.Folder, s)
	return t
}

// SetTemplate sets the template to render.
func (t *Item) SetTemplate(s string) *Item {
	t.template = path.Join(t.Folder, s)
	return t
}

// Render outputs the template to the ResponseWriter.
func (t *Item) Render(w http.ResponseWriter, r *http.Request) error {
	// Determine if there is an error in the template syntax.
	tc, err := template.ParseFiles(t.baseTemplate+"."+t.Extension,
		t.template+"."+t.Extension)
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
