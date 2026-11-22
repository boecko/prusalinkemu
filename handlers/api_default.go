package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/boecko/prusalinkemu/models"
	"github.com/labstack/echo/v4"
	"github.com/prusa3d/prusa-connect/sdk/deviceapi"
	"github.com/prusa3d/prusa-connect/sdk/storageapi"
)

// PrinterApiV1FilesStoragePathPut - upload file or create folder
func (c *Container) PrinterApiV1FilesStoragePathPut(ctx echo.Context) error {
	printer := ctx.Param("printer")
	storage := ctx.Param("storage")
	path := ctx.Param("path")
	//printAfterUpload := ctx.Request().Header.Get("Print-After-Upload")
	var _, err = findPrinter(c, printer)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.HelloWorld{
			Message: "Printer not found",
		})
	}

	r := ctx.Request()
	bytes, _ := io.ReadAll(r.Body)
	var writeParams = storageapi.WriteParams{}
	writeParams.Path = path
	writeParams.Storage = storage
	writeParams.Size = int64(len(bytes))

	_, errW := c.storageapi.Write(c.context, writeParams, r.Body)
	if errW != nil {
		return errW
	}

	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Slicer doesn't realy care about this",
	})
}

// PrinterApiV1StorageGet - storage info
func (c *Container) PrinterApiV1StorageGet(ctx echo.Context) error {
	storage := models.Storage{
		Path:      "/prusaconnect",
		Name:      "PrusaConnect",
		Type:      "USB",
		ReadOnly:  false,
		Available: true,
	}

	return ctx.JSON(http.StatusOK, map[string][]models.Storage{
		"storage_list": {storage},
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

	fmt.Fprintln(os.Stderr, "DEVICE %w", device)

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
		Limit: 100,
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
