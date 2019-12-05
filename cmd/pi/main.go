package main

import (
	"eddy.org/pi/drivers"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main() {
	distanceChan := make(chan int64, 100)
	a := raspi.NewAdaptor()
	r := drivers.NewWheelDriver(a, "11", "12")
	l := drivers.NewWheelDriver(a, "11", "12")
	c := drivers.NewCarDriver(r, l)
	u := drivers.NewUltrasonicSensorDriver(a, "13", "14", func(distance int64) {
		distanceChan <- distance
	})
	ac := drivers.NewAvoidanceCarDriver(c, u)

	_ = c.Front()
	go ac.Avoidance(distanceChan)

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			_ = u.Trig()
		})
	}

	robot := gobot.NewRobot("wheelBot",
		[]gobot.Connection{a},
		[]gobot.Device{ac},
		work,
	)

	_ = robot.Start()
}
