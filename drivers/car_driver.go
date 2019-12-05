package drivers

import (
	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
)

type CarDriver struct {
	name   string
	wheels wheels
	gobot.Commander
}

//0:right;1:left
type wheels [2]*WheelDriver

func NewCarDriver(right, left *WheelDriver) *CarDriver {
	return &CarDriver{
		name:      gobot.DefaultName("car"),
		wheels:    wheels{right, left},
		Commander: gobot.NewCommander(),
	}
}

func (c *CarDriver) Name() string {
	return c.name
}

func (c *CarDriver) SetName(s string) {
	c.name = s
}

func (c *CarDriver) Start() error {
	for _, wheel := range c.wheels {
		_ = wheel.Start()
	}
	return nil
}

func (c *CarDriver) Halt() error {
	return nil
}

func (c *CarDriver) Connection() gobot.Connection {
	return nil
}

func (c *CarDriver) Stop() error {
	log.Info("driver/CarDriver: stop")
	if err := c.wheels[0].Stop(); err != nil {
		return err
	}
	if err := c.wheels[1].Stop(); err != nil {
		return err
	}
	return nil
}

func (c *CarDriver) Front() error {
	log.Info("driver/CarDriver: front")
	if err := c.wheels[0].Front(); err != nil {
		return err
	}
	if err := c.wheels[1].Front(); err != nil {
		return err
	}
	return nil
}

func (c *CarDriver) Back() error {
	log.Info("driver/CarDriver: back")
	if err := c.wheels[0].Back(); err != nil {
		return err
	}
	if err := c.wheels[1].Back(); err != nil {
		return err
	}
	return nil
}

func (c *CarDriver) Left() error {
	log.Info("driver/CarDriver: left")
	if err := c.wheels[0].Front(); err != nil {
		return err
	}
	if err := c.wheels[1].Back(); err != nil {
		return err
	}
	return nil
}

func (c *CarDriver) Right() error {
	log.Info("driver/CarDriver: right")
	if err := c.wheels[0].Back(); err != nil {
		return err
	}
	if err := c.wheels[1].Front(); err != nil {
		return err
	}
	return nil
}
