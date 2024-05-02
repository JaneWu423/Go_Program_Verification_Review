// type = ALL
package main

import "time"

func main() {
	CP

	select {
	case <-time.After(3 * time.Second):
		CS
	}
}
