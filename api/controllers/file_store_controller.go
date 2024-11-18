package controllers

import (
	"net/http"
	"rice-wine-shop/common/log"
	"rice-wine-shop/core/entities"
	"rice-wine-shop/core/services"

	"github.com/gin-gonic/gin"
)

type FileStoreController struct {
	file *services.FileStoreSerVice
}

func NewFileStoreController(file *services.FileStoreSerVice) *FileStoreController {
	return &FileStoreController{
		file: file,
	}
}

func (u *FileStoreController) DeleteFileByID(ctx *gin.Context) {
	id := GetIdFromParam(ctx, "fileID")
	err := u.file.DeleteFileByID(ctx, id, GetUserID(ctx))
	if err != nil {
		log.Error(err, "error")
		RespondError(ctx, http.StatusInternalServerError, err)
		return
	}
	RespondSuccess(ctx, nil)
}

func (u *FileStoreController) UploadFile(ctx *gin.Context) {
	var req entities.CreateUploadFileRequest
	if !BindAndValidate(ctx, &req) {
		return
	}
	err := u.file.AddListFile(ctx, GetUserID(ctx), &req)
	if err != nil {
		RespondError(ctx, http.StatusInternalServerError, err)
		return
	}
	RespondSuccess(ctx, nil)
}
