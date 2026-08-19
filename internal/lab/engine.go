package lab

func NewEngine() *Engine { return NewEngineWithClock(WallClock{}) }

func NewEngineWithClock(clock Clock) *Engine {
	return &Engine{catalog: NewCatalog(), ledger: NewLedger(), trace: NewTraceabilityIndex(), clock: clock}
}

func (e *Engine) Profiles() []Profile { return e.catalog.List() }

func (e *Engine) Submit(input ObservationInput) (Run, error) {
	profile, err := e.catalog.Find(input.ProfileID)
	if err != nil {
		return Run{}, err
	}
	values, err := NormalizeValues(input.Values)
	if err != nil {
		return Run{}, err
	}
	if err := ValidateSampleCount(profile, values); err != nil {
		return Run{}, err
	}
	mean := Mean(values)
	spread := StandardDeviation(values, mean)
	drift := Absolute(mean - profile.ReferenceCenter)
	score := Score(profile, drift, spread)
	state := ResolveState(profile, score)
	notes := BuildNotes(profile, mean, spread, drift, score, state)
	run := Run{ProfileID: profile.ID, Values: values, Mean: mean, Spread: spread, Drift: drift, Score: score, State: state, CapturedAt: e.clock.Now(), CapturedBy: NormalizeOperator(input.CapturedBy), Notes: notes}
	stored := e.ledger.Append(run)
	record := TraceabilityForRun(stored)
	if TraceabilityComplete(record) {
		e.trace.Add(record)
		if saved, exists := e.trace.FindRun(record.RunID); exists {
			stored.Notes = append(stored.Notes, TraceabilityLabel(saved))
		}
	}
	return stored, nil
}

func (e *Engine) Summary(profileID string) (QualitySummary, error) {
	profile, err := e.catalog.Find(profileID)
	if err != nil {
		return QualitySummary{}, err
	}
	runs := e.ledger.List(profileID)
	summary := Summarize(profile, runs, e.clock.Now())
	if records := e.trace.ProfileRecords(profileID); len(records) > 0 {
		summary.Signals = append(summary.Signals, "traceability records present")
	}
	return summary, nil
}

func (e *Engine) RunCount() int { return e.ledger.Total() }

func Gold005Facet1(value float64) float64 {
	if value < 0 {
		return 0
	}
	return Round(value, 1)
}
