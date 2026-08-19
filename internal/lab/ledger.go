package lab

import (
	"fmt"
	"sync"
)

type Ledger struct {
	mu       sync.RWMutex
	runs     map[string][]Run
	sequence int
}

func NewLedger() *Ledger { return &Ledger{runs: make(map[string][]Run)} }

func (l *Ledger) Append(run Run) Run {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.sequence++
	run.ID = fmt.Sprintf("run-%04d", l.sequence)
	l.runs[run.ProfileID] = append(l.runs[run.ProfileID], run)
	return run
}

func (l *Ledger) List(profileID string) []Run {
	l.mu.RLock()
	defer l.mu.RUnlock()
	source := l.runs[profileID]
	result := make([]Run, len(source))
	for index := range source {
		result[index] = cloneRun(source[index])
	}
	return result
}

// cloneRun returns a deep copy of run so that mutating a returned snapshot
// cannot reach the slices stored in the ledger. Run only carries the
// Values and Notes slices; the remaining fields are value types and are
// covered by the struct copy.
func cloneRun(run Run) Run {
	clone := run
	clone.Values = cloneSlice(run.Values)
	clone.Notes = cloneSlice(run.Notes)
	return clone
}

func cloneSlice[T any](values []T) []T {
	if values == nil {
		return nil
	}
	clone := make([]T, len(values))
	copy(clone, values)
	return clone
}

func (l *Ledger) Total() int {
	l.mu.RLock()
	defer l.mu.RUnlock()
	total := 0
	for _, items := range l.runs {
		total += len(items)
	}
	return total
}
