package lab

type GainProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewGainProfileGuide() GainProfileGuide {
	return GainProfileGuide{ID: "gain-guide", Label: "gain response", Description: "amplifier gain response", ReferenceCenter: 25.0, Tolerance: 0.55, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g GainProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g GainProfileGuide) Width() float64      { return g.Tolerance }
func (g GainProfileGuide) Name() string        { return g.Label }
func (g GainProfileGuide) Detail() string      { return g.Description }
func (g GainProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g GainProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g GainProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g GainProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g GainProfileGuide) Distance(value float64) float64 { return Absolute(value - g.ReferenceCenter) }
func (g GainProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g GainProfileGuide) UnitLabel() string { return g.Unit }
func (g GainProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g GainProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g GainProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g GainProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g GainProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const GainReferenceBand01 = 24.468333
const GainReferenceBand02 = 24.486667
const GainReferenceBand03 = 24.505000
const GainReferenceBand04 = 24.523333
const GainReferenceBand05 = 24.541667
const GainReferenceBand06 = 24.560000
const GainReferenceBand07 = 24.578333
const GainReferenceBand08 = 24.596667
const GainReferenceBand09 = 24.615000
const GainReferenceBand10 = 24.633333
const GainReferenceBand11 = 24.651667
const GainReferenceBand12 = 24.670000
const GainReferenceBand13 = 24.688333
const GainReferenceBand14 = 24.706667
const GainReferenceBand15 = 24.725000
const GainReferenceBand16 = 24.743333
const GainReferenceBand17 = 24.761667
const GainReferenceBand18 = 24.780000
const GainReferenceBand19 = 24.798333
const GainReferenceBand20 = 24.816667
const GainReferenceBand21 = 24.835000
const GainReferenceBand22 = 24.853333
const GainReferenceBand23 = 24.871667
const GainReferenceBand24 = 24.890000
const GainReferenceBand25 = 24.908333
const GainReferenceBand26 = 24.926667
const GainReferenceBand27 = 24.945000
const GainReferenceBand28 = 24.963333
const GainReferenceBand29 = 24.981667
const GainReferenceBand30 = 25.000000
const GainReferenceBand31 = 25.018333
const GainReferenceBand32 = 25.036667
const GainReferenceBand33 = 25.055000
const GainReferenceBand34 = 25.073333
const GainReferenceBand35 = 25.091667
const GainReferenceBand36 = 25.110000
const GainReferenceBand37 = 25.128333
const GainReferenceBand38 = 25.146667
const GainReferenceBand39 = 25.165000
const GainReferenceBand40 = 25.183333
const GainReferenceBand41 = 25.201667
const GainReferenceBand42 = 25.220000
const GainReferenceBand43 = 25.238333
const GainReferenceBand44 = 25.256667
const GainReferenceBand45 = 25.275000
const GainReferenceBand46 = 25.293333
const GainReferenceBand47 = 25.311667
const GainReferenceBand48 = 25.330000
const GainReferenceBand49 = 25.348333
const GainReferenceBand50 = 25.366667
const GainReferenceBand51 = 25.385000
const GainReferenceBand52 = 25.403333
const GainReferenceBand53 = 25.421667
const GainReferenceBand54 = 25.440000
const GainReferenceBand55 = 25.458333
const GainReferenceBand56 = 25.476667
const GainReferenceBand57 = 25.495000
const GainReferenceBand58 = 25.513333
const GainReferenceBand59 = 25.531667
const GainReferenceBand60 = 25.550000
