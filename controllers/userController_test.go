package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserCreationHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			recorder.Code, http.StatusCreated)
	}

	userBody, _ := json.Marshal(map[string]string{
		"firstName": "Karan",
		"lastName":  "Nehra",
		"email":     "karan",
		"password":  "karan",
		"role":      "MASTER",
	})

	req, err = http.NewRequest("POST", "/user", bytes.NewBuffer(userBody))
	recorder = httptest.NewRecorder()
	handler = http.HandlerFunc(CreateUser)
	handler.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			recorder.Code, http.StatusCreated)
	}
}
