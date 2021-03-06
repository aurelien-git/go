/*

Exercise: Stringers

Make the 'IPAddr' type implement 'fmt.Stringer' to print the address as
a dotted quad.

For instance, IPAddr(1, 2, 3, 4) should print as "1.2.3.4"

*/

package main

import "fmt"

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

func main() {
	addrs := map[string]IPAddr{
		"lookback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
