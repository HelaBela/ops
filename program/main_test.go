package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

//table of expect4ed input and expected output.

//one function to loop over that - t.run() "table driven testing" - google that

//clues :D

//function name, expected output and stuts code - that goes to the table

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

func Test_HealthPoint(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	healthCheck(w, req)

	expected := "its ok"
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
