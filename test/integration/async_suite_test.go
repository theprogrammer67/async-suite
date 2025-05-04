package integration

// Basic imports
import (
	"async_suite/test/integration/asynctest"
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

// custom error in test
func (s *ExampleAsyncSuite) TestExample1() {
	async := asynctest.NewTest(&s.Suite, 10*time.Second)
	defer async.Done()

	start := func() {
		go func() {
			time.Sleep(100 * time.Millisecond)
			async.Fail("custom error")
		}()
	}

	start()
	async.Wait()

	s.Require().NoError(async.Err())
}

// NoError
func (s *ExampleAsyncSuite) TestExample2() {
	async := asynctest.NewTest(&s.Suite, 10*time.Second)
	defer async.Done()

	start := func() {
		go func() {
			time.Sleep(100 * time.Millisecond)
			err := errors.New("test error")
			async.NoError(err)
			async.Success()
		}()
	}

	start()
	async.Wait()
}

// Equal
func (s *ExampleAsyncSuite) TestExample3() {
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
}

// Assert
func (s *ExampleAsyncSuite) TestExample5() {
	async := asynctest.NewTest(&s.Suite, 10*time.Second)
	defer async.Done()

	start := func() {
		go func() {
			time.Sleep(100 * time.Millisecond)
			async.Assert(s.Equal(1, 2))
			async.Success()
		}()
	}

	start()
	async.Wait()
}

// context deadline exceeded error
func (s *ExampleAsyncSuite) TestExample4() {
	async := asynctest.NewTest(&s.Suite, 1*time.Second)
	defer async.Done()

	start := func() {
		go func() {
			time.Sleep(100 * time.Millisecond)
		}()
	}

	start()
	async.Wait()

	s.Require().NoError(async.Err())
}
