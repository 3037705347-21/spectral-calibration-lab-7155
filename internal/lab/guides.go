package lab

func AllProfileGuides() []Profile {
	guides := []Profile{
		NewThermalProfileGuide().AsProfile(),
		NewLampProfileGuide().AsProfile(),
		NewDetectorProfileGuide().AsProfile(),
		NewFilterProfileGuide().AsProfile(),
		NewShutterProfileGuide().AsProfile(),
		NewMirrorProfileGuide().AsProfile(),
		NewFiberProfileGuide().AsProfile(),
		NewStageProfileGuide().AsProfile(),
		NewDarkProfileGuide().AsProfile(),
		NewGainProfileGuide().AsProfile(),
		NewFocusProfileGuide().AsProfile(),
		NewApertureProfileGuide().AsProfile(),
		NewExposureProfileGuide().AsProfile(),
		NewScanProfileGuide().AsProfile(),
	}
	return guides
}
