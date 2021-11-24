package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testsLists map[string]struct {
	handler func(http.ResponseWriter, *http.Request)
	method  string
	path    string
	code    int
	body    string
}

func TestAll(t *testing.T) {
	tests := testsLists{
		"homePage": {
			handler: homePage,
			method:  "GET",
			path:    "/",
			code:    200,
			body:    "home page endpoint",
		},
		"healthPage": {
			handler: healthCheck,
			method:  "GET",
			path:    "/health",
			code:    200,
			body:    "its ok",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatalf("failed to create request: %s", err)
			}
			w := httptest.NewRecorder()
			tt.handler(w, req)

			expected := tt.body

			respo, err := ioutil.ReadAll(w.Result().Body)
			if err != nil {
				t.Errorf("cant read body")
			}

			expected_resp_code := 200
			resp_code := w.Result().StatusCode

			if expected_resp_code != resp_code {
				t.Errorf("wrong code. Expected was %d. Actual is %d", expected_resp_code, resp_code)
			}

			if expected != string(respo) {
				t.Errorf("dint work. Expected was %s. Actual is %s", expected, respo)
			}

		})
	}
}
