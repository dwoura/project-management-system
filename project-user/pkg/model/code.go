package model

import (
	"test.com/project-common/errs"
)

var (
	IllegalMobile = errs.NewError(2001, "手机号码不合法") //手机号码不合法
)
