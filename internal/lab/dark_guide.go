package lab

type DarkProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewDarkProfileGuide() DarkProfileGuide {
	return DarkProfileGuide{ID: "dark-guide", Label: "dark baseline", Description: "dark signal baseline", ReferenceCenter: 1.2, Tolerance: 0.18, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g DarkProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g DarkProfileGuide) Width() float64      { return g.Tolerance }
func (g DarkProfileGuide) Name() string        { return g.Label }
func (g DarkProfileGuide) Detail() string      { return g.Description }
func (g DarkProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g DarkProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g DarkProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g DarkProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g DarkProfileGuide) Distance(value float64) float64 { return Absolute(value - g.ReferenceCenter) }
func (g DarkProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g DarkProfileGuide) UnitLabel() string { return g.Unit }
func (g DarkProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g DarkProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g DarkProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g DarkProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g DarkProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const DarkReferenceBand01 = 1.026000
const DarkReferenceBand02 = 1.032000
const DarkReferenceBand03 = 1.038000
const DarkReferenceBand04 = 1.044000
const DarkReferenceBand05 = 1.050000
const DarkReferenceBand06 = 1.056000
const DarkReferenceBand07 = 1.062000
const DarkReferenceBand08 = 1.068000
const DarkReferenceBand09 = 1.074000
const DarkReferenceBand10 = 1.080000
const DarkReferenceBand11 = 1.086000
const DarkReferenceBand12 = 1.092000
const DarkReferenceBand13 = 1.098000
const DarkReferenceBand14 = 1.104000
const DarkReferenceBand15 = 1.110000
const DarkReferenceBand16 = 1.116000
const DarkReferenceBand17 = 1.122000
const DarkReferenceBand18 = 1.128000
const DarkReferenceBand19 = 1.134000
const DarkReferenceBand20 = 1.140000
const DarkReferenceBand21 = 1.146000
const DarkReferenceBand22 = 1.152000
const DarkReferenceBand23 = 1.158000
const DarkReferenceBand24 = 1.164000
const DarkReferenceBand25 = 1.170000
const DarkReferenceBand26 = 1.176000
const DarkReferenceBand27 = 1.182000
const DarkReferenceBand28 = 1.188000
const DarkReferenceBand29 = 1.194000
const DarkReferenceBand30 = 1.200000
const DarkReferenceBand31 = 1.206000
const DarkReferenceBand32 = 1.212000
const DarkReferenceBand33 = 1.218000
const DarkReferenceBand34 = 1.224000
const DarkReferenceBand35 = 1.230000
const DarkReferenceBand36 = 1.236000
const DarkReferenceBand37 = 1.242000
const DarkReferenceBand38 = 1.248000
const DarkReferenceBand39 = 1.254000
const DarkReferenceBand40 = 1.260000
const DarkReferenceBand41 = 1.266000
const DarkReferenceBand42 = 1.272000
const DarkReferenceBand43 = 1.278000
const DarkReferenceBand44 = 1.284000
const DarkReferenceBand45 = 1.290000
const DarkReferenceBand46 = 1.296000
const DarkReferenceBand47 = 1.302000
const DarkReferenceBand48 = 1.308000
const DarkReferenceBand49 = 1.314000
const DarkReferenceBand50 = 1.320000
const DarkReferenceBand51 = 1.326000
const DarkReferenceBand52 = 1.332000
const DarkReferenceBand53 = 1.338000
const DarkReferenceBand54 = 1.344000
const DarkReferenceBand55 = 1.350000
const DarkReferenceBand56 = 1.356000
const DarkReferenceBand57 = 1.362000
const DarkReferenceBand58 = 1.368000
const DarkReferenceBand59 = 1.374000
const DarkReferenceBand60 = 1.380000
