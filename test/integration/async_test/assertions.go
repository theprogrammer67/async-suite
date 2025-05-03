package asynctest

func (t *AsyncTest) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if t.suite.Equal(expected, actual, msgAndArgs...) {
		return
	}
	t.cancel(nil)

	return
}
