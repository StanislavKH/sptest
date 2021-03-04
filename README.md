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
goos: linux
goarch: amd64
BenchmarkFast-4                1        2014476139 ns/op
Download: 7.17 Mbps, Upload: 0.00 Mbps
BenchmarkSpeedtest-4           1        41369347497 ns/op
Download: 292.06 Mbps, Upload: 16.00 Mbps
```

