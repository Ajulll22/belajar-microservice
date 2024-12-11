package handling

import "github.com/gin-gonic/gin"

func ResponseSuccess(c *gin.Context, data interface{}, msg string, code int) BaseResponse {
	return BaseResponse{
		Message: msg,
		Code:    code,
		Data:    data,
	}
}
