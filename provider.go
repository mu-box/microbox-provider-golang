package provider

import (
)

var backend BackendAdaptor

// Start the provider using the provided back end
func Start(b BackendAdaptor) error {
	backend = b
	return mux.Run(":8080")
}