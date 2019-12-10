package drivers

import (
	"eddy.org/pi/drivers/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestM(t *testing.T) {
	assert := require.New(t)
	assert.True(int64(4)&mask != mask)
	assert.True(int64(5)&mask == mask)
	assert.True(int64(6)&mask != mask)
}

type avoidanceCarDriverTestSuit struct {
	avoidanceCarDriver *AvoidanceCarDriver
	checkChan          chan test.CheckValue
	readChan           chan int
	distanceChan       chan int64
	suite.Suite
}

func (a *avoidanceCarDriverTestSuit) SetupSuite() {
	cc := make(chan test.CheckValue, 100)
	rc := make(chan int, 100)
	dc := make(chan int64, 100)
	a.checkChan = cc
	a.readChan = rc
	a.distanceChan = dc
	ad := &test.Adaptor{
		N:         "N",
		WriteChan: cc,
		ReadChan:  rc,
	}
	cd := NewCarDriver(NewWheelDriver(
		ad,
		"1",
		"2",
	), NewWheelDriver(
		ad,
		"3",
		"4",
	))
	ud := NewUltrasonicSensorDriver(
		ad,
		"5",
		"6",
		func(distance int64) {
			dc <- distance
		},
	)
	a.avoidanceCarDriver = NewAvoidanceCarDriver(cd, ud)
}

func (a *avoidanceCarDriverTestSuit) TestAvoidanceCarDriver() {
	assert := require.New(a.T())
	err := a.avoidanceCarDriver.Start()
	assert.NoError(err)

	err = a.avoidanceCarDriver.carDriver.Front()
	assert.NoError(err)
	assert.Equal(test.CheckValue{Pin: "1", Val: 1}, <-a.checkChan)
	assert.Equal(test.CheckValue{Pin: "2", Val: 0}, <-a.checkChan)
	assert.Equal(test.CheckValue{Pin: "3", Val: 1}, <-a.checkChan)
	assert.Equal(test.CheckValue{Pin: "4", Val: 0}, <-a.checkChan)

	err = a.avoidanceCarDriver.ultrasonicSensorDriver.Trig()
	assert.NoError(err)
	assert.Equal(test.CheckValue{Pin: "5", Val: 1}, <-a.checkChan)
	assert.Equal(test.CheckValue{Pin: "5", Val: 0}, <-a.checkChan)

	a.readChan <- 1
	c := time.After(800 * time.Microsecond)
	<-c
	a.readChan <- 0
	go a.avoidanceCarDriver.Avoidance(a.distanceChan)
	time.Sleep(100 * time.Millisecond)

	assert.Equal(test.CheckValue{Pin: "1", Val: 0}, <-a.checkChan)
	assert.Equal(test.CheckValue{Pin: "2", Val: 0}, <-a.checkChan)
	assert.Equal(test.CheckValue{Pin: "3", Val: 0}, <-a.checkChan)
	assert.Equal(test.CheckValue{Pin: "4", Val: 0}, <-a.checkChan)
}

func TestAvoidanceCarDriver(t *testing.T) {
	suite.Run(t, new(avoidanceCarDriverTestSuit))
}
