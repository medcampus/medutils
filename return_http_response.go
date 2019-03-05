package medutils

import (
	"errors"
	"github.com/gin-gonic/gin"
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
	ctx.JSON(response.Code, response)
}

func ReturnErrorResponse(ctx *gin.Context, response ErrorResponse) {
	response.Status = false
	ctx.Error(errors.New(response.Message))
	ctx.JSON(response.Code, response)

}
