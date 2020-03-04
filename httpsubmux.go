package httpsubmux

import (
	"net/http"
)

type ServeMux interface {
	Handle(pattern string, handler http.Handler)
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type SubMux struct {
	Pattern string
	ParentRoute string
	Route string
	mux *http.ServeMux
}

func NewServeMux(parent string, pattern string) *SubMux {
	return &SubMux{
		Pattern: pattern,
		ParentRoute: parent,
		Route: parent + pattern,
		mux: http.NewServeMux(),
	}
}

func (submux *SubMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	submux.mux.ServeHTTP(w, r)
}

func (submux *SubMux) Handle(pattern string, handler http.Handler) {
	submux.mux.Handle(pattern, handler)
}

func (submux *SubMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	submux.mux.HandleFunc(pattern, handler)
}
