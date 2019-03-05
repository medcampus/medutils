package medutils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SuccessResponse struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func ReturnSuccessResponse(ctx *gin.Context, response SuccessResponse) {
	response.Status = true
	if response.Code != 0 {
		response.Code = http.StatusOK
	}
	ctx.JSON(response.Code, response)
}

func ReturnErrorResponse(ctx *gin.Context, response ErrorResponse) {
	response.Status = false
	ctx.Error(errors.New(response.Message))
	ctx.JSON(response.Code, response)

}
