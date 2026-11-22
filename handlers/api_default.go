package handlers

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/boecko/prusalinkemu/models"
	"github.com/labstack/echo/v4"
	"github.com/prusa3d/prusa-connect/sdk/deviceapi"
)

// PrinterApiV1FilesStoragePathDelete - Delete a file or folder
func (c *Container) PrinterApiV1FilesStoragePathDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrinterApiV1FilesStoragePathGet - File or folder metadata
func (c *Container) PrinterApiV1FilesStoragePathGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrinterApiV1FilesStoragePathHead - file presence and state check
func (c *Container) PrinterApiV1FilesStoragePathHead(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrinterApiV1FilesStoragePathPost - Start print of file if there's no print job running
func (c *Container) PrinterApiV1FilesStoragePathPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrinterApiV1FilesStoragePathPut - upload file or create folder
func (c *Container) PrinterApiV1FilesStoragePathPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrinterApiV1StorageGet - storage info
func (c *Container) PrinterApiV1StorageGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrinterApiV1TransferGet - transfer info
func (c *Container) PrinterApiV1TransferGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrinterApiV1TransferIdDelete - stop transfer
func (c *Container) PrinterApiV1TransferIdDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrinterApiVersionGet - api version information
func (c *Container) PrinterApiVersionGet(ctx echo.Context) error {
	printer := ctx.Param("printer")
	var device, err = findPrinter(c, printer)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.HelloWorld{
			Message: "Printer not found",
		})
	}

	fmt.Fprintf(os.Stderr, "DEVICE %s\n", device)

	var capabilities = models.VersionCapabilities{UploadByPut: true}
	return ctx.JSON(http.StatusOK, models.Version{
		Api:          "2.0.0",
		Text:         "PrusaLink " + printer,
		Firmware:     device.Firmware,
		Printer:      "7.1.0",
		Capabilities: capabilities,
	})
}

func toSlug(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "")
	s = strings.Trim(s, "")
	return s
}

func findPrinter(c *Container, wanted string) (deviceapi.Device, error) {
	listparams := deviceapi.ListParams{
		Limit: 10,
	}
	page, err := c.devicesapi.List(c.context, listparams)
	if err != nil {
		fmt.Fprintln(os.Stderr, "DEVICES err", err)
	}
	if len(page.Items) == 0 {
		return deviceapi.Device{}, fmt.Errorf("no devices found")
	}
	wantedSlug := toSlug(wanted)
	for _, device := range page.Items {
		deviceSlug := toSlug(device.Name)
		if deviceSlug == wantedSlug {
			return device, nil
		}
	}
	return deviceapi.Device{}, fmt.Errorf("no device found with name %s", wanted)
}
