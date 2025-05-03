package integration

// Basic imports
import (
	asynctest "async_suite/test/integration/async_test"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ExampleAsyncSuite struct {
	suite.Suite
}

func (s *ExampleAsyncSuite) SetupSuite() {
}

func (s *ExampleAsyncSuite) TearDownSuite() {
}

func TestAsyncSuite(t *testing.T) {
	suite.Run(t, &ExampleAsyncSuite{})
}

func (s *ExampleAsyncSuite) TestExample1() {
	async := asynctest.NewTest(&s.Suite, 10*time.Second)
	defer async.Done()

	start := func() {
		go func() {
			time.Sleep(100 * time.Millisecond)
			async.Fail(errors.New("test error"))
		}()
	}

	start()
	async.Wait()

	s.Require().NoError(async.Err())
}

func (s *ExampleAsyncSuite) TestExample2() {
	async := asynctest.NewTest(&s.Suite, 10*time.Second)
	defer async.Done()

	start := func() {
		go func() {
			time.Sleep(100 * time.Millisecond)
			async.Equal(99, 100)
			async.Success()
		}()
	}

	start()
	async.Wait()

	s.Require().NoError(async.Err())
}
