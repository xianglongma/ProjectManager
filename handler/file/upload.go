package file

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/xianglongma/ProjectManager/pkg/resp"
	"log"
)

func (A APIImpl) Upload(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	log.Println("filename", file)
	id := uuid.NewV4()
	newName := id.String() + "-" + file.Filename
	ctx.SaveUploadedFile(file, "./database/"+newName)
	resp.SendData(ctx, gin.H{
		"filename": newName,
	})
}
