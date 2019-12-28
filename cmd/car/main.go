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
	r.SetName("left wheel")
	l := drivers.NewWheelDriver(a, "13", "15")
	l.SetName("right wheel")
	c := drivers.NewCarDriver(r, l)

	work := func() {

		_ = car.Setup(func(direction car.Direction) {
			var err error
			switch direction {
			case car.Direction_FRONT:
				err = c.Front()
			case car.Direction_BACK:
				err = c.Back()
			case car.Direction_LEFT:
				err = c.Left()
			case car.Direction_RIGHT:
				err = c.Right()
			case car.Direction_STOP:
				err = c.Stop()
			default:
				logrus.WithField("command", direction.String()).Fatal("cmd/car: unhandled command")
			}
			if err != nil {
				logrus.Error(err)
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
