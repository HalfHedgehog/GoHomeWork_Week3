// +build wireinject

/**
    @Author: qiyou_wu
    @CreateDate: 2022/4/10
    @Description:
**/

package main

import (
	"GoHomeWork_week3/app/user/service/internal/biz"
	"github.com/google/wire"
)

// InitializeEvent 声明injector的函数签名
func InitializeEvent(msg string) biz.Event {
	wire.Build(biz.NewEvent, biz.NewGreeter, biz.NewMessage)
	return biz.Event{} //返回值没有实际意义，只需符合函数签名即可
}
