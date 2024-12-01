package uploader

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context, file *multipart.FileHeader) (filePath string, err error) {

	if _, err := os.Stat("storage"); os.IsNotExist(err) {
		os.Mkdir("storage", os.ModePerm)
	}

	// Menghasilkan UUID untuk nama file yang unik
	newFileName := uuid.New().String() + filepath.Ext(file.Filename)

	// Tentukan folder penyimpanan file (misalnya di folder "storage")
	filePath = filepath.Join("storage", newFileName)

	// Simpan file ke folder storage
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}
