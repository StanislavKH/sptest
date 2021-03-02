package speedtest

import (
	"github.com/johnsto/speedtest"
	"net"
	"net/http"
	"time"
)

// Default values
const (
	DefaultThreads    = 4
	DefaultMaxThreads = 16
)

// SpeedTest -
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
	settings, err := speedtest.FetchSettings()
	if err != nil {
		return nil, err
	}
	client := prepareClient()
	upRate := prepareRate(getUploadRate(client, settings))
	downRate := prepareRate(getDownloadRate(client, settings))
	return &Result{Upload: upRate, Download: downRate}, nil
}

func getDownloadRate(client http.Client, settings speedtest.Settings) int {
	benchmark := speedtest.NewDownloadBenchmark(client, settings.Servers[0])
	return speedtest.RunBenchmark(benchmark, DefaultThreads, DefaultMaxThreads, time.Second*10)
}

func getUploadRate(client http.Client, settings speedtest.Settings) int {
	benchmark := speedtest.NewUploadBenchmark(client, settings.Servers[0])
	return speedtest.RunBenchmark(benchmark, DefaultThreads, DefaultMaxThreads, time.Second*10)
}

func prepareClient() http.Client {
	client := http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(10*time.Second))
			},
		},
	}
	return client
}

func prepareRate(rate int) float64 {
	return float64(8 * rate) / 1024 / 1024
}