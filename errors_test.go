package httperrors

import (
	"fmt"
	"net/http"
	"testing"
)

type errorTestNotfound struct {
}

func (e *errorTestNotfound) StatusCode() int {
	return http.StatusNotFound
}

func (e *errorTestNotfound) Error() string {
	return "not-found"
}

type errorTestInternal struct {
}

func (e *errorTestInternal) StatusCode() int {
	return http.StatusInternalServerError
}

func (e *errorTestInternal) Error() string {
	return "internal-server-error"
}

type errorTestCustomMessage struct {
	message string
}

func (e *errorTestCustomMessage) StatusCode() int {
	return http.StatusBadRequest
}

func (e *errorTestCustomMessage) Error() string {
	return e.message
}

func TestIdentify(t *testing.T) {
	testcases := map[string]struct {
		err             error
		messageExpected string
		statusExpected  int
	}{
		"Not found": {
			err:             &errorTestNotfound{},
			messageExpected: "not-found",
			statusExpected:  http.StatusNotFound,
		},
		"Internal server error": {
			err:             &errorTestInternal{},
			messageExpected: "internal-server-error",
			statusExpected:  http.StatusInternalServerError,
		},
		"custom message error": {
			err:             &errorTestCustomMessage{message: "custom-message-example"},
			messageExpected: "custom-message-example",
			statusExpected:  http.StatusBadRequest,
		},
		"any error": {
			err:             fmt.Errorf("any-error"),
			messageExpected: "any-error",
			statusExpected:  http.StatusInternalServerError,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			resultStatus, resultMessage := Identify(tc.err)
			if resultStatus != tc.statusExpected {
				t.Errorf("status: expected %d, got %d", tc.statusExpected, resultStatus)
			}
			if resultMessage != tc.messageExpected {
				t.Errorf("message: expected %s, got %s", tc.messageExpected, resultMessage)
			}
		})
	}
}
