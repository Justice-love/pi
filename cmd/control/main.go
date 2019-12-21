package main

import (
	"eddy.org/pi/internal/car"
	"os"
	"os/signal"
)

func main() {

	car.Setup()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
