package mypkg

import "strings"

func New(val string) *MyType {
	return &MyType{
		PublicVal:  val,
		privateVal: strings.ToLower(val),
	}
}

func (myt *MyType) MyFunc(n int) string {
	if n > len(myt.privateVal) {
		n = len(myt.privateVal)
	}
	return myt.privateVal[:n]
}
