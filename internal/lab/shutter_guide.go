package lab

type ShutterProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewShutterProfileGuide() ShutterProfileGuide {
	return ShutterProfileGuide{ID: "shutter-guide", Label: "shutter timing", Description: "shutter travel timing", ReferenceCenter: 14.0, Tolerance: 0.6, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g ShutterProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g ShutterProfileGuide) Width() float64      { return g.Tolerance }
func (g ShutterProfileGuide) Name() string        { return g.Label }
func (g ShutterProfileGuide) Detail() string      { return g.Description }
func (g ShutterProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g ShutterProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g ShutterProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g ShutterProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g ShutterProfileGuide) Distance(value float64) float64 {
	return Absolute(value - g.ReferenceCenter)
}
func (g ShutterProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g ShutterProfileGuide) UnitLabel() string { return g.Unit }
func (g ShutterProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g ShutterProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g ShutterProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g ShutterProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g ShutterProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const ShutterReferenceBand01 = 13.420000
const ShutterReferenceBand02 = 13.440000
const ShutterReferenceBand03 = 13.460000
const ShutterReferenceBand04 = 13.480000
const ShutterReferenceBand05 = 13.500000
const ShutterReferenceBand06 = 13.520000
const ShutterReferenceBand07 = 13.540000
const ShutterReferenceBand08 = 13.560000
const ShutterReferenceBand09 = 13.580000
const ShutterReferenceBand10 = 13.600000
const ShutterReferenceBand11 = 13.620000
const ShutterReferenceBand12 = 13.640000
const ShutterReferenceBand13 = 13.660000
const ShutterReferenceBand14 = 13.680000
const ShutterReferenceBand15 = 13.700000
const ShutterReferenceBand16 = 13.720000
const ShutterReferenceBand17 = 13.740000
const ShutterReferenceBand18 = 13.760000
const ShutterReferenceBand19 = 13.780000
const ShutterReferenceBand20 = 13.800000
const ShutterReferenceBand21 = 13.820000
const ShutterReferenceBand22 = 13.840000
const ShutterReferenceBand23 = 13.860000
const ShutterReferenceBand24 = 13.880000
const ShutterReferenceBand25 = 13.900000
const ShutterReferenceBand26 = 13.920000
const ShutterReferenceBand27 = 13.940000
const ShutterReferenceBand28 = 13.960000
const ShutterReferenceBand29 = 13.980000
const ShutterReferenceBand30 = 14.000000
const ShutterReferenceBand31 = 14.020000
const ShutterReferenceBand32 = 14.040000
const ShutterReferenceBand33 = 14.060000
const ShutterReferenceBand34 = 14.080000
const ShutterReferenceBand35 = 14.100000
const ShutterReferenceBand36 = 14.120000
const ShutterReferenceBand37 = 14.140000
const ShutterReferenceBand38 = 14.160000
const ShutterReferenceBand39 = 14.180000
const ShutterReferenceBand40 = 14.200000
const ShutterReferenceBand41 = 14.220000
const ShutterReferenceBand42 = 14.240000
const ShutterReferenceBand43 = 14.260000
const ShutterReferenceBand44 = 14.280000
const ShutterReferenceBand45 = 14.300000
const ShutterReferenceBand46 = 14.320000
const ShutterReferenceBand47 = 14.340000
const ShutterReferenceBand48 = 14.360000
const ShutterReferenceBand49 = 14.380000
const ShutterReferenceBand50 = 14.400000
const ShutterReferenceBand51 = 14.420000
const ShutterReferenceBand52 = 14.440000
const ShutterReferenceBand53 = 14.460000
const ShutterReferenceBand54 = 14.480000
const ShutterReferenceBand55 = 14.500000
const ShutterReferenceBand56 = 14.520000
const ShutterReferenceBand57 = 14.540000
const ShutterReferenceBand58 = 14.560000
const ShutterReferenceBand59 = 14.580000
const ShutterReferenceBand60 = 14.600000
