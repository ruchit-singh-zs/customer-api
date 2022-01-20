package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetByID(t *testing.T) {
	cases := []struct {
		desc               string
		id                 string
		body               []byte
		expectedStatusCode int
		expectedResponse   string
	}{
		{"customer exists", "1", nil, http.StatusOK, `{"customerId":1,"name":"Ruchit S","phoneNo":"4523946525","address":"BLR"}`},
		{"customer does not exists", "10", nil, http.StatusNotFound, "No Record Exists"},
	}

	for i, tc := range cases {
		req := httptest.NewRequest(http.MethodGet, "http://customer", nil)
		r := mux.SetURLVars(req, map[string]string{"id": tc.id})
		w := httptest.NewRecorder()

		// fixme : Mock everything
		//Handler.GetByID(w,r)

		resp := w.Result()

		if resp.StatusCode != tc.expectedStatusCode {
			t.Errorf("TestCase %v FAILED\nExpected : %v\tGot : %v", i, tc.expectedStatusCode, w.Code)
		}

		expected := bytes.NewBuffer([]byte(tc.expectedResponse))
		if !reflect.DeepEqual(w.Body, expected) {
			t.Errorf("TestCase %v FAILED\nExpected : %v\tGot : %v", i, expected.String(), w.Body.String())
		}
	}
}

func TestDeleteByID(t *testing.T) {
	testcases := []struct {
		desc               string
		id                 string
		body               []byte
		expectedStatusCode int
		expectedResponse   string
	}{
		{"customer deleted succesfully", "6", nil, http.StatusOK, "Deleted Successfully"},
		{"customer record doesn't exist", "16", nil, http.StatusOK, "Deleted Successfully"},
	}

	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodDelete, "http://customer", nil)
		r := mux.SetURLVars(req, map[string]string{"id": v.id})
		w := httptest.NewRecorder()

		DeleteByID(w, r)

		resp := w.Result()

		if resp.StatusCode != v.expectedStatusCode {
			t.Errorf("TestCase %v FAILED\nExpected : %v\tGot : %v", i, v.expectedStatusCode, w.Code)
		}

		expected := bytes.NewBuffer([]byte(v.expectedResponse))
		if !reflect.DeepEqual(w.Body, expected) {
			t.Errorf("TestCase %v FAILED\nExpected : %v\tGot : %v", i, expected.String(), w.Body.String())
		}
	}
}

func TestCreate(t *testing.T) {
	testcases := []struct {
		desc               string
		id                 string
		body               []byte
		expectedStatusCode int
		expectedResponse   string
	}{
		{"customer created Succesfully", "6", []byte(`{"customerId":6,"name":"Rahul","phoneNo":"9909111143","address":"BG Road Bangalore"}`), http.StatusOK, "Succesfully created"},
	}
	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodPost, "http://customer", bytes.NewReader(v.body))
		r := mux.SetURLVars(req, map[string]string{"id": v.id})
		w := httptest.NewRecorder()

		Create(w, r)

		resp := w.Result()

		if resp.StatusCode != v.expectedStatusCode {
			t.Errorf("TestCase %v FAILED\nExpected : %v\tGot : %v", i, v.expectedStatusCode, w.Code)
		}

		expected := bytes.NewBuffer([]byte(v.expectedResponse))
		if !reflect.DeepEqual(w.Body, expected) {
			t.Errorf("TestCase %v FAILED\nExpected : %v\tGot : %v", i, expected.String(), w.Body.String())
		}
	}
}

func TestUpdateByID(t *testing.T) {
	testcases := []struct {
		desc               string
		id                 string
		body               []byte
		expectedStatusCode int
		expectedResponse   string
	}{
		{"customer updated successfully", "4", []byte(`{"customerId":4,"name":"Aakanksha J	","phoneNo":"9909111143","address":"HSR Bangalore"}`), http.StatusOK, "Updated Successfully"},
	}
	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodPut, "http://customer", bytes.NewReader(v.body))
		r := mux.SetURLVars(req, map[string]string{"id": v.id})
		w := httptest.NewRecorder()

		UpdateByID(w, r)

		resp := w.Result()

		if resp.StatusCode != v.expectedStatusCode {
			t.Errorf("TestCase %v FAILED\nExpected : %v\tGot : %v", i, v.expectedStatusCode, w.Code)
		}
	}
}
