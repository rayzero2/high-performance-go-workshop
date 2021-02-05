package main

import (
	"fmt"
	"unsafe"
)

// tag::struct[]
type S struct {
	a bool
	b float64
	c int32
}

// end::struct[]

// tag::padding[]
type S2 struct {
	a bool
	_ [7]byte // padding <1>
	b float64
	c int32
	_ [4]byte // padding <2>
}

// end::padding[]

func main() {
	// tag::sizeof[]
	var s S
	var s2 S
	fmt.Println(unsafe.Sizeof(s))  // <1>
	fmt.Println(unsafe.Sizeof(s2)) // <2>
	// end::sizeof[]
}
