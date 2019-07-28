package main

import (
	"fmt"
	"github.com/shirdonl/TP-Link-HS110"
)

func main() {
	plug := tplinkhs110.Hs110Plug{IPAddress: "192.168.0.196"}
	results, err := plug.MeterInfo()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(results)

}
