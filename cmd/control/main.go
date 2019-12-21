package main

import (
	"eddy.org/pi/internal/car"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {

	car.Setup(func(direction car.Direction) {
		logrus.Info(direction.String())
	})
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
