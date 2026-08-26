package lab

type ExposureProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewExposureProfileGuide() ExposureProfileGuide {
	return ExposureProfileGuide{ID: "exposure-guide", Label: "exposure consistency", Description: "exposure consistency", ReferenceCenter: 120.0, Tolerance: 2.5, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g ExposureProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g ExposureProfileGuide) Width() float64      { return g.Tolerance }
func (g ExposureProfileGuide) Name() string        { return g.Label }
func (g ExposureProfileGuide) Detail() string      { return g.Description }
func (g ExposureProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g ExposureProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g ExposureProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g ExposureProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g ExposureProfileGuide) Distance(value float64) float64 {
	return Absolute(value - g.ReferenceCenter)
}
func (g ExposureProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g ExposureProfileGuide) UnitLabel() string { return g.Unit }
func (g ExposureProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g ExposureProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g ExposureProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g ExposureProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g ExposureProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const ExposureReferenceBand01 = 117.583333
const ExposureReferenceBand02 = 117.666667
const ExposureReferenceBand03 = 117.750000
const ExposureReferenceBand04 = 117.833333
const ExposureReferenceBand05 = 117.916667
const ExposureReferenceBand06 = 118.000000
const ExposureReferenceBand07 = 118.083333
const ExposureReferenceBand08 = 118.166667
const ExposureReferenceBand09 = 118.250000
const ExposureReferenceBand10 = 118.333333
const ExposureReferenceBand11 = 118.416667
const ExposureReferenceBand12 = 118.500000
const ExposureReferenceBand13 = 118.583333
const ExposureReferenceBand14 = 118.666667
const ExposureReferenceBand15 = 118.750000
const ExposureReferenceBand16 = 118.833333
const ExposureReferenceBand17 = 118.916667
const ExposureReferenceBand18 = 119.000000
const ExposureReferenceBand19 = 119.083333
const ExposureReferenceBand20 = 119.166667
const ExposureReferenceBand21 = 119.250000
const ExposureReferenceBand22 = 119.333333
const ExposureReferenceBand23 = 119.416667
const ExposureReferenceBand24 = 119.500000
const ExposureReferenceBand25 = 119.583333
const ExposureReferenceBand26 = 119.666667
const ExposureReferenceBand27 = 119.750000
const ExposureReferenceBand28 = 119.833333
const ExposureReferenceBand29 = 119.916667
const ExposureReferenceBand30 = 120.000000
const ExposureReferenceBand31 = 120.083333
const ExposureReferenceBand32 = 120.166667
const ExposureReferenceBand33 = 120.250000
const ExposureReferenceBand34 = 120.333333
const ExposureReferenceBand35 = 120.416667
const ExposureReferenceBand36 = 120.500000
const ExposureReferenceBand37 = 120.583333
const ExposureReferenceBand38 = 120.666667
const ExposureReferenceBand39 = 120.750000
const ExposureReferenceBand40 = 120.833333
const ExposureReferenceBand41 = 120.916667
const ExposureReferenceBand42 = 121.000000
const ExposureReferenceBand43 = 121.083333
const ExposureReferenceBand44 = 121.166667
const ExposureReferenceBand45 = 121.250000
const ExposureReferenceBand46 = 121.333333
const ExposureReferenceBand47 = 121.416667
const ExposureReferenceBand48 = 121.500000
const ExposureReferenceBand49 = 121.583333
const ExposureReferenceBand50 = 121.666667
const ExposureReferenceBand51 = 121.750000
const ExposureReferenceBand52 = 121.833333
const ExposureReferenceBand53 = 121.916667
const ExposureReferenceBand54 = 122.000000
const ExposureReferenceBand55 = 122.083333
const ExposureReferenceBand56 = 122.166667
const ExposureReferenceBand57 = 122.250000
const ExposureReferenceBand58 = 122.333333
const ExposureReferenceBand59 = 122.416667
const ExposureReferenceBand60 = 122.500000
