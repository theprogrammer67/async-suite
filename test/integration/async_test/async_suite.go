package asynctest

import (
	"context"
	"errors"
	"time"

	"github.com/stretchr/testify/suite"
)

// type AsyncSuite struct {
// 	suite.Suite
// 	test *asyncTest
// }

type AsyncTest struct {
	suite  *suite.Suite
	ctx    context.Context
	cancel context.CancelCauseFunc
}

func NewTest(suite *suite.Suite, timeout time.Duration) *AsyncTest {
	res := &AsyncTest{suite: suite}
	ctxTimeout, _ := context.WithTimeout(context.TODO(), timeout)
	res.ctx, res.cancel = context.WithCancelCause(ctxTimeout)

	return res
}

// func (s *AsyncSuite) InitTest(timeout time.Duration) {
// 	s.test = newTest(timeout)
// }

func (t *AsyncTest) Done() {
	t.cancel(nil)
}

func (t *AsyncTest) Success() {
	t.Fail(nil)
}

func (t *AsyncTest) Fail(err error) {
	t.cancel(err)
}

func (t *AsyncTest) Wait() {
	<-t.ctx.Done()
}

func (t *AsyncTest) waitErr() error {
	err := t.ctx.Err()
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}

func (t *AsyncTest) testErr() error {
	err := context.Cause(t.ctx)
	if errors.Is(err, context.Canceled) {
		return nil
	}

	return err
}

func (t *AsyncTest) Err() error {
	err := t.waitErr()
	if err != nil {
		return err
	}

	err = t.testErr()
	if err != nil {
		return err
	}

	return nil
}
