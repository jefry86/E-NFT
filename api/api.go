package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nft_platform/global"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Api struct {
	*gin.Context
}

func (a *Api) c(c *gin.Context) *Api {
	if a.Context == nil {
		a.Context = c
	}
	return a
}

func (a *Api) JsonWithCodeAndData(c *gin.Context, code global.Code, data interface{}) {
	result := ResponseData{
		Code: int(code),
		Msg:  code.String(),
		Data: data,
	}
	a.c(c).JSON(http.StatusOK, result)
}

func (a *Api) JsonSuccessWithData(c *gin.Context, data interface{}) {
	ok := global.OK
	if data == nil {
		data = map[string]string{}
	}

	result := ResponseData{
		Code: int(ok),
		Msg:  ok.String(),
		Data: data,
	}
	a.c(c).JSON(http.StatusOK, result)
}

func (a *Api) JsonSuccess(c *gin.Context) {
	ok := global.OK
	result := ResponseData{
		Code: int(ok),
		Msg:  ok.String(),
		Data: map[string]interface{}{},
	}
	a.c(c).JSON(http.StatusOK, result)
}

func (a *Api) JsonError(c *gin.Context) {
	fail := global.Fail
	result := ResponseData{
		Code: int(fail),
		Msg:  fail.String(),
		Data: map[string]interface{}{},
	}
	a.c(c).JSON(http.StatusOK, result)
}

func (a *Api) JsonParamsError(c *gin.Context) {
	fail := global.ParamErr
	result := ResponseData{
		Code: int(fail),
		Msg:  fail.String(),
		Data: map[string]interface{}{},
	}
	a.c(c).JSON(http.StatusOK, result)
}

func (a *Api) JsonErrorWithMsg(c *gin.Context, msg string) {
	fail := global.ParamErr
	result := ResponseData{
		Code: int(fail),
		Msg:  msg,
		Data: map[string]interface{}{},
	}
	a.c(c).JSON(http.StatusOK, result)
}

func (a *Api) JsonWithCode(c *gin.Context, code global.Code) {
	result := ResponseData{
		Code: int(code),
		Msg:  code.String(),
		Data: map[string]interface{}{},
	}
	a.c(c).JSON(http.StatusOK, result)
}

func (a *Api) JsonWithData(c *gin.Context, code int, msg string, data interface{}) {
	result := ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	a.c(c).JSON(http.StatusOK, result)
}

func (a *Api) Json(c *gin.Context, code int, msg string) {
	result := ResponseData{
		Code: code,
		Msg:  msg,
		Data: map[string]interface{}{},
	}
	a.c(c).JSON(http.StatusOK, result)
}
