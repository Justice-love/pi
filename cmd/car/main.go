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
	r := drivers.NewWheelDriver(a, "12", "11")
	l := drivers.NewWheelDriver(a, "13", "15")
	r.SetName("right wheel")
	l.SetName("left wheel")
	c := drivers.NewCarDriver(r, l)

	work := func() {

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
