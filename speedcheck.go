package speedcheck

import (
	"errors"
	"github.com/StanislavKH/sptest/fast"
	"github.com/StanislavKH/sptest/speedtest"
)

// Possible types
const (
	TypeFast  = "fast"
	TypeSpeed = "speedtest"
)

// Response -
type Response struct {
	Download float64
	Upload   float64
}

// Options -
type Options struct {
	CheckType string
}

// SpCheck -
type SpCheck struct {
	stest *speedtest.SpeedTest
	ftest *fast.SpeedTest
}

// New init available checkers
func New() *SpCheck {
	return &SpCheck{
		stest: speedtest.New(),
		ftest: fast.New(),
	}
}

// GetNetworkSpeed process checking for provided option which define type (Ookla or Netflix)
func (s *SpCheck) GetNetworkSpeed(op *Options) (*Response, error) {
	switch op.CheckType {
	case TypeSpeed:
		res, err := s.stest.StartBenchmark()
		return &Response{Download: res.Download, Upload: res.Upload}, err
	case TypeFast:
		fres, err := s.ftest.StartBenchmark()
		return &Response{Download: fres.Download, Upload: fres.Upload}, err
	}
	return nil, errors.New("speedtest check type undefined or unsupported")
}
