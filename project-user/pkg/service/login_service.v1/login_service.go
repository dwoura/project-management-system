package login_service_v1

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"log"
	common "test.com/project-common"
	"test.com/project-user/pkg/dao"
	"test.com/project-user/pkg/repo"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}

}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	//1、获取参数
	mobile := msg.Mobile
	//2、校验参数
	if !common.VerifyMobile(mobile) {
		return nil, errors.New("手机号不合法")
	}
	//3、生成验证码 （随机4位1000-9999）
	code := "123456"
	//4、调用短信平台（三方 放入go协程中执行 接口可以快速响应）
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("短信平台调用成功，发送短信 INFO")
		//5、存储验证码 redis当中 过期时间15分钟
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := ls.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Printf("验证码存入redis出错：%v \n", err)
		}
	}()
	return &CaptchaResponse{Code: code}, nil
}
