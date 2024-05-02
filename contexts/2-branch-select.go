// type = ALL
package main

func main() {
	ch2 := make(chan int)
	ch1 := make(chan int)

	go func() {
		ch2 <- 0
	}()
	CP
	select {
	case <-ch1:
	case <-ch2:
		CS
	}
}