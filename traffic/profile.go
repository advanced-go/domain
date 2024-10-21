package traffic

import "time"

const (
	Peak              = "peak"
	OffPeak           = "off-peak"
	ScaleUp           = "scale-up"
	ScaleDown         = "scale-down"
	PeakDuration      = time.Minute * 1
	OffPeakDuration   = time.Minute * 5
	ScaleUpDuration   = time.Minute * 2
	ScaleDownDuration = time.Minute * 2
)

type Attention struct {
	Observation time.Duration
	Interval    time.Duration
}

type Window struct {
	Hour int
	Tag  string // Peak,Off-Peak,Scale-Up,Scale-Down
	Rate int
}

func NewWindow(hour, rate int, tag string) *Window {
	w := new(Window)
	w.Rate = rate
	w.Hour = hour
	w.Tag = tag
	return w
}

type Profile struct {
}
