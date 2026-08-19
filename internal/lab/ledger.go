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
	stored := cloneRunForStorage(run)
	l.runs[run.ProfileID] = append(l.runs[run.ProfileID], stored)
	return cloneRunForHistory(stored)
}

func (l *Ledger) List(profileID string) []Run {
	l.mu.RLock()
	defer l.mu.RUnlock()
	source := l.runs[profileID]
	return cloneRuns(source)
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
