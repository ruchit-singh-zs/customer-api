package customer

import (
	"bytes"
	"customer-api/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type mockService struct {
}

func (m mockService) CreateCustomer(customer models.Customer) error {
	return nil
}

func (m mockService) GetCustomer(id string) error {
	return nil
}

func (m mockService) UpdateCustomer(id string) error {
	return nil
}

func (m mockService) DeleteCustomer(id string) error {
	return nil
}

func TestCreate(t *testing.T) {
	h := New(mockService{})
	res := `{"id":7,"name":"Shreya S","phoneNo":"9909111122","address":"BG Road Bangalore"}`
	cases := []struct {
		desc               string
		id                 string
		body               []byte
		expectedStatusCode int
		expectedResponse   string
	}{
		{"customer created Successfully", "7", []byte(res), http.StatusOK, "successfully created"},
	}

	for x, v := range cases {
		req := httptest.NewRequest(http.MethodPost, "http://customer", bytes.NewReader(v.body))
		w := httptest.NewRecorder()

		h.Create(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != v.expectedStatusCode {
			t.Errorf("Test[%v] Failed\n desc: %v\nExpected: %v \tGot: %v", x, v.desc, v.expectedStatusCode, w.Code)
		}

		expected := bytes.NewBuffer([]byte(v.expectedResponse))
		if !reflect.DeepEqual(w.Body, expected) {
			t.Errorf("Test[%v] Failed\n desc: %v\nExpected: %v \tGot: %v", x, v.desc, expected.String(), w.Body.String())
		}
	}
}

//func TestGetByID(t *testing.T) {
//	h:=New(mockService{})
//	cases := []struct {
//		desc               string
//		id                 string
//		body               []byte
//		expectedStatusCode int
//		expectedResponse   string
//	}{
//		{"customer exists", "1", nil, http.StatusOK, `{"id":1,"name":"Ruchit S","phoneNo":"4523946525","address":"BLR"}`},
//		{"customer does not exists", "10", nil, http.StatusNotFound, "No Record Exists"},
//	}
//
//	for i, v := range cases {
//		req := httptest.NewRequest(http.MethodGet, "http://customer", nil)
//		r := mux.SetURLVars(req, map[string]string{"id": v.id})
//		w := httptest.NewRecorder()
//
//		h.GetByID(w, r)
//
//		resp := w.Result()
//		defer resp.Body.Close()
//
//		if resp.StatusCode != v.expectedStatusCode {
//			t.Errorf("Test[%v] Failed desc: %v\nExpected: %v \tGot %v", i, v.desc, v.expectedStatusCode, w.Code)
//		}
//
//		expected := bytes.NewBuffer([]byte(v.expectedResponse))
//		if !reflect.DeepEqual(w.Body, expected) {
//			t.Errorf("Test[%v] Failed\n desc: %v\n Expected: %v \tGot: %v", i, v.desc, expected.String(), w.Body.String())
//		}
//	}
//}
//
//func TestUpdateByID(t *testing.T) {
//	h:=New(mockService{})
//	resp := `{"id":7,"name":"Divya S","phoneNo":"9909111143","address":"HSR Bangalore"}`
//	testcases := []struct {
//		desc               string
//		id                 string
//		body               []byte
//		expectedStatusCode int
//		expectedResponse   string
//	}{
//		{"customer updated successfully", "7", []byte(resp), http.StatusOK, "Updated Successfully"},
//	}
//
//	for i, v := range testcases {
//		req := httptest.NewRequest(http.MethodPut, "http://customer", bytes.NewReader(v.body))
//		r := mux.SetURLVars(req, map[string]string{"id": v.id})
//		w := httptest.NewRecorder()
//
//		h.UpdateByID(w, r)
//
//		resp := w.Result()
//		defer resp.Body.Close()
//
//		if resp.StatusCode != v.expectedStatusCode {
//			t.Errorf("Test[%v] Failed \ndesc: %v \nExpected: %v \tGot: %v", v.desc, i, v.expectedStatusCode, w.Code)
//		}
//	}
//}
//
//func TestDeleteByID(t *testing.T) {
//	h:=New(mockService{})
//	cases := []struct {
//		desc               string
//		id                 string
//		body               []byte
//		expectedStatusCode int
//		expectedResponse   string
//	}{
//		{"customer deleted successfully", "7", nil, http.StatusOK, "Deleted Successfully"},
//		{"customer record doesn't exist", "16", nil, http.StatusOK, "Deleted Successfully"},
//	}
//
//	for i, v := range cases {
//		req := httptest.NewRequest(http.MethodDelete, "http://customer", nil)
//		r := mux.SetURLVars(req, map[string]string{"id": v.id})
//		w := httptest.NewRecorder()
//
//		h.DeleteByID(w, r)
//
//		resp := w.Result()
//		defer resp.Body.Close()
//
//		if resp.StatusCode != v.expectedStatusCode {
//			t.Errorf("Test[%v] Failed \ndesc: %v\nExpected: %v \tGot: %v", i, v.desc, v.expectedStatusCode, w.Code)
//		}
//
//		expected := bytes.NewBuffer([]byte(v.expectedResponse))
//		if !reflect.DeepEqual(w.Body, expected) {
//			t.Errorf("Test[%v] Failed\n desc: %v\nExpected: %v \tGot: %v", i, v.desc, expected.String(), w.Body.String())
//		}
//	}
//}
