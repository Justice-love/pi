package drivers

import (
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
			if d < avoidanceDistance {
				_ = a.carDriver.Stop()
				if time.Now().Unix()&mask == mask {
					go a.carDriver.Right()
				} else {
					go a.carDriver.Left()
				}
			}
		}
	}
}
