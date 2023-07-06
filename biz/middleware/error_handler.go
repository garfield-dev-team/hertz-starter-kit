package middleware

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/pkg/errors"
)

func GlobalErrorHandler(ctx context.Context, c *app.RequestContext) {
	c.Next(ctx)

	if len(c.Errors) == 0 {
		return
	}

	hertzErr := c.Errors[0]
	// 获取errors包装的err
	err := hertzErr.Unwrap()
	// 打印异常堆栈
	logger.CtxErrorf(ctx, "%+v", err)
	// 获取原始err
	err = errors.Unwrap(err)
	// todo 进行错误类型判断
	c.JSON(consts.StatusBadRequest, utils.H{
		"code":    400,
		"message": err.Error(),
	})
}
