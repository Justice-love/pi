package main

import (
	"eddy.org/pi/drivers"
	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

func init() {
	log.SetLevel(log.WarnLevel)
}

func main() {
	distanceChan := make(chan int64, 100)
	a := raspi.NewAdaptor()
	l := drivers.NewWheelDriver(a, "11", "12")
	l.SetName("left wheel")
	r := drivers.NewWheelDriver(a, "15", "13")
	r.SetName("right wheel")
	c := drivers.NewCarDriver(r, l)
	u := drivers.NewUltrasonicSensorDriver(a, "40", "38", func(distance int64) {
		distanceChan <- distance
	})
	ac := drivers.NewAvoidanceCarDriver(c, u)

	work := func() {

		_ = c.Front()
		go ac.Avoidance(distanceChan)

		gobot.Every(100*time.Millisecond, func() {
			_ = u.Trig()
		})
	}

	robot := gobot.NewRobot("avoidance_car",
		[]gobot.Connection{a},
		[]gobot.Device{ac},
		work,
	)

	_ = robot.Start()
}
