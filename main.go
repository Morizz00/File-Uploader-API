package main

import (
	"log"
	"os"

	"github.com/Morizz00/file-upload-api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := os.MkdirAll(handler.UploadDir, os.ModePerm); err != nil {
		log.Fatalf("Error to create upload dir: %v", err)
	}

	router := gin.Default()
	router.MaxMultipartMemory = 10 << 20

	router.POST("/upload", handler.HandleUpload)
	log.Println("Server running at http://localhost:8086")
	router.Run(":8086")
}
