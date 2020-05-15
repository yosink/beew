package utils

import (
	"github.com/astaxie/beego/context"
)

type B struct {
	C *context.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (b *B) JsonResponse(httpCode, errCode int, message string, data interface{}) {
	b.C.Output.SetStatus(httpCode)
	b.C.Output.JSON(Response{
		Code:    errCode,
		Message: message,
		Data:    data,
	}, false, false)
}
