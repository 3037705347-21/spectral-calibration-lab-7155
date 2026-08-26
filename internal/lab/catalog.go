package lab

import (
	"fmt"
	"sort"
)

type Catalog struct{ profiles map[string]Profile }

func NewCatalog() *Catalog {
	records := []Profile{
		{ID: "thermal-stability", Name: "Thermal stability", ReferenceCenter: 10, Tolerance: 0.45, MinimumSamples: 3, AcceptScore: 82, ReviewScore: 60, Unit: "nm", Description: "Checks thermal drift around the reference center."},
		{ID: "lamp-linearity", Name: "Lamp linearity", ReferenceCenter: 48, Tolerance: 1.1, MinimumSamples: 4, AcceptScore: 80, ReviewScore: 58, Unit: "mA", Description: "Checks repeatable line response across a lamp sequence."},
		{ID: "detector-noise", Name: "Detector noise", ReferenceCenter: 2.5, Tolerance: 0.3, MinimumSamples: 5, AcceptScore: 78, ReviewScore: 55, Unit: "counts", Description: "Checks noise spread in a detector baseline."},
		{ID: "filter-alignment", Name: "Filter alignment", ReferenceCenter: 72, Tolerance: 0.8, MinimumSamples: 3, AcceptScore: 84, ReviewScore: 63, Unit: "degrees", Description: "Checks an optical filter alignment reading."},
	}
	records = append(records, AllProfileGuides()...)
	profiles := make(map[string]Profile, len(records))
	for _, record := range records {
		profiles[record.ID] = record
	}
	return &Catalog{profiles: profiles}
}

func (c *Catalog) List() []Profile {
	result := make([]Profile, 0, len(c.profiles))
	for _, profile := range c.profiles {
		result = append(result, profile)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result
}

func (c *Catalog) Find(id string) (Profile, error) {
	profile, ok := c.profiles[id]
	if !ok {
		return Profile{}, fmt.Errorf("unknown calibration profile %q", id)
	}
	return profile, nil
}

func (c *Catalog) Count() int { return len(c.profiles) }
