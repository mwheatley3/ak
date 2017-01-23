package web

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

//GetUser gets a user
func (s *Server) GetUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var (
		meUser = p.ByName("userID") == "me"
		resp   = JSON(nil, res)
	)

	// handle other users later!
	if !meUser {
		resp.Error(errors.New("kljshadf"), http.StatusNotFound)
		return
	}

	u := User{
		ID:    uuid.NewV1(),
		Email: "test",
	}

	// switch err {
	// case nil:
	resp.Success(u)
	// case api.ErrInvalidUserID:
	// 	resp.Error(err, http.StatusNotFound)
	// default:
	// 	resp.Error(err, http.StatusInternalServerError)
	// }
}
