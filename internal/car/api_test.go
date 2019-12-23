package car

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CarServerSuit struct {
	server CarRpcServer
	check  chan Direction
	suite.Suite
}

func (c *CarServerSuit) SetupSuite() {
	c.server = NewCarServerAPI()
	c.check = make(chan Direction, 1)
	control = func(direction Direction) {
		c.check <- direction
	}
}

func (c *CarServerSuit) TestCarAPI() {
	c.T().Run("front", func(t *testing.T) {
		assert := require.New(t)
		_, _ = c.server.Command(context.Background(), &CarControlRequest{Direction: Direction_FRONT})
		assert.Equal(Direction_FRONT, <-c.check)
	})
	c.T().Run("back", func(t *testing.T) {
		assert := require.New(t)
		_, _ = c.server.Command(context.Background(), &CarControlRequest{Direction: Direction_BACK})
		assert.Equal(Direction_BACK, <-c.check)
	})
	c.T().Run("left", func(t *testing.T) {
		assert := require.New(t)
		_, _ = c.server.Command(context.Background(), &CarControlRequest{Direction: Direction_LEFT})
		assert.Equal(Direction_LEFT, <-c.check)
	})
	c.T().Run("right", func(t *testing.T) {
		assert := require.New(t)
		_, _ = c.server.Command(context.Background(), &CarControlRequest{Direction: Direction_RIGHT})
		assert.Equal(Direction_RIGHT, <-c.check)
	})

}

func TestCarAPI(t *testing.T) {
	suite.Run(t, new(CarServerSuit))
}
