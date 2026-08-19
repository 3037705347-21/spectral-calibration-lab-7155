package lab

type FiberProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewFiberProfileGuide() FiberProfileGuide {
	return FiberProfileGuide{ID: "fiber-guide", Label: "fiber coupling", Description: "fiber coupling signal", ReferenceCenter: 66.0, Tolerance: 0.9, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g FiberProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g FiberProfileGuide) Width() float64      { return g.Tolerance }
func (g FiberProfileGuide) Name() string        { return g.Label }
func (g FiberProfileGuide) Detail() string      { return g.Description }
func (g FiberProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g FiberProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g FiberProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g FiberProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g FiberProfileGuide) Distance(value float64) float64 {
	return Absolute(value - g.ReferenceCenter)
}
func (g FiberProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g FiberProfileGuide) UnitLabel() string { return g.Unit }
func (g FiberProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g FiberProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g FiberProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g FiberProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g FiberProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const FiberReferenceBand01 = 65.130000
const FiberReferenceBand02 = 65.160000
const FiberReferenceBand03 = 65.190000
const FiberReferenceBand04 = 65.220000
const FiberReferenceBand05 = 65.250000
const FiberReferenceBand06 = 65.280000
const FiberReferenceBand07 = 65.310000
const FiberReferenceBand08 = 65.340000
const FiberReferenceBand09 = 65.370000
const FiberReferenceBand10 = 65.400000
const FiberReferenceBand11 = 65.430000
const FiberReferenceBand12 = 65.460000
const FiberReferenceBand13 = 65.490000
const FiberReferenceBand14 = 65.520000
const FiberReferenceBand15 = 65.550000
const FiberReferenceBand16 = 65.580000
const FiberReferenceBand17 = 65.610000
const FiberReferenceBand18 = 65.640000
const FiberReferenceBand19 = 65.670000
const FiberReferenceBand20 = 65.700000
const FiberReferenceBand21 = 65.730000
const FiberReferenceBand22 = 65.760000
const FiberReferenceBand23 = 65.790000
const FiberReferenceBand24 = 65.820000
const FiberReferenceBand25 = 65.850000
const FiberReferenceBand26 = 65.880000
const FiberReferenceBand27 = 65.910000
const FiberReferenceBand28 = 65.940000
const FiberReferenceBand29 = 65.970000
const FiberReferenceBand30 = 66.000000
const FiberReferenceBand31 = 66.030000
const FiberReferenceBand32 = 66.060000
const FiberReferenceBand33 = 66.090000
const FiberReferenceBand34 = 66.120000
const FiberReferenceBand35 = 66.150000
const FiberReferenceBand36 = 66.180000
const FiberReferenceBand37 = 66.210000
const FiberReferenceBand38 = 66.240000
const FiberReferenceBand39 = 66.270000
const FiberReferenceBand40 = 66.300000
const FiberReferenceBand41 = 66.330000
const FiberReferenceBand42 = 66.360000
const FiberReferenceBand43 = 66.390000
const FiberReferenceBand44 = 66.420000
const FiberReferenceBand45 = 66.450000
const FiberReferenceBand46 = 66.480000
const FiberReferenceBand47 = 66.510000
const FiberReferenceBand48 = 66.540000
const FiberReferenceBand49 = 66.570000
const FiberReferenceBand50 = 66.600000
const FiberReferenceBand51 = 66.630000
const FiberReferenceBand52 = 66.660000
const FiberReferenceBand53 = 66.690000
const FiberReferenceBand54 = 66.720000
const FiberReferenceBand55 = 66.750000
const FiberReferenceBand56 = 66.780000
const FiberReferenceBand57 = 66.810000
const FiberReferenceBand58 = 66.840000
const FiberReferenceBand59 = 66.870000
const FiberReferenceBand60 = 66.900000
