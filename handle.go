package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const UploadDir = "./uploads"

func HandleUpload(c *gin.Context) {
	file, header, err := c.Request.FormFile("File")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file field is bad"})
		return
	}
	defer file.Close()

	outPath := filepath.Join(UploadDir, header.Filename)

	if err := c.SaveUploadedFile(header, outPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't save file"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded succesfully",
		"filename": header.Filename,
		"path":     outPath,
	})

}
