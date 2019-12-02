package drivers

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"time"
)

type UltrasonicSensorDriver struct {
	pinTrig    string
	pinEcho    string
	name       string
	trigChan   chan int8
	echoChan   chan float32
	connection gobot.Adaptor
	gobot.Commander
}

func (u *UltrasonicSensorDriver) Name() string {
	return u.name
}

func (u *UltrasonicSensorDriver) SetName(s string) {
	u.name = s
}

func (u *UltrasonicSensorDriver) Start() error {
	go func() {
		for {
			select {
			case <-u.trigChan:
				cr := checkHigh(u.connection.(gpio.DigitalReader), u.pinEcho)
				if !cr {
					continue
				}
				begin := time.Now()
				checkLow(u.connection.(gpio.DigitalReader), u.pinEcho)
				d := time.Since(begin)
				distance := float32(d.Milliseconds()) * 0.34 / 2.0
				u.echoChan <- distance
			}
		}
	}()
	return nil
}

func checkHigh(reader gpio.DigitalReader, pinEcho string) bool {
	now := time.Now()
	for {
		val, _ := reader.DigitalRead(pinEcho)
		if val > 0 {
			return true
		}
		if time.Since(now).Milliseconds() > 100 {
			return false
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func checkLow(reader gpio.DigitalReader, pinEcho string) {
	for {
		val, _ := reader.DigitalRead(pinEcho)
		if val <= 0 {
			return
		}
	}
}

func (u *UltrasonicSensorDriver) Halt() error {
	return nil
}

func (u *UltrasonicSensorDriver) Connection() gobot.Connection {
	return u.connection
}

func NewUltrasonicSensorDriver(adaptor gobot.Adaptor, pinTrig string, pinEcho string) *UltrasonicSensorDriver {
	return &UltrasonicSensorDriver{
		pinTrig:    pinTrig,
		pinEcho:    pinEcho,
		name:       gobot.DefaultName("Ultrasonic"),
		trigChan:   make(chan int8, 1),
		echoChan:   make(chan float32, 1),
		connection: adaptor,
		Commander:  gobot.NewCommander(),
	}
}

func (u *UltrasonicSensorDriver) Trig() error {
	if err := u.connection.(gpio.DigitalWriter).DigitalWrite(u.pinTrig, 1); err != nil {
		return err
	}
	c := time.After(10 * time.Microsecond)
	<-c
	if err := u.connection.(gpio.DigitalWriter).DigitalWrite(u.pinTrig, 0); err != nil {
		return err
	}
	u.trigChan <- 1
	return nil
}
