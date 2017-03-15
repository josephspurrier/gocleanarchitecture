package domain

import "net/http"

// ViewVars maps a string key to a variable.
type ViewVars map[string]interface{}

// ViewCase represents a service for managing templates.
type ViewCase interface {
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
