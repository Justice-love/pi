package drivers

import (
	"eddy.org/pi/drivers/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type UltrasonicSensorDriverTestSuite struct {
	driver       *UltrasonicSensorDriver
	readChan     chan int
	checkChan    <-chan test.CheckValue
	distanceChan chan int64
	suite.Suite
}

func (u *UltrasonicSensorDriverTestSuite) SetupSuite() {
	c := make(chan int, 100)
	cc := make(chan test.CheckValue, 100)
	dc := make(chan int64, 100)
	u.driver = NewUltrasonicSensorDriver(&test.Adaptor{
		N:         "test",
		ReadChan:  c,
		WriteChan: cc,
	}, "1", "2", func(distance int64) {
		dc <- distance
	})
	u.readChan = c
	u.checkChan = cc
	u.distanceChan = dc
}

func (u *UltrasonicSensorDriverTestSuite) TestUltrasonicSensorDriver() {
	assert := require.New(u.T())
	err := u.driver.Start()
	assert.NoError(err)
	u.T().Run("test trig", func(t *testing.T) {
		assert := require.New(t)
		err := u.driver.Trig()
		assert.NoError(err)
		assert.Equal(test.CheckValue{Pin: "1", Val: 1}, <-u.checkChan)
		assert.Equal(test.CheckValue{Pin: "1", Val: 0}, <-u.checkChan)
	})
	u.T().Run("test echo", func(t *testing.T) {
		assert := require.New(t)
		u.readChan <- 1
		tc := time.After(5 * time.Second)
		<-tc
		u.readChan <- 0
		df := <-u.distanceChan
		assert.Equal(int64(85000), df&0x1FFC8)
	})
	u.T().Run("test long", func(t *testing.T) {
		assert := require.New(t)
		err := u.driver.Trig()
		assert.NoError(err)
		time.Sleep(1 * time.Second)
		assert.Empty(0, len(u.distanceChan))
	})

}

func TestUltrasonicSensorDriver(t *testing.T) {
	suite.Run(t, new(UltrasonicSensorDriverTestSuite))
}
