package file

import "github.com/gin-gonic/gin"

type API interface {
	Upload(ctx *gin.Context)
	Download(ctx *gin.Context)
	DownloadImg(ctx *gin.Context)
}

type APIImpl struct {
}

func NewFileAPI() API {
	return APIImpl{}
}
