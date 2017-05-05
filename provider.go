package provider

import (
	"fmt"
	"github.com/plimble/ace"
)

var backend BackendAdaptor

// Start the provider using the provided back end
func Start(b BackendAdaptor, secure bool) error {
	mux = ace.New()
	setup()
	if secure {
		fmt.Println("secure??")
		mux.Use(redirect, ace.Logger())
	}
	backend = b
	mux.Run(":8080")
	return nil
}