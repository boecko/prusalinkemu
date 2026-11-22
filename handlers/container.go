package handlers

import (
	"context"

	"github.com/prusa3d/prusa-connect/sdk/deviceapi"
	"github.com/prusa3d/prusa-connect/sdk/storageapi"
	"github.com/prusa3d/prusa-connect/sdk/wirerpc"
)

// Container will hold all dependencies for your application.
type Container struct {
	devicesapi *deviceapi.Client
	storageapi *storageapi.Client
	context    context.Context
	wirerpc    *wirerpc.Conn
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer(ctx context.Context, devicesapi *deviceapi.Client, storageapi *storageapi.Client, conn *wirerpc.Conn) (Container, error) {
	c := Container{context: ctx, devicesapi: devicesapi, storageapi: storageapi, wirerpc: conn}
	return c, nil
}
