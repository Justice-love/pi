package main

import (
	"eddy.org/pi/drivers"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main() {
	r := raspi.NewAdaptor()
	w := drivers.NewWheelDriver(r, "11", "12")
	work := func() {
		gobot.Every(5*time.Second, func() {
			_ = w.Toggle()
		})
	}

	robot := gobot.NewRobot("wheelBot",
		[]gobot.Connection{r},
		[]gobot.Device{w},
		work,
	)

	_ = robot.Start()
}
