package asyncsuite

import (
	"context"
	"errors"
	"time"

	"github.com/stretchr/testify/suite"
)

type AsyncSuite struct {
	suite.Suite
	test *asyncTest
}

type asyncTest struct {
	ctx    context.Context
	cancel context.CancelCauseFunc
}

func newTest(timeout time.Duration) *asyncTest {
	res := &asyncTest{}
	ctxTimeout, _ := context.WithTimeout(context.TODO(), timeout)
	res.ctx, res.cancel = context.WithCancelCause(ctxTimeout)

	return res
}

func (s *AsyncSuite) InitTest(timeout time.Duration) {
	s.test = newTest(timeout)
}

func (s *AsyncSuite) DoneTest() {
	s.checkTest()
	s.test.cancel(nil)
	s.test = nil
}

func (s *AsyncSuite) Success() {
	s.Fail(nil)
}

func (s *AsyncSuite) Fail(err error) {
	s.checkTest()
	s.test.cancel(err)
}

func (s *AsyncSuite) Wait() {
	s.checkTest()
	<-s.test.ctx.Done()
}

func (s *AsyncSuite) WaitErr() error {
	s.checkTest()

	err := s.test.ctx.Err()
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}

func (s *AsyncSuite) TestErr() error {
	s.checkTest()

	err := context.Cause(s.test.ctx)
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}

func (s *AsyncSuite) checkTest() {
	if s.test == nil {
		panic("test not inited")
	}
}
