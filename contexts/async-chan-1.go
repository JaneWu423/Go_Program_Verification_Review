// type = CH
package main

func main() {

	ch := make(chan int, 1)

	ch <- 0
	<-ch

	CS

}
