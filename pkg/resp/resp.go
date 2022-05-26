package resp

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Action  string      `json:"action"`
	Data    interface{} `json:"data,omitempty"`
}

func SendData(c *gin.Context, data interface{}) {
	c.JSON(200, &Response{
		Code: 0,
		Data: data,
	})
}

func SendSuccess(ctx *gin.Context) {
	SendData(ctx, map[string]string{"result": "success"})
}

func SendError(c *gin.Context, err error) {
	c.JSON(200, &Response{
		Code:    10000,
		Message: err.Error(),
		Action:  "error",
	})
}
