package traffic

import "time"

// These are levels
// peak,off-peak,scale-up,scale-down
const (
	PeakLevel         = "peak"
	OffPeakLevel      = "off-peak"
	ScaleUpLevel      = "scale-up"
	ScaleDownLevel    = "scale-down"
	PeakDuration      = time.Minute * 1
	OffPeakDuration   = time.Minute * 5
	ScaleUpDuration   = time.Minute * 2
	ScaleDownDuration = time.Minute * 2
)

type Measure struct {
	Observation time.Duration
	Interval    time.Duration
}

type Window struct {
	Hour int
	From string // Peak,Off-Peak,Scale-Up,Scale-Down
	To   string
}

func NewWindow(hour, rate int, tag string) *Window {
	w := new(Window)
	//w.Rate = rate
	w.Hour = hour
	//w.Tag = tag
	return w
}

type Attention interface {
	Peak() Measure
	OffPeak() Measure
	ScaleUp() Measure
	ScaleDown() Measure
}

type Profile struct {
	windows   [24]Window
	attention map[string]Measure
}

func (p *Profile) Attention(ts time.Time) (from, to Measure) {
	w := p.windows[ts.Hour()]
	return p.attention[w.From], p.attention[w.To]
}

func (p *Profile) Level(ts time.Time) (from, to string) {
	w := p.windows[ts.Hour()]
	return w.From, w.To
}

func (p *Profile) IsPeak(ts time.Time) bool {
	w := p.windows[ts.Hour()]
	return w.From == PeakLevel
}

func (p *Profile) IsOffPeak(ts time.Time) bool {
	w := p.windows[ts.Hour()]
	return w.From == OffPeakLevel
}

func (p *Profile) IsScaleUp(ts time.Time) bool {
	w := p.windows[ts.Hour()]
	return w.From == ScaleUpLevel
}

func (p *Profile) IsScaleDown(ts time.Time) bool {
	w := p.windows[ts.Hour()]
	return w.From == ScaleDownLevel
}
