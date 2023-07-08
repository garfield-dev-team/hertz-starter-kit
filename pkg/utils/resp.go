package utils

import (
	"hertz-starter-kit/pkg/errcode"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type BaseResp struct {
	Code    int32    `json:"code"`
	Message string   `json:"message"`
	Result  any      `json:"result"`
	Details []string `json:"details,omitempty"`
}

type Response struct {
	Ctx *app.RequestContext
}

func NewResp(ctx *app.RequestContext) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) Success(data any) {
	if data == nil {
		data = utils.H{}
	}
	r.Ctx.JSON(consts.StatusOK, &BaseResp{Result: data})
}

func (r *Response) Error(err *errcode.ErrorCode) {
	resp := &BaseResp{
		Code:    err.GetCode(),
		Message: err.GetMessage(),
		Details: err.GetDetails(),
	}
	r.Ctx.JSON(consts.StatusOK, resp)
}
