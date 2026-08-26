package lab

type ThermalProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewThermalProfileGuide() ThermalProfileGuide {
	return ThermalProfileGuide{ID: "thermal-guide", Label: "thermal stability", Description: "temperature-conditioned optical alignment", ReferenceCenter: 10.0, Tolerance: 0.45, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g ThermalProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g ThermalProfileGuide) Width() float64      { return g.Tolerance }
func (g ThermalProfileGuide) Name() string        { return g.Label }
func (g ThermalProfileGuide) Detail() string      { return g.Description }
func (g ThermalProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g ThermalProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g ThermalProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g ThermalProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g ThermalProfileGuide) Distance(value float64) float64 {
	return Absolute(value - g.ReferenceCenter)
}
func (g ThermalProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g ThermalProfileGuide) UnitLabel() string { return g.Unit }
func (g ThermalProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g ThermalProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g ThermalProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g ThermalProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g ThermalProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const ThermalReferenceBand01 = 9.565000
const ThermalReferenceBand02 = 9.580000
const ThermalReferenceBand03 = 9.595000
const ThermalReferenceBand04 = 9.610000
const ThermalReferenceBand05 = 9.625000
const ThermalReferenceBand06 = 9.640000
const ThermalReferenceBand07 = 9.655000
const ThermalReferenceBand08 = 9.670000
const ThermalReferenceBand09 = 9.685000
const ThermalReferenceBand10 = 9.700000
const ThermalReferenceBand11 = 9.715000
const ThermalReferenceBand12 = 9.730000
const ThermalReferenceBand13 = 9.745000
const ThermalReferenceBand14 = 9.760000
const ThermalReferenceBand15 = 9.775000
const ThermalReferenceBand16 = 9.790000
const ThermalReferenceBand17 = 9.805000
const ThermalReferenceBand18 = 9.820000
const ThermalReferenceBand19 = 9.835000
const ThermalReferenceBand20 = 9.850000
const ThermalReferenceBand21 = 9.865000
const ThermalReferenceBand22 = 9.880000
const ThermalReferenceBand23 = 9.895000
const ThermalReferenceBand24 = 9.910000
const ThermalReferenceBand25 = 9.925000
const ThermalReferenceBand26 = 9.940000
const ThermalReferenceBand27 = 9.955000
const ThermalReferenceBand28 = 9.970000
const ThermalReferenceBand29 = 9.985000
const ThermalReferenceBand30 = 10.000000
const ThermalReferenceBand31 = 10.015000
const ThermalReferenceBand32 = 10.030000
const ThermalReferenceBand33 = 10.045000
const ThermalReferenceBand34 = 10.060000
const ThermalReferenceBand35 = 10.075000
const ThermalReferenceBand36 = 10.090000
const ThermalReferenceBand37 = 10.105000
const ThermalReferenceBand38 = 10.120000
const ThermalReferenceBand39 = 10.135000
const ThermalReferenceBand40 = 10.150000
const ThermalReferenceBand41 = 10.165000
const ThermalReferenceBand42 = 10.180000
const ThermalReferenceBand43 = 10.195000
const ThermalReferenceBand44 = 10.210000
const ThermalReferenceBand45 = 10.225000
const ThermalReferenceBand46 = 10.240000
const ThermalReferenceBand47 = 10.255000
const ThermalReferenceBand48 = 10.270000
const ThermalReferenceBand49 = 10.285000
const ThermalReferenceBand50 = 10.300000
const ThermalReferenceBand51 = 10.315000
const ThermalReferenceBand52 = 10.330000
const ThermalReferenceBand53 = 10.345000
const ThermalReferenceBand54 = 10.360000
const ThermalReferenceBand55 = 10.375000
const ThermalReferenceBand56 = 10.390000
const ThermalReferenceBand57 = 10.405000
const ThermalReferenceBand58 = 10.420000
const ThermalReferenceBand59 = 10.435000
const ThermalReferenceBand60 = 10.450000
