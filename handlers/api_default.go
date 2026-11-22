package handlers

import (
	"net/http"

	"github.com/boecko/prusalinkemu/models"
	"github.com/labstack/echo/v4"
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
	//var device, _ = findPrinter(c, printer)
	var capabilities = models.VersionCapabilities{UploadByPut: true}
	return ctx.JSON(http.StatusOK, models.Version{
		Api:          "2.0.0",
		Text:         "PrusaLink " + printer,
		Firmware:     "8.0.2+17092", // TODO from printer
		Printer:      "7.1.0",       // TODO from printer
		Capabilities: capabilities,
	})
}

func findPrinter(c *Container, wanted string) (deviceapi.Device, error) {
	return deviceapi.Device{}, nil
}
