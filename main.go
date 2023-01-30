package main

import (
	"fmt"
	"time"
)

func main() {
	timer()
}

func timer() {
	timer := time.NewTicker(10 * time.Second)
	fmt.Println("Hello World")
	<-timer.C
	fmt.Println("Goodbye world")
}
