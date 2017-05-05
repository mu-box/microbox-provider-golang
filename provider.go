package provider

import (
)

var backend BackendAdaptor

// Start the provider using the provided back end
func Start(b BackendAdaptor, secure bool) error {
	if secure {
		mux.Use(redirect)
	}
	backend = b
	mux.Run(":8080")
	return nil
}