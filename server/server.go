package server

import (
	"context"
	"github.com/mwheatley3/ak/server/twitter"
	"net/http"
)

//Server wraps an http server
type Server struct {
	TwitterClient *twitter.Client
	HTTPServer    http.Server
	Router        *http.ServeMux
}

// WithTwitterClient allows handlers to access the twitter client
func (s *Server) WithTwitterClient(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.Background(), "twitterClient", s.TwitterClient)
		fn(w, r.WithContext(ctx))
	}
}
