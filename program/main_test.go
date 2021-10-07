package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_HomePage(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	homePage(w, req)

	expected := "home page endpoint"

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

}
