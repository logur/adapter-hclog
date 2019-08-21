package hclog_test

import (
	hclogadapter "logur.dev/adapter/hclog"
)

func ExampleNew() {
	var l interface{}
	logger := hclogadapter.New(l)

	// Output:
	_ = logger
}

// If logger is nil, a default instance is created.
func ExampleNew_default() {
	logger := hclogadapter.New(nil)

	// Output:
	_ = logger
}
