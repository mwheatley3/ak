package web

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mwheatley3/ak/server/twitter"
)

//Server wraps an http server
type Server struct {
	TwitterClient *twitter.Client
	HTTPServer    http.Server
	Router        *httprouter.Router
}

// WithTwitterClient allows handlers to access the twitter client
func (s *Server) WithTwitterClient(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := context.WithValue(context.Background(), "twitterClient", s.TwitterClient)
		fn(w, r.WithContext(ctx), p)
	}
}
