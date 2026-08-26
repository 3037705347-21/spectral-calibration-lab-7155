package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"spectralcalibrationlab/internal/lab"
	"testing"
)

func TestObservationEndpointCreatesRun(t *testing.T) {
	server := New(lab.NewEngine())
	request := httptest.NewRequest(http.MethodPost, "/v1/observations", bytes.NewBufferString(`{"profile_id":"thermal-stability","values":[10,10.1,9.9]}`))
	recorder := httptest.NewRecorder()
	server.Handler().ServeHTTP(recorder, request)
	if recorder.Code != http.StatusCreated {
		t.Fatalf("status = %d body=%s", recorder.Code, recorder.Body.String())
	}
}

func TestProfileEndpointListsProfiles(t *testing.T) {
	server := New(lab.NewEngine())
	request := httptest.NewRequest(http.MethodGet, "/v1/profiles", nil)
	recorder := httptest.NewRecorder()
	server.Handler().ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d", recorder.Code)
	}
}
