package examplepkg

import "strings"

// New creates a new instance of MyType and sets the privateVal
func New(val string) *ExampleType {
	return &ExampleType{
		PublicVal:  val,
		privateVal: strings.ToLower(val),
	}
}

// MyFunc returns the privateVal up until the specified index
func (et *ExampleType) MyFunc(n int) string {
	if n > len(et.privateVal) {
		n = len(et.privateVal)
	}
	return et.privateVal[:n]
}
