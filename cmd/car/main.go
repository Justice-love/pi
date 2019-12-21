package main

import (
	"eddy.org/pi/drivers"
	"eddy.org/pi/internal/car"
	"github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
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
		_ = car.Setup(func(direction car.Direction) {
			switch direction {
			case car.Direction_FRONT:
				_ = c.Front()
			case car.Direction_BACK:
				_ = c.Back()
			case car.Direction_LEFT:
				_ = c.Left()
			case car.Direction_RIGHT:
				_ = c.Right()
			default:
				logrus.WithField("command", direction.String()).Fatal("cmd/car: unhandled command")
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
