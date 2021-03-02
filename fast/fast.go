package fast

import (
	"time"
	"errors"
	"github.com/npotts/go-fast"
)

const (
	DefaultWorkersCount     = 3
	DefaultChankSize        = 50000
)

// SpeedTetst -
type SpeedTest struct{}

// Result in bytes, direct int values
type Result struct {
	Upload   float64
	Download float64
}

// New -
func New() *SpeedTest {
	return &SpeedTest{}
}

// StartBenchmark -
func (st *SpeedTest) StartBenchmark() (*Result, error) {
	gf := gofast.New()
        cfg := gofast.Settings{
		MaxBytes: int64(DefaultChankSize),
		Timeout: time.Duration(5) * time.Second,
		Workers: DefaultWorkersCount,
		Network: time.Duration(1) * time.Second,
	}
	results := <-gf.Measure(cfg)
	if results.Workers == 0 {
		return nil, errors.New("unexpected")
	}
	return &Result{Upload: 0, Download: results.Mbps}, nil
}
