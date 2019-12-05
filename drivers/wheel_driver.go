package drivers

import (
	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"sync"
)

const (
	front = iota
	back
	stop
)

type wheelState int8

type WheelDriver struct {
	pinRight   string
	pinLeft    string
	name       string
	state      wheelState
	lock       sync.Mutex
	connection gpio.DigitalWriter
	gobot.Commander
}

func NewWheelDriver(a gpio.DigitalWriter, pinRight, pinLeft string) *WheelDriver {
	return &WheelDriver{
		pinRight:   pinRight,
		pinLeft:    pinLeft,
		name:       gobot.DefaultName("Wheel"),
		state:      stop,
		connection: a,
		Commander:  gobot.NewCommander(),
	}
}

func (w *WheelDriver) Name() string {
	return w.name
}

func (w *WheelDriver) SetName(s string) {
	w.name = s
}

func (w *WheelDriver) Start() error {
	return nil
}

func (w *WheelDriver) Halt() error {
	return nil
}

func (w *WheelDriver) Connection() gobot.Connection {
	return w.connection.(gobot.Connection)
}

func (w *WheelDriver) Stop() error {
	log.WithField("wheel", w.name).Info("driver/WheelDriver: stop")
	if w.state == stop {
		return nil
	}
	w.lock.Lock()
	defer w.lock.Unlock()

	if err := w.connection.DigitalWrite(w.pinRight, 0); err != nil {
		return err
	}
	if err := w.connection.DigitalWrite(w.pinLeft, 0); err != nil {
		return err
	}
	w.state = stop

	return nil
}

func (w *WheelDriver) Front() error {
	log.WithField("wheel", w.name).Info("driver/WheelDriver: front")
	if w.state == front {
		return nil
	}
	w.lock.Lock()
	defer w.lock.Unlock()
	if err := w.connection.DigitalWrite(w.pinRight, 1); err != nil {
		return err
	}
	if err := w.connection.DigitalWrite(w.pinLeft, 0); err != nil {
		return err
	}
	w.state = front

	return nil
}

func (w *WheelDriver) Back() error {
	log.WithField("wheel", w.name).Info("driver/WheelDriver: back")
	if w.state == back {
		return nil
	}
	w.lock.Lock()
	defer w.lock.Unlock()
	if err := w.connection.DigitalWrite(w.pinRight, 0); err != nil {
		return err
	}
	if err := w.connection.DigitalWrite(w.pinLeft, 1); err != nil {
		return err
	}
	w.state = back

	return nil
}

func (w *WheelDriver) Toggle() error {
	log.WithField("wheel", w.name).Info("driver/WheelDriver: toggle")
	if w.state == stop || w.state == back {
		if err := w.Front(); err != nil {
			return err
		}
		w.state = front
	} else {
		if err := w.Back(); err != nil {
			return err
		}
		w.state = back
	}
	return nil
}
