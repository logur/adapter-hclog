package hclog_test

import (
	"github.com/hashicorp/go-hclog"

	hclogadapter "logur.dev/adapter/hclog"
)

func ExampleNew() {
	logger := hclogadapter.New(hclog.Default())

	// Output:
	_ = logger
}

// If logger is nil, a default instance is created.
func ExampleNew_default() {
	logger := hclogadapter.New(nil)

	// Output:
	_ = logger
}
