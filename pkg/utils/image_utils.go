package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// ReadImage function checks image details and uploads.
func ReadImage(c *fiber.Ctx, field string) (*multipart.FileHeader, error) {
	const maxFileSize = 5 * 1024 * 1024

	file, err := c.FormFile(field)
	if err != nil {
		return nil, errlst.NewBadRequestError(err.Error())
	}

	if file.Size > maxFileSize {
		return nil, errlst.NewBadRequestError("file size exceeds from limit 5MB")
	}

	return file, nil
}

func SaveImage(c *fiber.Ctx, file *multipart.FileHeader, facultyName, groupCode, studentName, username, semester string) (string, error) {
	// Base directory
	basePath := "./internal/uploads"

	err := os.MkdirAll(basePath, 0755)
	if err != nil {
		return "", errlst.NewInternalServerError(fmt.Sprintf("failed to create base directory: %s", err.Error()))
	}

	// cleanedFileName removes free spaces in file.Filename and replace with '_'
	cleanedFileName := strings.ReplaceAll(file.Filename, " ", "_")

	// Build subdirectory path
	subDir := fmt.Sprintf("%s/%s/%s", basePath, facultyName, groupCode)

	err = os.MkdirAll(subDir, 0755)
	if err != nil {
		return "", errlst.NewInternalServerError(fmt.Sprintf("failed to create group directory: %s", err.Error()))
	}

	// Final file path
	fileDstPath := fmt.Sprintf(
		"%s/%s_%s_%s",
		subDir,
		studentName,
		username,
		cleanedFileName,
	)

	// Save the file
	err = c.SaveFile(file, fileDstPath)
	if err != nil {
		return "", errlst.NewInternalServerError(fmt.Sprintf("failed to save image to path %s: %s", fileDstPath, err.Error()))
	}

	// Return logical file identifier
	return fmt.Sprintf("%s/%s/%s_%s_%s_%s%s", facultyName, groupCode, studentName, username, groupCode, semester, filepath.Ext(cleanedFileName)), nil
}
