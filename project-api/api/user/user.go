package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	common "test.com/project-common"
	loginServiceV1 "test.com/project-user/pkg/service/login_service.v1"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//project-api远程调用project-user里的方法
	rsp, err := LoginServiceClient.GetCaptcha(c, &loginServiceV1.CaptchaMessage{Mobile: mobile})
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(2001, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(rsp.Code))
}
