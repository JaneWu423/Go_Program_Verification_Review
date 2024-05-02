// type = ALL
package main

import "os"

func main() {
	CP
	applyOnce(func() {
		CS
	})
}

func applyOnce(f func()) {
	if len(os.Args) > 1 {
		f()
	}
}
