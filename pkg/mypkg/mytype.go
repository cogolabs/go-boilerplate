package mypkg

import "strings"

// New creates a new instance of MyType and sets the privateVal
func New(val string) *MyType {
	return &MyType{
		PublicVal:  val,
		privateVal: strings.ToLower(val),
	}
}

// MyFunc returns the privateVal up until the specified index
func (myt *MyType) MyFunc(n int) string {
	if n > len(myt.privateVal) {
		n = len(myt.privateVal)
	}
	return myt.privateVal[:n]
}
