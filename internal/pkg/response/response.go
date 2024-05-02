package response

import (
	customErr "github.com/AkbarFikri/mooistudio_technical_test/internal/pkg/error"
	"github.com/gin-gonic/gin"
)

type Response struct {
	HttpCode int         `json:"-"`
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	Payload  interface{} `json:"payload,omitempty"`
	Error    string      `json:"error,omitempty"`
}

func New(params ...func(*Response) *Response) Response {
	var resp = Response{
		Success: true,
	}

	for _, param := range params {
		param(&resp)
	}

	return resp
}

func WithHttpCode(httpCode int) func(*Response) *Response {
	return func(r *Response) *Response {
		r.HttpCode = httpCode
		return r
	}
}

func WithMessage(message string) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Message = message
		return r
	}
}

func WithPayload(payload interface{}) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Payload = payload
		return r
	}
}

func WithError(err error) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Success = false

		myErr, ok := err.(customErr.Error)
		if !ok {
			myErr = customErr.ErrorGeneral
		}

		r.Error = myErr.Message
		r.HttpCode = myErr.HttpCode

		return r
	}
}

func (r Response) Send(ctx *gin.Context) {
	ctx.JSON(r.HttpCode, r)
	return
}
