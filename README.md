##### Netflix as I know don't provide upload test, only download via Fast
```
        sp := sptest.New()
        res, err := sp.GetNetworkSpeed(&sptest.Options{CheckType: sptest.TypeFast})
        if err != nil {
                fmt.Println(err)
        }
        fmt.Printf("Download: %.2f Mbps, Upload: %.2f Mbps \n", res.Download, res.Upload)
```
##### Download: 9.22 Mbps, Upload: 0.00 Mbps


##### Speedtest.net up/down rate
```
        sp := sptest.New()
        res, err := sp.GetNetworkSpeed(&sptest.Options{CheckType: sptest.TypeSpeed})
        if err != nil {
                fmt.Println(err)
        }
        fmt.Printf("Download: %.2f Mbps, Upload: %.2f Mbps \n", res.Download, res.Upload)
```
##### Download: 794.96 Mbps, Upload: 16.00 Mbps

###### Benchmark
###### Speedtest much complex and additionaly work with upload test
```
go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/StanislavKH/sptest
BenchmarkFast-4                1        1814710986 ns/op         8065808 B/op     162626 allocs/op
BenchmarkSpeedtest-4           1        41377012898 ns/op        4289560 B/op      17733 allocs/op
PASS
ok      github.com/StanislavKH/sptest   43.206s
```

