package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"rice-wine-shop/common/utils"
	"rice-wine-shop/core/configs"

	"github.com/gin-gonic/gin"
)

type ControllerSaveFile struct {
}

func NewControllerSaveFile() *ControllerSaveFile {
	return &ControllerSaveFile{}
}

func (u *ControllerSaveFile) SaveFile(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Could not retrieve form data"})
		return
	}

	files, ok := form.File["upload[]"]
	if !ok || len(files) == 0 {
		ctx.JSON(400, gin.H{"error": "No files were uploaded"})
		return
	}

	var fileList []string
	dstFolder := "publics"

	if err := os.MkdirAll(dstFolder, os.ModePerm); err != nil {
		ctx.JSON(500, gin.H{"error": "Could not create directory"})
		return
	}

	for _, file := range files {
		genNameFile := utils.GenerateNameFile() + filepath.Ext(file.Filename)
		dst := filepath.Join(dstFolder, genNameFile)

		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		fileURL := fmt.Sprintf("%s/%s", configs.Get().ConfigPathFile, genNameFile)
		fileList = append(fileList, fileURL)
	}

	ctx.JSON(200, gin.H{
		"code":           0,
		"uploaded_files": fileList,
	})
}
