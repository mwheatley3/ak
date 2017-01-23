package web

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// A Handler is an http.Handler with some additional functionality
type Handler interface {
	Handle(context.Context, http.ResponseWriter, *http.Request, httprouter.Params)
}

// HandlerFunc is a func that satisfies Handler
type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request, httprouter.Params)

// Handle satisfies the Handler interface
func (fn HandlerFunc) Handle(c context.Context, w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fn(c, w, r, p)
}

func wrapHandler(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h.Handle(context.Background(), w, r, p)
	}
}
