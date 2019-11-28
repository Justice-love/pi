package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main() {
	r := raspi.NewAdaptor()

	led1 := gpio.NewLedDriver(r, "11")
	led2 := gpio.NewLedDriver(r, "12")

	led1.On()
	led2.Off()
	work := func() {
		gobot.Every(5*time.Second, func() {

			led1.Toggle()
			led2.Toggle()
		})
	}

	robot := gobot.NewRobot("motorBot",
		[]gobot.Connection{r},
		[]gobot.Device{led1, led2},
		work,
	)

	_ = robot.Start()
}
