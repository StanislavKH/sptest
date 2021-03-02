package main

import(
	"fmt"
	sptest "github.com/StanislavKH/sptest"
)

func main() {
	sp := sptest.New()
	res, err := sp.GetNetworkSpeed(&sptest.Options{CheckType: sptest.TypeSpeed})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Download: %.2f Mbps, Upload: %.2f Mbps \n", res.Download, res.Upload)
}