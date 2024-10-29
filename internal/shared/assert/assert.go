package assert

func Assert(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}

func AssertFunc(condition func() bool, msg string) {
	if !condition() {
		panic(msg)
	}
}

func NotNil(obj interface{}, msg string) {
	if obj == nil {
		panic(msg)
	}
}

func Equal(expected, actual interface{}, msg string) {
	if expected != actual {
		panic(msg)
	}
}
