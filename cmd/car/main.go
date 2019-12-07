package main

import (
	"eddy.org/pi/drivers"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func main() {
	a := raspi.NewAdaptor()
	l := drivers.NewWheelDriver(a, "11", "12")
	l.SetName("left wheel")
	r := drivers.NewWheelDriver(a, "15", "13")
	r.SetName("right wheel")
	c := drivers.NewCarDriver(r, l)

	work := func() {

		_ = c.Front()
		gobot.Every(100*time.Millisecond, func() {
			switch time.Now().Unix() % 4 {
			case 0:
				_ = c.Front()
			case 1:
				_ = c.Start()
			case 2:
				_ = c.Left()
			case 3:
				_ = c.Right()
			}
		})
	}

	robot := gobot.NewRobot("car",
		[]gobot.Connection{a},
		[]gobot.Device{c},
		work,
	)

	_ = robot.Start()
}
