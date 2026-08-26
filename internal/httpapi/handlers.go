package httpapi

import (
	"encoding/json"
	"net/http"
	"spectralcalibrationlab/internal/lab"
)

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]any{"status": "ok", "runs": s.engine.RunCount()})
}

func (s *Server) profiles(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]any{"profiles": s.engine.Profiles()})
}

func (s *Server) observations(w http.ResponseWriter, r *http.Request) {
	var input lab.ObservationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		WriteError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if err := lab.ValidateObservation(input); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	run, err := s.engine.Submit(input)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	WriteJSON(w, http.StatusCreated, map[string]any{"status": run.State, "run": run})
}

func (s *Server) reports(w http.ResponseWriter, r *http.Request) {
	profileID := r.URL.Query().Get("profile_id")
	if err := lab.ValidateProfileID(profileID); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	summary, err := s.engine.Summary(profileID)
	if err != nil {
		WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	WriteJSON(w, http.StatusOK, summary)
}
