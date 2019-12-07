package drivers

import (
	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"time"
)

const (
	avoidanceDistance = 5
	mask              = 0x0000000000000001
)

type AvoidanceCarDriver struct {
	carDriver              *CarDriver
	ultrasonicSensorDriver *UltrasonicSensorDriver
	name                   string
	gobot.Commander
}

func NewAvoidanceCarDriver(cDriver *CarDriver, uDriver *UltrasonicSensorDriver) *AvoidanceCarDriver {
	return &AvoidanceCarDriver{
		carDriver:              cDriver,
		ultrasonicSensorDriver: uDriver,
		name:                   "avoidanceCar",
		Commander:              gobot.NewCommander(),
	}
}

func (a *AvoidanceCarDriver) Name() string {
	return a.name
}

func (a *AvoidanceCarDriver) SetName(s string) {
	a.name = s
}

func (a *AvoidanceCarDriver) Start() error {
	if err := a.carDriver.Start(); err != nil {
		return err
	}
	if err := a.ultrasonicSensorDriver.Start(); err != nil {
		return err
	}
	return nil
}

func (a *AvoidanceCarDriver) Halt() error {
	return nil
}

func (a *AvoidanceCarDriver) Connection() gobot.Connection {
	return nil
}

func (a *AvoidanceCarDriver) Avoidance(distanceChan chan int64) {
	for {
		select {
		case d := <-distanceChan:
			log.WithField("distance", d).Info("driver/AvoidanceCarDriver: receive distance")
			if d < avoidanceDistance {
				log.WithField("distance", d).Warn("driver/AvoidanceCarDriver: will change direction")
				_ = a.carDriver.Stop()
				if time.Now().Unix()&mask == mask {
					log.WithField("distance", d).Warn("driver/AvoidanceCarDriver: turn right")
					go func() {
						_ = a.carDriver.Back()
						time.Sleep(500 * time.Millisecond)
						_ = a.carDriver.Right()
						time.Sleep(2 * time.Second)
						_ = a.carDriver.Front()
					}()
				} else {
					log.WithField("distance", d).Warn("driver/AvoidanceCarDriver: turn left")
					go func() {
						_ = a.carDriver.Back()
						time.Sleep(500 * time.Millisecond)
						_ = a.carDriver.Left()
						time.Sleep(2 * time.Second)
						_ = a.carDriver.Front()
					}()
				}
			}
		}
	}
}
