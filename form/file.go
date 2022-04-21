package form

import (
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	filePkg "github.com/goer-project/goer-utils/file"
	"github.com/goer-project/goer/config"
)

func Save(c *gin.Context, file *multipart.FileHeader) (dir string, filename string, err error) {
	// Mkdir
	storagePath := config.NewDir.Upload
	dirName := fmt.Sprintf("/uploads/%s/", time.Now().Format("2006/01/02"))
	_ = os.MkdirAll(storagePath+dirName, 0755)

	// Random filename
	fileName := filePkg.RandomFilename(file)
	dir = storagePath + dirName
	path := dir + fileName
	if err := c.SaveUploadedFile(file, path); err != nil {
		return "", "", err
	}

	return dir, filename, nil
}

func SaveUploadedFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	// Save
	dir, filename, err := Save(c, file)
	if err != nil {
		return "", err
	}
	path := dir + filename

	return path, nil
}

func SaveAndCompress(c *gin.Context, file *multipart.FileHeader, compressedRatio float64, removeOldFile bool) (string, error) {
	// Save
	dir, filename, err := Save(c, file)
	if err != nil {
		return "", err
	}
	path := dir + filename

	if !IsResizable(c, file) {
		return path, nil
	}

	// Resize
	resizedPath, _ := filePkg.Resize(dir, filename, file, compressedRatio, removeOldFile)
	if resizedPath == "" {
		resizedPath = path
	}

	return resizedPath, nil
}

func IsResizable(c *gin.Context, header *multipart.FileHeader) bool {
	fileContent, err := header.Open()

	mime, err := filePkg.GetContentType(fileContent)
	if err != nil {
		return false
	}

	resizableTypes := []string{
		"image/jpg",
		"image/jpeg",
		"image/png",
	}

	for _, resizableType := range resizableTypes {
		if resizableType == mime {
			return true
		}
	}

	return false
}
