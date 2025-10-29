package response

import (
	"net/http"
	"time"

	"github.com/binhbeng/goex/internal/pkg/errors"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Cost string      `json:"cost"`
}

type Response struct {
	httpCode int
	result   *Result
}

func Resp() *Response {
	return &Response{
		httpCode: http.StatusOK,
		result: &Result{
			Code: 0,
			Msg:  "",
			Data: nil,
			Cost: "",
		},
	}
}

func (r *Response) Fail(c *gin.Context, code int, msg string, data ...any) {
	r.SetCode(code)
	r.SetMessage(msg)
	if data != nil {
		r.WithData(data[0])
	}
	r.json(c)
}

func (r *Response) FailCode(c *gin.Context, code int, msg ...string) {
	r.SetCode(code)
	if msg != nil {
		r.SetMessage(msg[0])
	}
	r.json(c)
}

func (r *Response) Success(c *gin.Context) {
	r.SetCode(errors.SUCCESS)
	r.json(c)
}

func (r *Response) WithDataSuccess(c *gin.Context, data interface{}) {
	r.SetCode(errors.SUCCESS)
	r.WithData(data)
	r.json(c)
}

func (r *Response) SetCode(code int) *Response {
	r.result.Code = code
	return r
}

func (r *Response) SetHttpCode(code int) *Response {
	r.httpCode = code
	return r
}

type defaultRes struct {
	Result any `json:"result"`
}

func (r *Response) WithData(data any) *Response {
	switch data.(type) {
	case string, int, bool:
		r.result.Data = &defaultRes{Result: data}
	default:
		r.result.Data = data
	}
	return r
}

func (r *Response) SetMessage(message string) *Response {
	r.result.Msg = message
	return r
}

var ErrorText = errors.NewErrorText("zh_CN")

func (r *Response) json(c *gin.Context) {
	if r.result.Msg == "" {
		r.result.Msg = ErrorText.Text(r.result.Code)
	}

	// if r.Data == nil {
	// 	r.Data = struct{}{}
	// }

	r.result.Cost = time.Since(c.GetTime("requestStartTime")).String()
	c.AbortWithStatusJSON(r.httpCode, r.result)
}

func Success(c *gin.Context, data ...any) {
	if data != nil {
		Resp().WithDataSuccess(c, data[0])
		return
	}
	Resp().Success(c)
}

func FailCode(c *gin.Context, code int, data ...any) {
	if data != nil {
		Resp().WithData(data[0]).FailCode(c, code)
		return
	}
	Resp().FailCode(c, code)
}

func Fail(c *gin.Context, code int, message string, data ...any) {
	if data != nil {
		Resp().WithData(data[0]).FailCode(c, code, message)
		return
	}
	Resp().FailCode(c, code, message)
}