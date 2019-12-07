package drivers

import (
	log "github.com/sirupsen/logrus"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"time"
)

type UltrasonicSensorDriver struct {
	pinTrig    string
	pinEcho    string
	name       string
	trigChan   chan int8
	echo       func(distance int64)
	connection gobot.Adaptor
	gobot.Commander
}

func NewUltrasonicSensorDriver(adaptor gobot.Adaptor, pinTrig, pinEcho string, echo func(distance int64)) *UltrasonicSensorDriver {
	return &UltrasonicSensorDriver{
		pinTrig:    pinTrig,
		pinEcho:    pinEcho,
		name:       gobot.DefaultName("Ultrasonic"),
		trigChan:   make(chan int8, 1),
		connection: adaptor,
		Commander:  gobot.NewCommander(),
		echo:       echo,
	}
}

func (u *UltrasonicSensorDriver) Name() string {
	return u.name
}

func (u *UltrasonicSensorDriver) SetName(s string) {
	u.name = s
}

func (u *UltrasonicSensorDriver) Start() error {
	log.Info("driver/UltrasonicSensorDriver: start ")
	go func() {
		for {
			select {
			case <-u.trigChan:
				log.Info("driver/UltrasonicSensorDriver: trig ")
				cr := checkHigh(u.connection.(gpio.DigitalReader), u.pinEcho)
				log.WithField("high check", cr).Info("driver/UltrasonicSensorDriver: trig and high check")
				if !cr {
					log.Warn("driver/UltrasonicSensorDriver: trig and high check fail")
					continue
				}
				begin := time.Now()
				checkLow(u.connection.(gpio.DigitalReader), u.pinEcho)
				d := time.Since(begin)
				distance := d.Milliseconds() * 17
				log.WithFields(log.Fields{
					"duration": d,
					"distance": distance,
				}).Info("driver/UltrasonicSensorDriver: receive ultrasonic")
				u.echo(distance)
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
