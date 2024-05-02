// type = ALL
// bounds = 10
package main

func main() {
	CP
	rec(CP_ARG, bound)
}

func rec(CP_PARAM, i int) {

	if i > 0 {
		CS
		rec(CP_ARG_NO_PTR, i-1)
	}
}
