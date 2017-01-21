package web

// wrap the handlers with access to the twitter client:
// https://medium.com/@matryer/the-http-handlerfunc-wrapper-technique-in-golang-c60bf76e6124#.skx82vp57
// also check out this guy's book

// func (s *Server) WithTwitterClient(fn http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		dbcopy := s.dbsession.Copy()
// 		defer dbcopy.Close()
// 		context.Set(r, "db", dbcopy)
// 		fn(w, r)
// 	}
// }
