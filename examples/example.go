package main

import (
	"fmt"
	"github.com/shirdonl/tplinkhs110"
)

func main() {
	plug := tplinkhs110.Hs1xxPlug{IPAddress: "192.168.0.196"}
	results, err := plug.MeterInfo()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(results)

}
