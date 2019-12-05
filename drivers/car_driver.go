package drivers

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

type CarDriver struct {
	name       string
	wheels     wheels
	connection gpio.DigitalWriter
	gobot.Commander
}

//0:right;1:left
type wheels [2]*WheelDriver

func NewCarDriver(a gpio.DigitalWriter) *CarDriver {
	return &CarDriver{
		name:       gobot.DefaultName("car"),
		connection: a,
		Commander:  gobot.NewCommander(),
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
}

func (c *CarDriver) Halt() error {
	return nil
}

func (c *CarDriver) Connection() gobot.Connection {
	return c.connection.(gobot.Connection)
}

func (c *CarDriver) Stop() error {
	if err := c.wheels[0].Stop(); err != nil {
		return err
	}
	if err := c.wheels[1].Stop(); err != nil {
		return err
	}
	return nil
}

func (c *CarDriver) Front() error {
	if err := c.wheels[0].Front(); err != nil {
		return err
	}
	if err := c.wheels[1].Front(); err != nil {
		return err
	}
	return nil
}

func (c *CarDriver) Back() error {
	if err := c.wheels[0].Back(); err != nil {
		return err
	}
	if err := c.wheels[1].Back(); err != nil {
		return err
	}
	return nil
}

func (c *CarDriver) Left() error {
	if err := c.wheels[0].Front(); err != nil {
		return err
	}
	if err := c.wheels[1].Back(); err != nil {
		return err
	}
	return nil
}

func (c *CarDriver) Right() error {
	if err := c.wheels[0].Back(); err != nil {
		return err
	}
	if err := c.wheels[1].Front(); err != nil {
		return err
	}
	return nil
}
