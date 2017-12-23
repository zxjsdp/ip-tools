package main

import (
	"fmt"

	"github.com/zxjsdp/ip-tools/ip"
)

func main() {
	input := "['41.188.12.16', '41.188.12.29', '58.28.63.15', '58.28.64.19','58.28.64.23', '58.28.64.29', '58.28.64.30', '58.28.64.34', '66.199.151.143']"

	input = ip.PrepareInputString(input)

	ips := ip.GetSingleIPsByRegexp(input)
	fmt.Println(ips)

	result := ip.GetRange(ips)
	fmt.Println(result)
}
