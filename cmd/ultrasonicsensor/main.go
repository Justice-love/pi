package main

import (
	"eddy.org/pi/drivers"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main() {
	r := raspi.NewAdaptor()
	u := drivers.NewUltrasonicSensorDriver(r, "12", "11")
	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			_ = u.Trig()

		})
	}

	robot := gobot.NewRobot("UBot",
		[]gobot.Connection{r},
		[]gobot.Device{u},
		work,
	)

	_ = robot.Start()
}
