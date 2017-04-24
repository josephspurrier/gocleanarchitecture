package adapter

import (
	"net/http"

	"github.com/justinas/alice"
)

// IRouterService is the interface for page routing.
type IRouterService interface {
	Chain(c ...alice.Constructor) []alice.Constructor
	ChainHandler(h http.Handler, c ...alice.Constructor) http.Handler
	Delete(path string, fn http.HandlerFunc, c ...alice.Constructor)
	Get(path string, fn http.HandlerFunc, c ...alice.Constructor)
	Patch(path string, fn http.HandlerFunc, c ...alice.Constructor)
	Post(path string, fn http.HandlerFunc, c ...alice.Constructor)
	Put(path string, fn http.HandlerFunc, c ...alice.Constructor)
}
