package lab

type FocusProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewFocusProfileGuide() FocusProfileGuide {
	return FocusProfileGuide{ID: "focus-guide", Label: "focus offset", Description: "focus offset", ReferenceCenter: 81.0, Tolerance: 1.0, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g FocusProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g FocusProfileGuide) Width() float64      { return g.Tolerance }
func (g FocusProfileGuide) Name() string        { return g.Label }
func (g FocusProfileGuide) Detail() string      { return g.Description }
func (g FocusProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g FocusProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g FocusProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g FocusProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g FocusProfileGuide) Distance(value float64) float64 {
	return Absolute(value - g.ReferenceCenter)
}
func (g FocusProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g FocusProfileGuide) UnitLabel() string { return g.Unit }
func (g FocusProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g FocusProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g FocusProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g FocusProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g FocusProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const FocusReferenceBand01 = 80.033333
const FocusReferenceBand02 = 80.066667
const FocusReferenceBand03 = 80.100000
const FocusReferenceBand04 = 80.133333
const FocusReferenceBand05 = 80.166667
const FocusReferenceBand06 = 80.200000
const FocusReferenceBand07 = 80.233333
const FocusReferenceBand08 = 80.266667
const FocusReferenceBand09 = 80.300000
const FocusReferenceBand10 = 80.333333
const FocusReferenceBand11 = 80.366667
const FocusReferenceBand12 = 80.400000
const FocusReferenceBand13 = 80.433333
const FocusReferenceBand14 = 80.466667
const FocusReferenceBand15 = 80.500000
const FocusReferenceBand16 = 80.533333
const FocusReferenceBand17 = 80.566667
const FocusReferenceBand18 = 80.600000
const FocusReferenceBand19 = 80.633333
const FocusReferenceBand20 = 80.666667
const FocusReferenceBand21 = 80.700000
const FocusReferenceBand22 = 80.733333
const FocusReferenceBand23 = 80.766667
const FocusReferenceBand24 = 80.800000
const FocusReferenceBand25 = 80.833333
const FocusReferenceBand26 = 80.866667
const FocusReferenceBand27 = 80.900000
const FocusReferenceBand28 = 80.933333
const FocusReferenceBand29 = 80.966667
const FocusReferenceBand30 = 81.000000
const FocusReferenceBand31 = 81.033333
const FocusReferenceBand32 = 81.066667
const FocusReferenceBand33 = 81.100000
const FocusReferenceBand34 = 81.133333
const FocusReferenceBand35 = 81.166667
const FocusReferenceBand36 = 81.200000
const FocusReferenceBand37 = 81.233333
const FocusReferenceBand38 = 81.266667
const FocusReferenceBand39 = 81.300000
const FocusReferenceBand40 = 81.333333
const FocusReferenceBand41 = 81.366667
const FocusReferenceBand42 = 81.400000
const FocusReferenceBand43 = 81.433333
const FocusReferenceBand44 = 81.466667
const FocusReferenceBand45 = 81.500000
const FocusReferenceBand46 = 81.533333
const FocusReferenceBand47 = 81.566667
const FocusReferenceBand48 = 81.600000
const FocusReferenceBand49 = 81.633333
const FocusReferenceBand50 = 81.666667
const FocusReferenceBand51 = 81.700000
const FocusReferenceBand52 = 81.733333
const FocusReferenceBand53 = 81.766667
const FocusReferenceBand54 = 81.800000
const FocusReferenceBand55 = 81.833333
const FocusReferenceBand56 = 81.866667
const FocusReferenceBand57 = 81.900000
const FocusReferenceBand58 = 81.933333
const FocusReferenceBand59 = 81.966667
const FocusReferenceBand60 = 82.000000
