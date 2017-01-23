package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/Sirupsen/logrus"
)

var badMarshalResponse, _ = json.Marshal(&JSONResponse{
	Error: &JSONError{
		Message: "Unknown Error",
	},
})

// A JSONResponder provides functionality for writing json
// http responses
type JSONResponder struct {
	w       http.ResponseWriter
	logger  *logrus.Logger
	st      time.Time
	written int32
}

func (j *JSONResponder) init() {
	j.st = time.Now()
	j.w.Header().Add("Content-Type", "application/json")
}

// Error returns a json error response
func (j *JSONResponder) Error(err error, code int) {
	j.write(&JSONResponse{Error: NewJSONError(err)}, code)
}

// Success returns a json success response
func (j *JSONResponder) Success(data interface{}) {
	j.write(&JSONResponse{Data: data}, http.StatusOK)
}

// Write writes a success or error response depending on
// if err is nil.  If an err is passed in, a 500 status code is
// returned as well. For more fine grained control, use the
// Error method
func (j *JSONResponder) Write(data interface{}, err error) {
	if err != nil {
		j.Error(err, http.StatusInternalServerError)
	} else {
		j.Success(data)
	}
}

func (j *JSONResponder) write(resp *JSONResponse, code int) {
	first := atomic.CompareAndSwapInt32(&j.written, 0, 1)

	if !first {
		panic("Multiple calls to JSONResponder.write. Don't do that")
	}

	resp.Duration = int(time.Since(j.st) / time.Millisecond)

	b, err := json.Marshal(resp)

	if err != nil {
		j.logger.Errorf("JSONResponder unmarshal error: %s", err.Error())
		j.w.WriteHeader(500)
		j.w.Write(badMarshalResponse)
	} else {
		j.w.WriteHeader(code)
		j.w.Write(b)
	}
}

// JSON intializes a json http response.  The function that's returned
// should be called with the data or error that should be returned
func JSON(l *logrus.Logger, w http.ResponseWriter) *JSONResponder {
	js := &JSONResponder{w: w, logger: l}
	js.init()

	return js
}

// JSONResponse is the envelope returned from an api endpoint handler
type JSONResponse struct {
	Data     interface{} `json:"data"`
	Error    *JSONError  `json:"error"`
	Duration int         `json:"duration"`
}

// NewJSONError returns a JSONError initialized from err
func NewJSONError(err error) *JSONError {
	// ignore json marshalling errors for now
	b, _ := json.Marshal(err)

	return &JSONError{
		Message: err.Error(),
		Data:    b,
	}
}

// JSONError is the error representation in the api response
type JSONError struct {
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

// DecodeJSON decodes a json payload into dest and makes
// sure that the body is closed
func DecodeJSON(r *http.Request, dest interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(dest)
}

func (j *JSONError) String() string {
	return fmt.Sprintf("Message: %s Data: %s", j.Message, string(j.Data))
}
