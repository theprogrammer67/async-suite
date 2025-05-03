package integration

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

var (
	_ suite.SetupAllSuite    = (*TestSuite)(nil)
	_ suite.TearDownAllSuite = (*TestSuite)(nil)
)

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) SetupSuite() {
}

func (s *TestSuite) TearDownSuite() {
}

func Test(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) TestAsync() {
	ctxTimeout, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	ctx, cancel := context.WithCancelCause(ctxTimeout)
	defer cancel(nil)

	v := 100

	start := func() {
		go func() {
			time.Sleep(time.Second)
			if !s.Equal(100, v) {
				cancel(nil)
			}
			// cancel(errors.New("test-error"))
			cancel(nil)
		}()

	}

	start()
	<-ctx.Done()

	s.Require().ErrorIs(ctx.Err(), context.Canceled, "error waiting test")
	s.Require().ErrorIs(context.Cause(ctx), context.Canceled, "test error")
}
