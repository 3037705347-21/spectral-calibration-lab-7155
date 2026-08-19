package lab

func cloneRun(run Run) Run {
	run.Values = append([]float64(nil), run.Values...)
	run.Notes = append([]string(nil), run.Notes...)
	return run
}

func cloneRuns(runs []Run) []Run {
	if len(runs) == 0 {
		return nil
	}
	result := make([]Run, len(runs))
	for index, run := range runs {
		result[index] = cloneRun(run)
	}
	return result
}

func cloneRunForStorage(run Run) Run {
	return cloneRun(run)
}

func cloneRunForHistory(run Run) Run {
	return cloneRun(run)
}
