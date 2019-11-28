package drivers

import (
	"eddy.org/pi/drivers/test"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type WheelDriverTestSuite struct {
	driver *WheelDriver
	suite.Suite
	checkChan <-chan test.CheckValue
}

func (w *WheelDriverTestSuite) SetupSuite() {
	c := make(chan test.CheckValue, 100)
	w.checkChan = c
	a := &test.Adaptor{
		N:         "test",
		WriteChan: c,
	}
	w.driver = NewWheelDriver(a, "1", "2")
}

func (s *WheelDriverTestSuite) TestWheelDriver() {

	s.T().Run("test toggle", func(t *testing.T) {
		assert := require.New(t)
		err := s.driver.Toggle()
		assert.NoError(err)
		assert.Equal(test.CheckValue{"1", 1}, <-s.checkChan)
		assert.Equal(test.CheckValue{"2", 0}, <-s.checkChan)
	})
	s.T().Run("test stop", func(t *testing.T) {
		assert := require.New(t)
		err := s.driver.Stop()
		assert.NoError(err)
		assert.Equal(test.CheckValue{"1", 0}, <-s.checkChan)
		assert.Equal(test.CheckValue{"2", 0}, <-s.checkChan)
	})
	s.T().Run("test front", func(t *testing.T) {
		assert := require.New(t)
		err := s.driver.Front()
		assert.NoError(err)
		assert.Equal(test.CheckValue{"1", 1}, <-s.checkChan)
		assert.Equal(test.CheckValue{"2", 0}, <-s.checkChan)
	})
	s.T().Run("test back", func(t *testing.T) {
		assert := require.New(t)
		err := s.driver.Back()
		assert.NoError(err)
		assert.Equal(test.CheckValue{"1", 0}, <-s.checkChan)
		assert.Equal(test.CheckValue{"2", 1}, <-s.checkChan)
	})
}

func TestWheelDriver(t *testing.T) {
	suite.Run(t, new(WheelDriverTestSuite))
}
