package middleware

import (
	"context"
	"go.uber.org/zap"
	"hertz-starter-kit/pkg/errcode"
	"hertz-starter-kit/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
)

func GlobalErrorHandler(ctx context.Context, c *app.RequestContext) {
	c.Next(ctx)

	if len(c.Errors) == 0 {
		return
	}

	log := zap.L().Named("GlobalErrorHandler")

	hertzErr := c.Errors.Last()
	// 获取errors包装的err
	err := hertzErr.Unwrap()
	// 打印异常堆栈
	log.Error("", zap.Error(err))
	resp := utils.NewResp(c)

	// 注意这种写法断言失败会 panic
	//code := errors.Cause(err).(*errcode.ErrorCode)
	var code *errcode.ErrorCode
	// 这里可以用 `errors.As()` 做断言
	// 好处是会自动递归调用 `errors.Unwrap()`
	// 不需要再单独调用 `errors.Cause()` 函数进行递归 Unwrap 操作
	// errors.As 函数会将 err 递归地与目标类型进行比较，直到找到一个与目标类型匹配的错误
	// 如果找到一个与目标类型匹配的错误，errors.As 函数会将其赋值给传入的指针，并返回 true
	// 如果没有找到匹配的错误，errors.As 函数会将传入的指针赋值为 nil，并返回 false
	if errors.As(err, &code) {
		resp.Error(code)
	} else {
		resp.Error(errcode.NewServerError(errors.New("unknown error")))
	}
}
