package domain

import "net/http"

// ViewVars maps a string key to a variable.
type ViewVars map[string]interface{}

// IViewService is the interface for HTML templates.
type IViewService interface {
	Render(w http.ResponseWriter, r *http.Request) error
	SetFolder(relativeFolderPath string)
	SetExtension(fileExtension string)
	SetBaseTemplate(relativeFilePath string)
	SetTemplate(relativeFilePath string)

	AddVar(key string, value interface{})
	DelVar(key string)
	GetVar(key string) interface{}
	SetVars(vars ViewVars)
}
