package controllers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// UploadController handle user's upload
type UploadController struct {
}

// UploadHandler will handle user's upload
func (uploader *UploadController) UploadHandler(c echo.Context) error {
	fmt.Println("Hello Uploader")
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	if err := uploader.storeFile(file); err != nil {
		return err
	}

	return c.HTML(http.StatusOK,
		fmt.Sprintf("<p>File %s uploaded successfully with fields.</p>",
			file.Filename))
}

func (uploader *UploadController) storeFile(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	storePath, err := filepath.Abs(fmt.Sprintf("./store/%s", file.Filename))
	if err != nil {
		return fmt.Errorf("Get store path error %v", err)
	}
	dst, err := os.Create(storePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
