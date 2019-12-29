package drivers

import (
	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"time"
)

type CarDriver struct {
	name   string
	wheels wheels
	gobot.Commander
}

//0:right;1:left
type wheels [2]*WheelDriver

var t *time.Ticker

func NewCarDriver(right, left *WheelDriver) *CarDriver {
	driver := &CarDriver{
		name:      gobot.DefaultName("car"),
		wheels:    wheels{right, left},
		Commander: gobot.NewCommander(),
	}
	driver.Commander.AddCommand("Front", func(m map[string]interface{}) interface{} {
		return driver.Front()
	})
	driver.Commander.AddCommand("Back", func(m map[string]interface{}) interface{} {
		return driver.Back()
	})
	driver.Commander.AddCommand("Left", func(m map[string]interface{}) interface{} {
		return driver.Left()
	})
	driver.Commander.AddCommand("Right", func(m map[string]interface{}) interface{} {
		return driver.Right()
	})

	return driver

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
	st()
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
	st()
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
	st()
	log.Info("driver/CarDriver: back")
	if err := c.wheels[0].Back(); err != nil {
		return err
	}
	if err := c.wheels[1].Back(); err != nil {
		return err
	}
	return nil
}

func st() {
	if t != nil {
		t.Stop()
	}
}

func (c *CarDriver) every(f func() error) {
	t = time.NewTicker(10 * time.Microsecond)
	go func() {
		for {
			select {
			case <-t.C:
				_ = f()
				_ = c.Stop()
			}
		}
	}()
}

func (c *CarDriver) Left() error {
	st()
	log.Info("driver/CarDriver: left")
	c.every(func() error {
		if err := c.wheels[0].Front(); err != nil {
			return err
		}
		if err := c.wheels[1].Back(); err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (c *CarDriver) Right() error {
	st()
	log.Info("driver/CarDriver: right")
	c.every(func() error {
		if err := c.wheels[0].Back(); err != nil {
			return err
		}
		if err := c.wheels[1].Front(); err != nil {
			return err
		}
		return nil
	})
	return nil
}
