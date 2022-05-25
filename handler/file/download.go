package file

import (
	"github.com/gin-gonic/gin"
	"github.com/xianglongma/ProjectManager/pkg/resp"
)

type DownloadRequest struct {
	FileName string `query:"filename"`
}

func (A APIImpl) Download(ctx *gin.Context) {
	filename := ctx.Query("filename")
	fileContentDisposition := "attachment;filename=\"" + filename + "\""
	ctx.Header("Content-Type", "application/zip") // 这里是压缩文件类型 .zip
	ctx.Header("Content-Disposition", fileContentDisposition)
	resp.SendSuccess(ctx)
}

func (A APIImpl) DownloadImg(ctx *gin.Context) {
	filename := ctx.Query("filename")
	ctx.File("./database/" + filename)
	resp.SendSuccess(ctx)
}
