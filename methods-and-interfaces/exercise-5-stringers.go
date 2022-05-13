package tour

import (
	"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var s string
	for _, v := range ip {
		s += fmt.Sprintf("%v.", v)
	}
	return s[:len(s)-1]
}

func (ip IPAddr) String2() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
