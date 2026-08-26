package lab

type StageProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewStageProfileGuide() StageProfileGuide {
	return StageProfileGuide{ID: "stage-guide", Label: "stage return", Description: "translation stage return", ReferenceCenter: 5.0, Tolerance: 0.25, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g StageProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g StageProfileGuide) Width() float64      { return g.Tolerance }
func (g StageProfileGuide) Name() string        { return g.Label }
func (g StageProfileGuide) Detail() string      { return g.Description }
func (g StageProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g StageProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g StageProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g StageProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g StageProfileGuide) Distance(value float64) float64 {
	return Absolute(value - g.ReferenceCenter)
}
func (g StageProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g StageProfileGuide) UnitLabel() string { return g.Unit }
func (g StageProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g StageProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g StageProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g StageProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g StageProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const StageReferenceBand01 = 4.758333
const StageReferenceBand02 = 4.766667
const StageReferenceBand03 = 4.775000
const StageReferenceBand04 = 4.783333
const StageReferenceBand05 = 4.791667
const StageReferenceBand06 = 4.800000
const StageReferenceBand07 = 4.808333
const StageReferenceBand08 = 4.816667
const StageReferenceBand09 = 4.825000
const StageReferenceBand10 = 4.833333
const StageReferenceBand11 = 4.841667
const StageReferenceBand12 = 4.850000
const StageReferenceBand13 = 4.858333
const StageReferenceBand14 = 4.866667
const StageReferenceBand15 = 4.875000
const StageReferenceBand16 = 4.883333
const StageReferenceBand17 = 4.891667
const StageReferenceBand18 = 4.900000
const StageReferenceBand19 = 4.908333
const StageReferenceBand20 = 4.916667
const StageReferenceBand21 = 4.925000
const StageReferenceBand22 = 4.933333
const StageReferenceBand23 = 4.941667
const StageReferenceBand24 = 4.950000
const StageReferenceBand25 = 4.958333
const StageReferenceBand26 = 4.966667
const StageReferenceBand27 = 4.975000
const StageReferenceBand28 = 4.983333
const StageReferenceBand29 = 4.991667
const StageReferenceBand30 = 5.000000
const StageReferenceBand31 = 5.008333
const StageReferenceBand32 = 5.016667
const StageReferenceBand33 = 5.025000
const StageReferenceBand34 = 5.033333
const StageReferenceBand35 = 5.041667
const StageReferenceBand36 = 5.050000
const StageReferenceBand37 = 5.058333
const StageReferenceBand38 = 5.066667
const StageReferenceBand39 = 5.075000
const StageReferenceBand40 = 5.083333
const StageReferenceBand41 = 5.091667
const StageReferenceBand42 = 5.100000
const StageReferenceBand43 = 5.108333
const StageReferenceBand44 = 5.116667
const StageReferenceBand45 = 5.125000
const StageReferenceBand46 = 5.133333
const StageReferenceBand47 = 5.141667
const StageReferenceBand48 = 5.150000
const StageReferenceBand49 = 5.158333
const StageReferenceBand50 = 5.166667
const StageReferenceBand51 = 5.175000
const StageReferenceBand52 = 5.183333
const StageReferenceBand53 = 5.191667
const StageReferenceBand54 = 5.200000
const StageReferenceBand55 = 5.208333
const StageReferenceBand56 = 5.216667
const StageReferenceBand57 = 5.225000
const StageReferenceBand58 = 5.233333
const StageReferenceBand59 = 5.241667
const StageReferenceBand60 = 5.250000
