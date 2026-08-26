package httpapi

import (
	"bytes"
	"encoding/json"
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

func TestObservationEndpointReportsRunOutcome(t *testing.T) {
	server := New(lab.NewEngine())
	request := httptest.NewRequest(http.MethodPost, "/v1/observations", bytes.NewBufferString(`{"profile_id":"thermal-stability","values":[20,20,20]}`))
	recorder := httptest.NewRecorder()
	server.Handler().ServeHTTP(recorder, request)
	if recorder.Code != http.StatusCreated {
		t.Fatalf("status = %d body=%s", recorder.Code, recorder.Body.String())
	}

	var payload struct {
		Status string  `json:"status"`
		Run    lab.Run `json:"run"`
	}
	if err := json.NewDecoder(recorder.Body).Decode(&payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload.Run.State != "repeat" {
		t.Fatalf("run state = %s, want repeat", payload.Run.State)
	}
	if payload.Status != payload.Run.State {
		t.Fatalf("top-level status = %s, want %s", payload.Status, payload.Run.State)
	}
}
