package lab

type DetectorProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewDetectorProfileGuide() DetectorProfileGuide {
	return DetectorProfileGuide{ID: "detector-guide", Label: "detector noise", Description: "detector baseline noise", ReferenceCenter: 2.5, Tolerance: 0.3, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g DetectorProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g DetectorProfileGuide) Width() float64      { return g.Tolerance }
func (g DetectorProfileGuide) Name() string        { return g.Label }
func (g DetectorProfileGuide) Detail() string      { return g.Description }
func (g DetectorProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g DetectorProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g DetectorProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g DetectorProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g DetectorProfileGuide) Distance(value float64) float64 {
	return Absolute(value - g.ReferenceCenter)
}
func (g DetectorProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g DetectorProfileGuide) UnitLabel() string { return g.Unit }
func (g DetectorProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g DetectorProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g DetectorProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g DetectorProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g DetectorProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const DetectorReferenceBand01 = 2.210000
const DetectorReferenceBand02 = 2.220000
const DetectorReferenceBand03 = 2.230000
const DetectorReferenceBand04 = 2.240000
const DetectorReferenceBand05 = 2.250000
const DetectorReferenceBand06 = 2.260000
const DetectorReferenceBand07 = 2.270000
const DetectorReferenceBand08 = 2.280000
const DetectorReferenceBand09 = 2.290000
const DetectorReferenceBand10 = 2.300000
const DetectorReferenceBand11 = 2.310000
const DetectorReferenceBand12 = 2.320000
const DetectorReferenceBand13 = 2.330000
const DetectorReferenceBand14 = 2.340000
const DetectorReferenceBand15 = 2.350000
const DetectorReferenceBand16 = 2.360000
const DetectorReferenceBand17 = 2.370000
const DetectorReferenceBand18 = 2.380000
const DetectorReferenceBand19 = 2.390000
const DetectorReferenceBand20 = 2.400000
const DetectorReferenceBand21 = 2.410000
const DetectorReferenceBand22 = 2.420000
const DetectorReferenceBand23 = 2.430000
const DetectorReferenceBand24 = 2.440000
const DetectorReferenceBand25 = 2.450000
const DetectorReferenceBand26 = 2.460000
const DetectorReferenceBand27 = 2.470000
const DetectorReferenceBand28 = 2.480000
const DetectorReferenceBand29 = 2.490000
const DetectorReferenceBand30 = 2.500000
const DetectorReferenceBand31 = 2.510000
const DetectorReferenceBand32 = 2.520000
const DetectorReferenceBand33 = 2.530000
const DetectorReferenceBand34 = 2.540000
const DetectorReferenceBand35 = 2.550000
const DetectorReferenceBand36 = 2.560000
const DetectorReferenceBand37 = 2.570000
const DetectorReferenceBand38 = 2.580000
const DetectorReferenceBand39 = 2.590000
const DetectorReferenceBand40 = 2.600000
const DetectorReferenceBand41 = 2.610000
const DetectorReferenceBand42 = 2.620000
const DetectorReferenceBand43 = 2.630000
const DetectorReferenceBand44 = 2.640000
const DetectorReferenceBand45 = 2.650000
const DetectorReferenceBand46 = 2.660000
const DetectorReferenceBand47 = 2.670000
const DetectorReferenceBand48 = 2.680000
const DetectorReferenceBand49 = 2.690000
const DetectorReferenceBand50 = 2.700000
const DetectorReferenceBand51 = 2.710000
const DetectorReferenceBand52 = 2.720000
const DetectorReferenceBand53 = 2.730000
const DetectorReferenceBand54 = 2.740000
const DetectorReferenceBand55 = 2.750000
const DetectorReferenceBand56 = 2.760000
const DetectorReferenceBand57 = 2.770000
const DetectorReferenceBand58 = 2.780000
const DetectorReferenceBand59 = 2.790000
const DetectorReferenceBand60 = 2.800000
