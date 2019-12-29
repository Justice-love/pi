package drivers

import (
	"eddy.org/pi/drivers/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type CarDriverTestSuit struct {
	carDriver *CarDriver
	checkChan chan test.CheckValue
	suite.Suite
}

func (c *CarDriverTestSuit) SetupSuite() {
	cc := make(chan test.CheckValue, 100)
	a := &test.Adaptor{
		N:         "N",
		WriteChan: cc,
	}
	c.checkChan = cc
	c.carDriver = NewCarDriver(NewWheelDriver(
		a,
		"1",
		"2",
	), NewWheelDriver(
		a,
		"3",
		"4",
	))
}

func (c *CarDriverTestSuit) TestCarDriver() {
	c.T().Run("front test", func(t *testing.T) {
		assert := require.New(t)
		err := c.carDriver.Front()
		assert.NoError(err)
		assert.Equal(test.CheckValue{Pin: "1", Val: 1}, <-c.checkChan)
		assert.Equal(test.CheckValue{Pin: "2", Val: 0}, <-c.checkChan)
		assert.Equal(test.CheckValue{Pin: "3", Val: 1}, <-c.checkChan)
		assert.Equal(test.CheckValue{Pin: "4", Val: 0}, <-c.checkChan)
	})
	c.T().Run("back test", func(t *testing.T) {
		assert := require.New(t)
		err := c.carDriver.Back()
		assert.NoError(err)
		assert.Equal(test.CheckValue{Pin: "1", Val: 0}, <-c.checkChan)
		assert.Equal(test.CheckValue{Pin: "2", Val: 1}, <-c.checkChan)
		assert.Equal(test.CheckValue{Pin: "3", Val: 0}, <-c.checkChan)
		assert.Equal(test.CheckValue{Pin: "4", Val: 1}, <-c.checkChan)
	})
	c.T().Run("stop test", func(t *testing.T) {
		assert := require.New(t)
		err := c.carDriver.Stop()
		assert.NoError(err)
		assert.Equal(test.CheckValue{Pin: "1", Val: 0}, <-c.checkChan)
		assert.Equal(test.CheckValue{Pin: "2", Val: 0}, <-c.checkChan)
		assert.Equal(test.CheckValue{Pin: "3", Val: 0}, <-c.checkChan)
		assert.Equal(test.CheckValue{Pin: "4", Val: 0}, <-c.checkChan)
	})
	//c.T().Run("left test", func(t *testing.T) {
	//	assert := require.New(t)
	//	err := c.carDriver.Left()
	//	assert.NoError(err)
	//	assert.Equal(test.CheckValue{Pin: "1", Val: 1}, <-c.checkChan)
	//	assert.Equal(test.CheckValue{Pin: "2", Val: 0}, <-c.checkChan)
	//	assert.Equal(test.CheckValue{Pin: "3", Val: 0}, <-c.checkChan)
	//	assert.Equal(test.CheckValue{Pin: "4", Val: 1}, <-c.checkChan)
	//})
	//c.T().Run("right test", func(t *testing.T) {
	//	assert := require.New(t)
	//	err := c.carDriver.Right()
	//	assert.NoError(err)
	//	assert.Equal(test.CheckValue{Pin: "1", Val: 0}, <-c.checkChan)
	//	assert.Equal(test.CheckValue{Pin: "2", Val: 1}, <-c.checkChan)
	//	assert.Equal(test.CheckValue{Pin: "3", Val: 1}, <-c.checkChan)
	//	assert.Equal(test.CheckValue{Pin: "4", Val: 0}, <-c.checkChan)
	//})
	c.T().Run("every func test", func(t *testing.T) {
		c.carDriver.every(func() error {
			println("run")
			return nil
		})
		a := time.After(3 * time.Second)
		<-a
		st()
	})
}

func TestCarDriver(t *testing.T) {
	suite.Run(t, new(CarDriverTestSuit))
}
