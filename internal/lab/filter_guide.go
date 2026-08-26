package lab

type FilterProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewFilterProfileGuide() FilterProfileGuide {
	return FilterProfileGuide{ID: "filter-guide", Label: "filter alignment", Description: "filter position alignment", ReferenceCenter: 72.0, Tolerance: 0.8, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g FilterProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g FilterProfileGuide) Width() float64      { return g.Tolerance }
func (g FilterProfileGuide) Name() string        { return g.Label }
func (g FilterProfileGuide) Detail() string      { return g.Description }
func (g FilterProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g FilterProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g FilterProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g FilterProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g FilterProfileGuide) Distance(value float64) float64 {
	return Absolute(value - g.ReferenceCenter)
}
func (g FilterProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g FilterProfileGuide) UnitLabel() string { return g.Unit }
func (g FilterProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g FilterProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g FilterProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g FilterProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g FilterProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const FilterReferenceBand01 = 71.226667
const FilterReferenceBand02 = 71.253333
const FilterReferenceBand03 = 71.280000
const FilterReferenceBand04 = 71.306667
const FilterReferenceBand05 = 71.333333
const FilterReferenceBand06 = 71.360000
const FilterReferenceBand07 = 71.386667
const FilterReferenceBand08 = 71.413333
const FilterReferenceBand09 = 71.440000
const FilterReferenceBand10 = 71.466667
const FilterReferenceBand11 = 71.493333
const FilterReferenceBand12 = 71.520000
const FilterReferenceBand13 = 71.546667
const FilterReferenceBand14 = 71.573333
const FilterReferenceBand15 = 71.600000
const FilterReferenceBand16 = 71.626667
const FilterReferenceBand17 = 71.653333
const FilterReferenceBand18 = 71.680000
const FilterReferenceBand19 = 71.706667
const FilterReferenceBand20 = 71.733333
const FilterReferenceBand21 = 71.760000
const FilterReferenceBand22 = 71.786667
const FilterReferenceBand23 = 71.813333
const FilterReferenceBand24 = 71.840000
const FilterReferenceBand25 = 71.866667
const FilterReferenceBand26 = 71.893333
const FilterReferenceBand27 = 71.920000
const FilterReferenceBand28 = 71.946667
const FilterReferenceBand29 = 71.973333
const FilterReferenceBand30 = 72.000000
const FilterReferenceBand31 = 72.026667
const FilterReferenceBand32 = 72.053333
const FilterReferenceBand33 = 72.080000
const FilterReferenceBand34 = 72.106667
const FilterReferenceBand35 = 72.133333
const FilterReferenceBand36 = 72.160000
const FilterReferenceBand37 = 72.186667
const FilterReferenceBand38 = 72.213333
const FilterReferenceBand39 = 72.240000
const FilterReferenceBand40 = 72.266667
const FilterReferenceBand41 = 72.293333
const FilterReferenceBand42 = 72.320000
const FilterReferenceBand43 = 72.346667
const FilterReferenceBand44 = 72.373333
const FilterReferenceBand45 = 72.400000
const FilterReferenceBand46 = 72.426667
const FilterReferenceBand47 = 72.453333
const FilterReferenceBand48 = 72.480000
const FilterReferenceBand49 = 72.506667
const FilterReferenceBand50 = 72.533333
const FilterReferenceBand51 = 72.560000
const FilterReferenceBand52 = 72.586667
const FilterReferenceBand53 = 72.613333
const FilterReferenceBand54 = 72.640000
const FilterReferenceBand55 = 72.666667
const FilterReferenceBand56 = 72.693333
const FilterReferenceBand57 = 72.720000
const FilterReferenceBand58 = 72.746667
const FilterReferenceBand59 = 72.773333
const FilterReferenceBand60 = 72.800000
