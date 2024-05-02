// type = ALL
package main

import "os"

type I interface {
	f(CP_TYPE)
}

type A struct {
}

func (a A) f(CP_PARAM) {
	CS
}

type B struct {
}

func (b B) f(CP_PARAM) {

}

func main() {

	var v I

	if len(os.Args) > 0 {
		v = A{}
	} else {
		v = B{}
	}
	CP
	v.f(CP_ARG)
}
