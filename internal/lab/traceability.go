package lab

import "sync"

type TraceabilityRecord struct {
	RunID             string
	ProfileID         string
	MethodID          string
	InstrumentID      string
	OperatorID        string
	BatchID           string
	SampleID          string
	ReferenceID       string
	OpticID           string
	DetectorID        string
	LampID            string
	FilterID          string
	RoomID            string
	SessionID         string
	RevisionID        string
	ReviewID          string
	ArchiveID         string
	SignatureID       string
	SourceID          string
	SequenceID        string
	StationID         string
	ChannelID         string
	SpectrumID        string
	WindowID          string
	SeriesID          string
	Purpose           string
	MethodVersion     string
	InstrumentVersion string
	ReferenceVersion  string
	CaptureMode       string
	ReviewState       string
	ReleaseState      string
	RetentionClass    string
	EvidenceClass     string
	LocationCode      string
	TimeZone          string
	Checksum          string
	OperatorNote      string
	ReviewNote        string
	ReleaseNote       string
}

type TraceabilityIndex struct {
	mu        sync.RWMutex
	Records   []TraceabilityRecord
	ByRun     map[string]TraceabilityRecord
	ByProfile map[string][]TraceabilityRecord
	Revision  int
}

func NewTraceabilityIndex() *TraceabilityIndex {
	return &TraceabilityIndex{ByRun: make(map[string]TraceabilityRecord), ByProfile: make(map[string][]TraceabilityRecord)}
}

func (i *TraceabilityIndex) Add(record TraceabilityRecord) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.Records = append(i.Records, record)
	i.ByRun[record.RunID] = record
	i.ByProfile[record.ProfileID] = append(i.ByProfile[record.ProfileID], record)
	i.Revision++
}

func (i *TraceabilityIndex) FindRun(runID string) (TraceabilityRecord, bool) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	record, ok := i.ByRun[runID]
	return record, ok
}

func (i *TraceabilityIndex) ProfileRecords(profileID string) []TraceabilityRecord {
	i.mu.RLock()
	defer i.mu.RUnlock()
	source := i.ByProfile[profileID]
	result := make([]TraceabilityRecord, len(source))
	copy(result, source)
	return result
}

func TraceabilityForRun(run Run) TraceabilityRecord {
	return TraceabilityRecord{
		RunID:          run.ID,
		ProfileID:      run.ProfileID,
		MethodID:       "spectral-calibration",
		OperatorID:     run.CapturedBy,
		BatchID:        run.CapturedAt.Format("20060102"),
		ReviewState:    run.State,
		ReleaseState:   FormatState(run.State),
		RetentionClass: "research",
		EvidenceClass:  "observation",
		CaptureMode:    "in-memory",
	}
}

func TraceabilityLabel(record TraceabilityRecord) string {
	return record.ProfileID + ":" + record.RunID + ":" + record.ReviewState
}

func TraceabilityComplete(record TraceabilityRecord) bool {
	return record.RunID != "" && record.ProfileID != "" && record.MethodID != ""
}
