package controllers

import (
	"fmt"
	"int_ecosys/services"
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

// UploadHandler will handle user's upload.
// Accept mutiple files upload
// 1. Store uploaded scripts
// 2. Process these scripts, when the complete uploading api is called
// 3. Run docker
// TODO:
// 1. Process files need to add require in these scripts according to some conditionn
func (uploader *UploadController) UploadHandler(c echo.Context) error {
	fmt.Printf("Hello Uploader %v\n", c.Request())
	fileType := c.FormValue("type")
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fmt.Printf("==>Upload type: %s, file name: %s\n", fileType, file.Filename)

	// Store files
	if err := uploader.storeFile(file); err != nil {
		return err
	}
	// Execute runner
	services.RunParser()

	return c.HTML(http.StatusOK,
		fmt.Sprintf("<p>File %s uploaded successfully with fields.</p>",
			file.Filename))
}

// UploadCompleteHandler indicates that upload of all files are done
func (uploader *UploadController) UploadCompleteHandler(c echo.Context) error {
	var errorMessage string
	ret := models.Ret{
		Status:  "200",
		Message: "Done",
	}
	// Process files
	parser, err := services.NewParser("", "")
	if err != nil {
		errorMessage = fmt.Sprintf("Process uploaded files error %v", err)
		ret.Status = "400"
		ret.Message = errorMessage

		return c.JSON(http.StatusOK, ret)
	}

	err = parser.RunParser()
	if err != nil {
		errorMessage = fmt.Sprintf("Process uploaded files error %v", err)
		ret.Status = "400"
		ret.Message = errorMessage

		return c.JSON(http.StatusOK, ret)
	}

	// Execute runner
	err = services.Run()
	if err != nil {
		errorMessage = fmt.Sprintf("Process uploaded files error %v", err)
		ret.Status = "400"
		ret.Message = errorMessage

		return c.JSON(http.StatusOK, ret)
	}

	return c.JSON(http.StatusOK, ret)
}

// TODO: move this function to services
func (uploader *UploadController) storeFile(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	storePath, err := filepath.Abs(fmt.Sprintf("./store/temp/%s", file.Filename))
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
