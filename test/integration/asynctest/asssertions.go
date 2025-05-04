package asynctest

func (t *AsyncTest) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if t.suite.Equal(expected, actual, msgAndArgs...) {
		return
	}
	t.cancel(nil)

	return
}

func (t *AsyncTest) NoError(err error, msgAndArgs ...interface{}) {
	if t.suite.NoError(err, msgAndArgs...) {
		return
	}
	t.cancel(nil)

	return
}
