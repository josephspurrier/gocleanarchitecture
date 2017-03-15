package view

import (
	"html/template"
	"net/http"
	"path"

	"github.com/josephspurrier/gocleanarchitecture/domain"
)

// Item represents a view template.
type Item struct {
	extension string
	folder    string

	baseTemplate string
	template     string
	vars         domain.ViewVars
}

// New returns a new template.
func New(folder string, extension string) *Item {
	v := new(Item)

	// Set the initial values.
	v.SetFolder(folder)
	v.SetExtension(extension)
	v.SetBaseTemplate("base")
	v.SetTemplate("default")
	v.SetVars(domain.ViewVars{})

	return v
}

// SetFolder sets the folder containing the templates.
func (v *Item) SetFolder(s string) {
	v.folder = s
}

// SetExtension sets the extensions of the templates.
func (v *Item) SetExtension(s string) {
	v.extension = s
}

// SetBaseTemplate sets the base template to render.
func (v *Item) SetBaseTemplate(s string) {
	v.baseTemplate = path.Join(v.folder, s)
}

// SetTemplate sets the template to render.
func (v *Item) SetTemplate(s string) {
	v.template = path.Join(v.folder, s)
}

// AddVar adds a variable to the template variable map.
func (v *Item) AddVar(key string, value interface{}) {
	v.vars[key] = value
}

// DelVar removes a variable from the template variable map.
func (v *Item) DelVar(key string) {
	delete(v.vars, key)
}

// GetVar returns a value from the template variable map.
func (v *Item) GetVar(key string) interface{} {
	value, ok := v.vars[key]
	if !ok {
		return nil
	}
	return value
}

// SetVars sets the template variable map.
func (v *Item) SetVars(vars domain.ViewVars) {
	v.vars = vars
}

// Render outputs the template to the ResponseWriter.
func (v *Item) Render(w http.ResponseWriter, r *http.Request) error {
	// Determine if there is an error in the template syntax.
	tc, err := template.ParseFiles(v.baseTemplate+"."+v.extension,
		v.template+"."+v.extension)
	if err != nil {
		http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
		return err
	}

	// Display the content to the screen.
	err = tc.Execute(w, v.vars)
	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}
