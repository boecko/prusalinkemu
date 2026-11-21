package handlers

import (
	"net/http"

	"github.com/boecko/prusalinkemu/models"
	"github.com/labstack/echo/v4"
)

// PrusalinkApiV1FilesStoragePathDelete - Delete a file or folder
func (c *Container) PrusalinkApiV1FilesStoragePathDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrusalinkApiV1FilesStoragePathGet - File or folder metadata
func (c *Container) PrusalinkApiV1FilesStoragePathGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrusalinkApiV1FilesStoragePathHead - file presence and state check
func (c *Container) PrusalinkApiV1FilesStoragePathHead(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrusalinkApiV1FilesStoragePathPost - Start print of file if there's no print job running
func (c *Container) PrusalinkApiV1FilesStoragePathPost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrusalinkApiV1FilesStoragePathPut - upload file or create folder
func (c *Container) PrusalinkApiV1FilesStoragePathPut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrusalinkApiV1StorageGet - storage info
func (c *Container) PrusalinkApiV1StorageGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrusalinkApiV1TransferGet - transfer info
func (c *Container) PrusalinkApiV1TransferGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrusalinkApiV1TransferIdDelete - stop transfer
func (c *Container) PrusalinkApiV1TransferIdDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// PrusalinkApiVersionGet - api version information
func (c *Container) PrusalinkApiVersionGet(ctx echo.Context) error {
	var capabilities = models.VersionCapabilities{UploadByPut: true}
	return ctx.JSON(http.StatusOK, models.Version{
		Api:          "2.0.0",
		Text:         "PrusaLink",
		Firmware:     "8.0.2+17092", // TODO from printer
		Printer:      "7.1.0",       // TODO from printer
		Capabilities: capabilities,
	})
}
