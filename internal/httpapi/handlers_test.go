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

func TestReportRecommendsAcceptanceForStableRun(t *testing.T) {
	server := New(lab.NewEngine())
	create := httptest.NewRequest(http.MethodPost, "/v1/observations", bytes.NewBufferString(`{"profile_id":"thermal-stability","values":[10,10.1,9.9]}`))
	server.Handler().ServeHTTP(httptest.NewRecorder(), create)
	req := httptest.NewRequest(http.MethodGet, "/v1/reports?profile_id=thermal-stability", nil)
	rr := httptest.NewRecorder()
	server.Handler().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d", rr.Code)
	}
	if !bytes.Contains(rr.Body.Bytes(), []byte(`"recommended_action":"accept the current calibration run"`)) {
		t.Fatalf("body = %s", rr.Body.String())
	}
}
