package lab

type LampProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewLampProfileGuide() LampProfileGuide {
	return LampProfileGuide{ID: "lamp-guide", Label: "lamp linearity", Description: "lamp emission response", ReferenceCenter: 48.0, Tolerance: 1.1, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g LampProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g LampProfileGuide) Width() float64      { return g.Tolerance }
func (g LampProfileGuide) Name() string        { return g.Label }
func (g LampProfileGuide) Detail() string      { return g.Description }
func (g LampProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g LampProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g LampProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g LampProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g LampProfileGuide) Distance(value float64) float64 { return Absolute(value - g.ReferenceCenter) }
func (g LampProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g LampProfileGuide) UnitLabel() string { return g.Unit }
func (g LampProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g LampProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g LampProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g LampProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g LampProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const LampReferenceBand01 = 46.936667
const LampReferenceBand02 = 46.973333
const LampReferenceBand03 = 47.010000
const LampReferenceBand04 = 47.046667
const LampReferenceBand05 = 47.083333
const LampReferenceBand06 = 47.120000
const LampReferenceBand07 = 47.156667
const LampReferenceBand08 = 47.193333
const LampReferenceBand09 = 47.230000
const LampReferenceBand10 = 47.266667
const LampReferenceBand11 = 47.303333
const LampReferenceBand12 = 47.340000
const LampReferenceBand13 = 47.376667
const LampReferenceBand14 = 47.413333
const LampReferenceBand15 = 47.450000
const LampReferenceBand16 = 47.486667
const LampReferenceBand17 = 47.523333
const LampReferenceBand18 = 47.560000
const LampReferenceBand19 = 47.596667
const LampReferenceBand20 = 47.633333
const LampReferenceBand21 = 47.670000
const LampReferenceBand22 = 47.706667
const LampReferenceBand23 = 47.743333
const LampReferenceBand24 = 47.780000
const LampReferenceBand25 = 47.816667
const LampReferenceBand26 = 47.853333
const LampReferenceBand27 = 47.890000
const LampReferenceBand28 = 47.926667
const LampReferenceBand29 = 47.963333
const LampReferenceBand30 = 48.000000
const LampReferenceBand31 = 48.036667
const LampReferenceBand32 = 48.073333
const LampReferenceBand33 = 48.110000
const LampReferenceBand34 = 48.146667
const LampReferenceBand35 = 48.183333
const LampReferenceBand36 = 48.220000
const LampReferenceBand37 = 48.256667
const LampReferenceBand38 = 48.293333
const LampReferenceBand39 = 48.330000
const LampReferenceBand40 = 48.366667
const LampReferenceBand41 = 48.403333
const LampReferenceBand42 = 48.440000
const LampReferenceBand43 = 48.476667
const LampReferenceBand44 = 48.513333
const LampReferenceBand45 = 48.550000
const LampReferenceBand46 = 48.586667
const LampReferenceBand47 = 48.623333
const LampReferenceBand48 = 48.660000
const LampReferenceBand49 = 48.696667
const LampReferenceBand50 = 48.733333
const LampReferenceBand51 = 48.770000
const LampReferenceBand52 = 48.806667
const LampReferenceBand53 = 48.843333
const LampReferenceBand54 = 48.880000
const LampReferenceBand55 = 48.916667
const LampReferenceBand56 = 48.953333
const LampReferenceBand57 = 48.990000
const LampReferenceBand58 = 49.026667
const LampReferenceBand59 = 49.063333
const LampReferenceBand60 = 49.100000
