package ingress

import "time"

// Algorithm architecture
// Given an observation:
type metric int
type slope int

type Profile interface {
	Peak() bool
	OffPeak() bool
	ScaleUp() bool
	ScaleDown() bool
}

// T could be any metric: latency, range of status codes, percentile...
// T probably needs to be an interface
type observation[T any] struct {
	duration  time.Duration // Duration of the observation
	value     T
	threshold T
	// Current load
	rps int

	// TODO: how does rate of change affect the magnitude of an action.
	rateOfChange int // How fast is the value changing, higher rate of change is faster
}

// What happens if the observation changes within the duration?
type rateLimitingObservation struct {
	duration time.Duration // Duration of the observation

	// Current constraints
	limit int
	burst int

	// Percentage of traffic that was constrained/rate limited
	pctOfTraffic int
}

func (r *rateLimitingObservation) IsLimiting() bool {
	return r.limit > 0
}

// What happens if the observation changes within the duration?
type redirectObservation struct {
	duration time.Duration // Duration of the observation

	// Current counts
	primary   int
	secondary int

	// Percentage of traffic that was constrained/rate limited
	pctOfTraffic int
}

type frameContext struct {
	// TODO: how does current and future load affect the magnitude of an action
	profile Profile // What is the current and future load on the system

	// TODO: how does experience affect the magnitude of an action
	experience any

	// TODO: how does guidance affect the magnitude of an action
	guidance any
}

func shouldAct(observed, threshold metric, rateOfChange slope) bool {

	if observed < threshold {
		return false
	}

	return true
}

func action[T any](data observation[T], rateLimit rateLimitingObservation) int {
	// if no change in observation, stay with current action which could be rate limited or not
	if data.rateOfChange == 0 {
		return 0
	}
	// if the observed metric is decreasing, then reduce the rate limiting if applicable
	if data.rateOfChange < 0 {
		if rateLimit.IsLimiting() {
			return -1
		}
		return 0
	}

	return 0
}
