package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jumayevgadam/tsu-toleg/pkg/constants"
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
)

// ReadImage function checks image details and uploads.
func ReadImage(c *fiber.Ctx, field string) (*multipart.FileHeader, error) {
	const maxFileSize = 5 * 1024 * 1024

	file, err := c.FormFile(field)
	if err != nil {
		return nil, errlst.ErrNoUploadedFile
	}

	if file.Size > maxFileSize {
		return nil, errlst.ErrFileSize
	}

	return file, nil
}

func SaveImage(c *fiber.Ctx, file *multipart.FileHeader, facultyName, groupCode, studentName, username, semester string) (
	string, error,
) {
	// Base directory
	basePath := "./uploads"

	// Create base directory with proper permissions
	err := os.MkdirAll(basePath, constants.ZeroSevenFiveFive)
	if err != nil {
		return "", errlst.NewInternalServerError(fmt.Sprintf("failed to create base directory: %s", err.Error()))
	}

	cleanedFacultyName := strings.ReplaceAll(facultyName, " ", "_")

	// Create the subdirectory
	err = os.MkdirAll(basePath, constants.ZeroSevenFiveFive)
	if err != nil {
		return "", errlst.NewInternalServerError(fmt.Sprintf("failed to create subdirectory: %s", err.Error()))
	}

	uuidSuffix := uuid.New().String()[:5]

	// Final file path for saving
	fileDstPath := fmt.Sprintf(
		"%s/%s_%s_%s_%s_%s_%s%s",
		basePath,
		cleanedFacultyName,
		groupCode,
		semester,
		studentName,
		username,
		uuidSuffix,
		filepath.Ext(file.Filename),
	)

	// Save the file to the constructed path
	err = c.SaveFile(file, fileDstPath)
	if err != nil {
		return "", errlst.NewInternalServerError(fmt.Sprintf("failed to save image to path %s: %s", fileDstPath, err.Error()))
	}

	// Return a relative URL for the saved file
	return fmt.Sprintf(
		"%s_%s_%s_%s_%s_%s%s",
		cleanedFacultyName,
		groupCode,
		semester,
		studentName,
		username,
		uuidSuffix,
		filepath.Ext(file.Filename),
	), nil
}
