package asyncsuite

func (s *AsyncSuite) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	s.checkTest()

	if s.Suite.Equal(expected, actual, msgAndArgs...) {
		return
	}
	s.test.cancel(nil)

	return
}
