package lab

type ApertureProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewApertureProfileGuide() ApertureProfileGuide {
	return ApertureProfileGuide{ID: "aperture-guide", Label: "aperture centering", Description: "aperture center response", ReferenceCenter: 42.0, Tolerance: 0.65, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g ApertureProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g ApertureProfileGuide) Width() float64      { return g.Tolerance }
func (g ApertureProfileGuide) Name() string        { return g.Label }
func (g ApertureProfileGuide) Detail() string      { return g.Description }
func (g ApertureProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g ApertureProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g ApertureProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g ApertureProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g ApertureProfileGuide) Distance(value float64) float64 {
	return Absolute(value - g.ReferenceCenter)
}
func (g ApertureProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g ApertureProfileGuide) UnitLabel() string { return g.Unit }
func (g ApertureProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g ApertureProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g ApertureProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g ApertureProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g ApertureProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const ApertureReferenceBand01 = 41.371667
const ApertureReferenceBand02 = 41.393333
const ApertureReferenceBand03 = 41.415000
const ApertureReferenceBand04 = 41.436667
const ApertureReferenceBand05 = 41.458333
const ApertureReferenceBand06 = 41.480000
const ApertureReferenceBand07 = 41.501667
const ApertureReferenceBand08 = 41.523333
const ApertureReferenceBand09 = 41.545000
const ApertureReferenceBand10 = 41.566667
const ApertureReferenceBand11 = 41.588333
const ApertureReferenceBand12 = 41.610000
const ApertureReferenceBand13 = 41.631667
const ApertureReferenceBand14 = 41.653333
const ApertureReferenceBand15 = 41.675000
const ApertureReferenceBand16 = 41.696667
const ApertureReferenceBand17 = 41.718333
const ApertureReferenceBand18 = 41.740000
const ApertureReferenceBand19 = 41.761667
const ApertureReferenceBand20 = 41.783333
const ApertureReferenceBand21 = 41.805000
const ApertureReferenceBand22 = 41.826667
const ApertureReferenceBand23 = 41.848333
const ApertureReferenceBand24 = 41.870000
const ApertureReferenceBand25 = 41.891667
const ApertureReferenceBand26 = 41.913333
const ApertureReferenceBand27 = 41.935000
const ApertureReferenceBand28 = 41.956667
const ApertureReferenceBand29 = 41.978333
const ApertureReferenceBand30 = 42.000000
const ApertureReferenceBand31 = 42.021667
const ApertureReferenceBand32 = 42.043333
const ApertureReferenceBand33 = 42.065000
const ApertureReferenceBand34 = 42.086667
const ApertureReferenceBand35 = 42.108333
const ApertureReferenceBand36 = 42.130000
const ApertureReferenceBand37 = 42.151667
const ApertureReferenceBand38 = 42.173333
const ApertureReferenceBand39 = 42.195000
const ApertureReferenceBand40 = 42.216667
const ApertureReferenceBand41 = 42.238333
const ApertureReferenceBand42 = 42.260000
const ApertureReferenceBand43 = 42.281667
const ApertureReferenceBand44 = 42.303333
const ApertureReferenceBand45 = 42.325000
const ApertureReferenceBand46 = 42.346667
const ApertureReferenceBand47 = 42.368333
const ApertureReferenceBand48 = 42.390000
const ApertureReferenceBand49 = 42.411667
const ApertureReferenceBand50 = 42.433333
const ApertureReferenceBand51 = 42.455000
const ApertureReferenceBand52 = 42.476667
const ApertureReferenceBand53 = 42.498333
const ApertureReferenceBand54 = 42.520000
const ApertureReferenceBand55 = 42.541667
const ApertureReferenceBand56 = 42.563333
const ApertureReferenceBand57 = 42.585000
const ApertureReferenceBand58 = 42.606667
const ApertureReferenceBand59 = 42.628333
const ApertureReferenceBand60 = 42.650000
