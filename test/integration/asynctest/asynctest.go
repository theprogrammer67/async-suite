package asynctest

import (
	"context"
	"errors"
	"time"

	"github.com/stretchr/testify/suite"
)

type AsyncTest struct {
	suite  *suite.Suite
	ctx    context.Context
	cancel context.CancelCauseFunc
}

func NewTest(suite *suite.Suite, timeout time.Duration) *AsyncTest {
	t := &AsyncTest{suite: suite}
	ctxTimeout, _ := context.WithTimeout(context.TODO(), timeout)
	t.ctx, t.cancel = context.WithCancelCause(ctxTimeout)

	return t
}

func (t *AsyncTest) Done() {
	t.cancel(nil)
}

func (t *AsyncTest) Success() {
	t.cancel(nil)
}

func (t *AsyncTest) Fail(message string) {
	t.cancel(errors.New(message))
}

func (t *AsyncTest) Wait() {
	<-t.ctx.Done()
}

func (t *AsyncTest) Assert(b bool) {
	if !b {
		t.cancel(nil)
	}
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
