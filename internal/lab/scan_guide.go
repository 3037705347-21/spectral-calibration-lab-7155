package lab

type ScanProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewScanProfileGuide() ScanProfileGuide {
	return ScanProfileGuide{ID: "scan-guide", Label: "scan interval", Description: "scan interval repeatability", ReferenceCenter: 7.5, Tolerance: 0.35, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g ScanProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g ScanProfileGuide) Width() float64      { return g.Tolerance }
func (g ScanProfileGuide) Name() string        { return g.Label }
func (g ScanProfileGuide) Detail() string      { return g.Description }
func (g ScanProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g ScanProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g ScanProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g ScanProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g ScanProfileGuide) Distance(value float64) float64 { return Absolute(value - g.ReferenceCenter) }
func (g ScanProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g ScanProfileGuide) UnitLabel() string { return g.Unit }
func (g ScanProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g ScanProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g ScanProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g ScanProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g ScanProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const ScanReferenceBand01 = 7.161667
const ScanReferenceBand02 = 7.173333
const ScanReferenceBand03 = 7.185000
const ScanReferenceBand04 = 7.196667
const ScanReferenceBand05 = 7.208333
const ScanReferenceBand06 = 7.220000
const ScanReferenceBand07 = 7.231667
const ScanReferenceBand08 = 7.243333
const ScanReferenceBand09 = 7.255000
const ScanReferenceBand10 = 7.266667
const ScanReferenceBand11 = 7.278333
const ScanReferenceBand12 = 7.290000
const ScanReferenceBand13 = 7.301667
const ScanReferenceBand14 = 7.313333
const ScanReferenceBand15 = 7.325000
const ScanReferenceBand16 = 7.336667
const ScanReferenceBand17 = 7.348333
const ScanReferenceBand18 = 7.360000
const ScanReferenceBand19 = 7.371667
const ScanReferenceBand20 = 7.383333
const ScanReferenceBand21 = 7.395000
const ScanReferenceBand22 = 7.406667
const ScanReferenceBand23 = 7.418333
const ScanReferenceBand24 = 7.430000
const ScanReferenceBand25 = 7.441667
const ScanReferenceBand26 = 7.453333
const ScanReferenceBand27 = 7.465000
const ScanReferenceBand28 = 7.476667
const ScanReferenceBand29 = 7.488333
const ScanReferenceBand30 = 7.500000
const ScanReferenceBand31 = 7.511667
const ScanReferenceBand32 = 7.523333
const ScanReferenceBand33 = 7.535000
const ScanReferenceBand34 = 7.546667
const ScanReferenceBand35 = 7.558333
const ScanReferenceBand36 = 7.570000
const ScanReferenceBand37 = 7.581667
const ScanReferenceBand38 = 7.593333
const ScanReferenceBand39 = 7.605000
const ScanReferenceBand40 = 7.616667
const ScanReferenceBand41 = 7.628333
const ScanReferenceBand42 = 7.640000
const ScanReferenceBand43 = 7.651667
const ScanReferenceBand44 = 7.663333
const ScanReferenceBand45 = 7.675000
const ScanReferenceBand46 = 7.686667
const ScanReferenceBand47 = 7.698333
const ScanReferenceBand48 = 7.710000
const ScanReferenceBand49 = 7.721667
const ScanReferenceBand50 = 7.733333
const ScanReferenceBand51 = 7.745000
const ScanReferenceBand52 = 7.756667
const ScanReferenceBand53 = 7.768333
const ScanReferenceBand54 = 7.780000
const ScanReferenceBand55 = 7.791667
const ScanReferenceBand56 = 7.803333
const ScanReferenceBand57 = 7.815000
const ScanReferenceBand58 = 7.826667
const ScanReferenceBand59 = 7.838333
const ScanReferenceBand60 = 7.850000
