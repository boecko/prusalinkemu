package handlers

import (
	"bytes"
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
	"github.com/prusa3d/prusa-connect/sdk/wirerpc"
)

// PrinterApiV1FilesStoragePathPut - upload file or create folder
func (c *Container) PrinterApiV1FilesStoragePathPut(ctx echo.Context) error {
	printer := ctx.Param("printer")
	storage := ctx.Param("storage")
	path := ctx.Param("path")
	//printAfterUpload := ctx.Request().Header.Get("Print-After-Upload")

	listing, _ := c.storageapi.List(c.context, storageapi.ListParams{})
	fmt.Fprintln(os.Stderr, "Listing -", listing)
	fmt.Fprintln(os.Stderr, "Parameters - printer:", printer, "storage:", storage, "path:", path)

	var _, err = findPrinter(c, printer)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.HelloWorld{
			Message: "Printer not found",
		})
	}

	r := ctx.Request()
	fileBytes, _ := io.ReadAll(r.Body)
	var writeParams = storageapi.WriteParams{}
	writeParams.Path = path
	writeParams.Storage = "local"
	writeParams.Size = int64(len(fileBytes))
	fmt.Fprintf(os.Stderr, "writeParams - Path: %s, Storage: %s, Size: %d\n", writeParams.Path, writeParams.Storage, writeParams.Size)

	_, errW := c.storageapi.Write(c.context, writeParams, bytes.NewReader(fileBytes))
	if errW != nil {
		return errW
	}

	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Slicer doesn't realy care about this",
	})
}

// PrinterApiV1StorageGet - storage info
func (c *Container) PrinterApiV1StorageGet(ctx echo.Context) error {
	//printAfterUpload := ctx.Request().Header.Get("Print-After-Upload")

	//listing, err := c.storageapi.List(c.context, storageapi.ListParams{Storage: "local"})
	_, err := findStorages(c)
	if err != nil {
		return err
	}
	//fmt.Fprintln(os.Stderr, "Listing -", listing)

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

	//fmt.Fprintln(os.Stderr, "DEVICE %w", device)

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

func findStorages(c *Container) ([]string, error) {
	const MethodResources wirerpc.Method = "resources"
	resources, err := c.wirerpc.Call(c.context, MethodResources, nil)
	fmt.Fprintln(os.Stderr, "Resources -", resources)

	return nil, err
}
