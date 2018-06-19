package instruction

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MyError struct {
	Debug   string
	Message string
	Code    int
}

func (e *MyError) StatusCode() int {
	return e.Code
}

func (e *MyError) Error() string {
	return e.Message
}

func (e *MyError) WithDebug(debug string) *MyError {
	err := new(MyError)
	*err = *e
	err.Debug = debug
	return err
}

var (
	ErrNotFound = &MyError{
		Message: "could not be found",
		Code:    http.StatusNotFound,
	}
)

func TestError(t *testing.T) {
	var err error = ErrNotFound.WithDebug("My debug message")

	type statusCodeCarrier interface {
		StatusCode() int
	}

	if e, ok := err.(statusCodeCarrier); ok {
		t.Logf("Send status code: %d", e.StatusCode())
	}

	t.Logf("Could not handle request because %+v", err)
}

func TestHandler(t *testing.T) {
	router := mux.NewRouter()
	handler := &Handler{}
	handler.RegisterRoutes(router)
	ts := httptest.NewServer(router)

	for k, tc := range []struct {
		u  string
		ec int
	}{
		{
			u:  ts.URL + "/instruction",
			ec: http.StatusNoContent,
		},
		{
			u:  ts.URL + "/not-found",
			ec: http.StatusNotFound,
		},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			response, err := http.Get(tc.u)
			require.NoError(t, err)
			assert.EqualValues(t, tc.ec, response.StatusCode)
		})
	}
}
