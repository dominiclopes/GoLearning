package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	tests := []struct {
		name        string
		query_param string
		double      int
		status      int
		err         string
	}{
		{
			name:        "double of 2",
			query_param: "?v=2",
			double:      4,
			status:      http.StatusOK,
			err:         "",
		},
		{
			name:        "double of a",
			query_param: "?v=a",
			status:      http.StatusBadRequest,
			err:         "not a number: a",
		},
		{
			name:        "double when no query parameter provided",
			query_param: "?v=",
			status:      http.StatusBadRequest,
			err:         "missing value",
		},
		{
			name:   "double when no value provided",
			status: http.StatusBadRequest,
			err:    "missing value",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "localhost:8080/doble"+test.query_param, nil)
			if err != nil {
				t.Fatalf("error initializing the request, %v", err)
			}
			rw := httptest.NewRecorder()
			doubleHandler(rw, req)

			resp := rw.Result()
			defer resp.Body.Close()

			if resp.StatusCode != test.status {
				t.Fatalf("expected status %v, got %v", test.status, resp.StatusCode)
			}

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("could not read response, %v", err)
			}
			b = bytes.TrimSpace(b)

			if test.status == http.StatusOK {
				d, err := strconv.Atoi(string(b))
				if err != nil {
					t.Fatalf("expected an integer, got %s", b)
				}

				if d != test.double {
					t.Fatalf("expected double to be %v, got %v", test.double, d)
				}
			} else {
				if string(b) != test.err {
					t.Fatalf("expected error %q, got %q", test.err, b)
				}
			}
		})
	}
}
