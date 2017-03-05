package view

import (
	"html/template"
	"net/http"
	"path"
)

// Vars maps a string key to a variable.
type Vars map[string]interface{}

// Item represents a view template.
type Item struct {
	Extension string
	Folder    string

	baseTemplate string
	template     string
	vars         Vars
}

// Service represents a service for managing a template.
type Service interface {
	Render(w http.ResponseWriter, r *http.Request) error
	SetBaseTemplate(relativeFilePath string) *Item
	SetTemplate(relativeFilePath string) *Item

	AddVar(key string, value interface{}) *Item
	DelVar(key string) *Item
	GetVar(key string) interface{}
	SetVars(vars Vars) *Item
}

// New returns a new template.
func New(folder string, extension string) *Item {
	v := new(Item)

	// Set the initial values.
	v.Folder = folder
	v.Extension = extension
	v.SetBaseTemplate("base")
	v.SetTemplate("default")
	v.SetVars(Vars{})

	return v
}

// SetBaseTemplate sets the base template to render.
func (v *Item) SetBaseTemplate(s string) *Item {
	v.baseTemplate = path.Join(v.Folder, s)
	return v
}

// SetTemplate sets the template to render.
func (v *Item) SetTemplate(s string) *Item {
	v.template = path.Join(v.Folder, s)
	return v
}

// AddVar adds a variable to the template variable map.
func (v *Item) AddVar(key string, value interface{}) *Item {
	v.vars[key] = value
	return v
}

// DelVar removes a variable from the template variable map.
func (v *Item) DelVar(key string) *Item {
	delete(v.vars, key)
	return v
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
func (v *Item) SetVars(vars Vars) *Item {
	v.vars = vars
	return v
}

// Render outputs the template to the ResponseWriter.
func (v *Item) Render(w http.ResponseWriter, r *http.Request) error {
	// Determine if there is an error in the template syntax.
	tc, err := template.ParseFiles(v.baseTemplate+"."+v.Extension,
		v.template+"."+v.Extension)
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
