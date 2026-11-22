package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"

	"github.com/boecko/prusalinkemu/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prusa3d/prusa-connect/sdk/deviceapi"
	"github.com/prusa3d/prusa-connect/sdk/extension"
	"github.com/prusa3d/prusa-connect/sdk/fault"
	"github.com/prusa3d/prusa-connect/sdk/routesapi"
	"github.com/prusa3d/prusa-connect/sdk/storageapi"
	"github.com/prusa3d/prusa-connect/sdk/wire"
	"github.com/prusa3d/prusa-connect/sdk/wirerpc"
)

const (
	name    = "prusalinkemu"
	release = extension.Version("0.1.0")
)
const defaultMessage = "Prusalinkemu is serving."

type Settings struct {
	Message string `json:"message"`
}

//nolint:gochecknoglobals // the one piece of state this skeleton has.
var held atomic.Pointer[Settings]

//nolint:gochecknoglobals // whether the echo server in startPrusalink is listening.
var echoServerUp atomic.Bool

func settings() Settings {
	if current := held.Load(); current != nil {
		return *current
	}

	return Settings{Message: defaultMessage}
}

//go:embed openapi.json
var openAPI []byte

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
		os.Exit(1)
	}
}

func startPrusalink(ctx context.Context, devicesapi *deviceapi.Client, storageapi *storageapi.Client) error {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetOutput(os.Stderr)
	e.Logger.SetLevel(99)
	c, _ := handlers.NewContainer(ctx, devicesapi, storageapi)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// PrinterApiV1FilesStoragePathPut - upload file or create folder
	e.PUT("/:printer/api/v1/files/:storage/:path", c.PrinterApiV1FilesStoragePathPut)

	// PrinterApiV1StorageGet - storage info
	e.GET("/:printer/api/v1/storage", c.PrinterApiV1StorageGet)

	// PrinterApiVersionGet - api version information
	e.GET("/:printer/api/version", c.PrinterApiVersionGet)

	echoServerUp.Store(true)
	if err := e.Start("0.0.0.0:9090"); err != nil && err != http.ErrServerClosed {
		slog.Error("failed to start server", "error", err)
		echoServerUp.Store(false)
	}
	return nil
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	line, err := wire.Handshake{
		Protocol: wire.ProtocolMajor,
		MinMinor: wire.MinMinor,
		MaxMinor: wire.MaxMinor,
		Address:  wire.StdioAddress,
	}.Line()
	if err != nil {
		return fmt.Errorf("render the handshake: %w", err)
	}

	if _, err = os.Stdout.WriteString(line); err != nil {
		return fmt.Errorf("write the handshake: %w", err)
	}

	conn, err := wirerpc.NewConn(stdio{}, wirerpc.Options{
		Role:    wirerpc.RoleChild,
		Methods: methods(),
	})
	if err != nil {
		return fmt.Errorf("build the session: %w", err)
	}

	fmt.Fprintf(os.Stderr, "%s %s is serving on stdio as pid %d\n", name, release, os.Getpid())
	go func() {
		startPrusalink(ctx, deviceapi.NewClient(conn), storageapi.NewClient(conn))
	}()

	if err = conn.Serve(ctx); err != nil {
		return fmt.Errorf("serve: %w", err)
	}

	return nil
}

func methods() wirerpc.Methods {
	return wirerpc.Methods{
		wirerpc.MethodInfo: func(context.Context, wirerpc.Call) (any, error) {
			return routesapi.Describe(
				wirerpc.Info{
					Name:     name,
					Version:  release,
					Contract: wirerpc.Contract,
					Serves:   []wirerpc.Method{wirerpc.MethodConfigure},
				}, openAPI)
		},
		wirerpc.MethodConfigure: wirerpc.Configure(func(_ context.Context, sent Settings) error {
			if sent.Message == "" {
				sent.Message = defaultMessage
			}

			held.Store(&sent)

			return nil
		}),
		wirerpc.MethodShutdown: func(context.Context, wirerpc.Call) (any, error) {
			fmt.Fprintf(os.Stderr, "%s was asked to stop\n", name)

			return nil, nil
		},
		routesapi.MethodServe: routesapi.Serve(routes()),
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /status", func(w http.ResponseWriter, r *http.Request) {
		state := "stopped"
		if echoServerUp.Load() {
			state = "running"
		}

		answer(w, http.StatusOK, map[string]any{"state": state})
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		refuse(w, r, RouteNotFound.
			New("this module serves no route at "+r.URL.Path).
			With(fault.StringParam("path", r.URL.Path)))
	})

	return mux
}

func answer(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(body)
}

type stdio struct{}

func (stdio) Read(p []byte) (int, error)  { return os.Stdin.Read(p) }
func (stdio) Write(p []byte) (int, error) { return os.Stdout.Write(p) }
func (stdio) Close() error                { return os.Stdin.Close() }
