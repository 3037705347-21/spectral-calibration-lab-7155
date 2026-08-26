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
	copy(result, source)
	return result
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
