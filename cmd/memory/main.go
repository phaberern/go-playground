package main

import "strings"

//go:noinline
func output(x int) int {
	y := x + 5
	return y
}

//go:noinline
func output2(x int) *int {
	y := x + 5
	return &y
}

// the way most programmers concatenate strings
//
//go:noinline
func strconcat() string {
	x := ""
	for i := 0; i < 10000; i++ {
		x += "y"
	}
	return x
}

// a more efficient approach to concatenating strings
//
//go:noinline
func strconcatEfficient() string {
	strBldr := strings.Builder{}
	strBldr.WriteString("")
	for i := 0; i < 10000; i++ {
		strBldr.WriteString("-")
	}
	return strBldr.String()
}

func main() {
	_ = output(5)
	_ = output2(5)
}

// tools to inspect allocation behavior
// go build -gcflags '-m'
// go test -bench . -benchmem
