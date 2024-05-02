// type = CH
package main

func main() {

	ch := make(chan int, 4)

	ch <- 0
	ch <- 0
	ch <- 0
	ch <- 0
	<-ch
	<-ch
	<-ch
	<-ch

	CS

}
