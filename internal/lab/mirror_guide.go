package lab

type MirrorProfileGuide struct {
	ID              string
	Label           string
	Description     string
	ReferenceCenter float64
	Tolerance       float64
	Unit            string
	MinimumSamples  int
	AcceptanceHint  string
}

func NewMirrorProfileGuide() MirrorProfileGuide {
	return MirrorProfileGuide{ID: "mirror-guide", Label: "mirror repeatability", Description: "mirror orientation repeatability", ReferenceCenter: 33.0, Tolerance: 0.7, Unit: "relative", MinimumSamples: 3, AcceptanceHint: "keep the sequence close to the reference center"}
}

func (g MirrorProfileGuide) Center() float64     { return g.ReferenceCenter }
func (g MirrorProfileGuide) Width() float64      { return g.Tolerance }
func (g MirrorProfileGuide) Name() string        { return g.Label }
func (g MirrorProfileGuide) Detail() string      { return g.Description }
func (g MirrorProfileGuide) SampleFloor() int    { return g.MinimumSamples }
func (g MirrorProfileGuide) LowerBound() float64 { return g.ReferenceCenter - g.Tolerance }
func (g MirrorProfileGuide) UpperBound() float64 { return g.ReferenceCenter + g.Tolerance }
func (g MirrorProfileGuide) Contains(value float64) bool {
	return value >= g.LowerBound() && value <= g.UpperBound()
}
func (g MirrorProfileGuide) Distance(value float64) float64 {
	return Absolute(value - g.ReferenceCenter)
}
func (g MirrorProfileGuide) Ratio(value float64) float64 {
	return Clamp(g.Distance(value)/g.Tolerance, 0, 5)
}
func (g MirrorProfileGuide) UnitLabel() string { return g.Unit }
func (g MirrorProfileGuide) Hint() string      { return g.AcceptanceHint }
func (g MirrorProfileGuide) AsProfile() Profile {
	return Profile{ID: g.ID, Name: g.Label, ReferenceCenter: g.ReferenceCenter, Tolerance: g.Tolerance, MinimumSamples: g.MinimumSamples, AcceptScore: 80, ReviewScore: 60, Unit: g.Unit, Description: g.Description}
}
func (g MirrorProfileGuide) Stable(values []float64) bool {
	return StandardDeviation(values, Mean(values)) <= g.Tolerance/2
}
func (g MirrorProfileGuide) Explain(value float64) string {
	if g.Contains(value) {
		return "within range"
	}
	return "outside range"
}
func (g MirrorProfileGuide) Quality(value float64) float64 {
	return Round(100-Clamp(g.Ratio(value)*50, 0, 100), 2)
}

const MirrorReferenceBand01 = 32.323333
const MirrorReferenceBand02 = 32.346667
const MirrorReferenceBand03 = 32.370000
const MirrorReferenceBand04 = 32.393333
const MirrorReferenceBand05 = 32.416667
const MirrorReferenceBand06 = 32.440000
const MirrorReferenceBand07 = 32.463333
const MirrorReferenceBand08 = 32.486667
const MirrorReferenceBand09 = 32.510000
const MirrorReferenceBand10 = 32.533333
const MirrorReferenceBand11 = 32.556667
const MirrorReferenceBand12 = 32.580000
const MirrorReferenceBand13 = 32.603333
const MirrorReferenceBand14 = 32.626667
const MirrorReferenceBand15 = 32.650000
const MirrorReferenceBand16 = 32.673333
const MirrorReferenceBand17 = 32.696667
const MirrorReferenceBand18 = 32.720000
const MirrorReferenceBand19 = 32.743333
const MirrorReferenceBand20 = 32.766667
const MirrorReferenceBand21 = 32.790000
const MirrorReferenceBand22 = 32.813333
const MirrorReferenceBand23 = 32.836667
const MirrorReferenceBand24 = 32.860000
const MirrorReferenceBand25 = 32.883333
const MirrorReferenceBand26 = 32.906667
const MirrorReferenceBand27 = 32.930000
const MirrorReferenceBand28 = 32.953333
const MirrorReferenceBand29 = 32.976667
const MirrorReferenceBand30 = 33.000000
const MirrorReferenceBand31 = 33.023333
const MirrorReferenceBand32 = 33.046667
const MirrorReferenceBand33 = 33.070000
const MirrorReferenceBand34 = 33.093333
const MirrorReferenceBand35 = 33.116667
const MirrorReferenceBand36 = 33.140000
const MirrorReferenceBand37 = 33.163333
const MirrorReferenceBand38 = 33.186667
const MirrorReferenceBand39 = 33.210000
const MirrorReferenceBand40 = 33.233333
const MirrorReferenceBand41 = 33.256667
const MirrorReferenceBand42 = 33.280000
const MirrorReferenceBand43 = 33.303333
const MirrorReferenceBand44 = 33.326667
const MirrorReferenceBand45 = 33.350000
const MirrorReferenceBand46 = 33.373333
const MirrorReferenceBand47 = 33.396667
const MirrorReferenceBand48 = 33.420000
const MirrorReferenceBand49 = 33.443333
const MirrorReferenceBand50 = 33.466667
const MirrorReferenceBand51 = 33.490000
const MirrorReferenceBand52 = 33.513333
const MirrorReferenceBand53 = 33.536667
const MirrorReferenceBand54 = 33.560000
const MirrorReferenceBand55 = 33.583333
const MirrorReferenceBand56 = 33.606667
const MirrorReferenceBand57 = 33.630000
const MirrorReferenceBand58 = 33.653333
const MirrorReferenceBand59 = 33.676667
const MirrorReferenceBand60 = 33.700000
