package integration

// Basic imports
import (
	asyncsuite "async_suite/test/integration/async_suite"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ExampleAsyncSuite struct {
	suite.Suite
	async *asyncsuite.AsyncSuite
}

func (s *ExampleAsyncSuite) SetupSuite() {
	s.async.SetT(s.T())
}

func (s *ExampleAsyncSuite) TearDownSuite() {
}

func TestAsyncSuite(t *testing.T) {
	suite.Run(t, &ExampleAsyncSuite{async: &asyncsuite.AsyncSuite{}})
}

func (s *ExampleAsyncSuite) TestExample() {
	s.async.InitTest(10 * time.Second)
	defer s.async.DoneTest()

	start := func() {
		go func() {
			time.Sleep(100 * time.Millisecond)
			s.async.Equal(99, 100)
			// s.async.Fail(errors.New("test error"))
			s.async.Success()
		}()
	}

	start()
	s.async.Wait()

	s.Require().NoError(s.async.WaitErr())
	s.Require().NoError(s.async.TestErr())
}
