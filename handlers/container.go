package handlers

import (
	"context"

	"github.com/prusa3d/prusa-connect/sdk/deviceapi"
)

// Container will hold all dependencies for your application.
type Container struct {
	devicesapi *deviceapi.Client
	context    context.Context
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer(ctx context.Context, devicesapi *deviceapi.Client) (Container, error) {
	c := Container{context: ctx, devicesapi: devicesapi}
	return c, nil
}
