package ingress

import "time"

type facts struct {
	observed     metric
	threshold    metric
	rateOfChange int
	rps          int
	duration     time.Duration // Duration of the observation
	rateLimiting int           // Percentage of traffic that is currently being rate limited
}

type observed struct {
	saturation   int // Percentage of observed metric/threshold metric
	rateOfChange int // Rate of change of metric
	rateLimited  int // Percentage of traffic currently being limited
}

// Need to know the rate of change of the profile window for a given time.
// Need to know how far ahead into the future to look for saturation reaching 90-100%
// Rate of change is 0, and traffic is just starting to scale up.
// TODO : how to determine
func rate0_OffPeak_ScaleUp(saturation int, limited int, f facts, profile Profile) (int, bool) {
	//saturation := f.observed / f.threshold

	// TODO: determine the duration when the observed will reach the threshold. This should be based
	// on current saturation, profile, and rate of change.
	return 0, false
}

func findChange(o observed) (int, bool) {
	return 0, false
}
