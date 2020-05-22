/**
 * @Author: entere@126.com
 * @Description:
 * @File:  演示海报生成过程
 * @Version: 1.0.0
 * @Date: 2020/5/21 21:54
 */

package main

import (
	"fmt"
	"github.com/hitailang/poster/core"
	"github.com/hitailang/poster/handler"
	"github.com/rs/xid"
)

// 海报组件使用示例
func main() {

	nullHandler := &handler.NullHandler{}
	ctx := &handler.Context{
		//图片都绘在这个PNG载体上
		PngCarrier: core.NewPNG(0, 0, 750, 1334),
	}
	//绘制背景图
	backgroundHandler := &handler.BackgroundHandler{
		X:    0,
		Y:    0,
		Path: "./assets/background.png",
	}
	//绘制圆形图像
	imageCircleHandler := &handler.ImageCircleHandler{
		X:   30,
		Y:   50,
		URL: "http://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTLJT9ncWLPov6rAzn4VCPSC4QoAvdangHRB1JgszqCvffggAysvzpm5MDb72Io4g9YAScHEw7xSWg/132",
	}
	//绘制本地图像
	imageLocalHandler := &handler.ImageLocalHandler{
		X:    30,
		Y:    400,
		Path: "./assets/reward.png",
	}

	//绘制二维码
	qrCodeHandler := &handler.QRCodeHandler{
		X:   30,
		Y:   860,
		URL: "https://github.com/hitailang/poster",
	}
	//绘制文字
	textHandler1 := &handler.TextHandler{
		Next:     handler.Next{},
		X:        180,
		Y:        105,
		Size:     20,
		R:        255,
		G:        241,
		B:        250,
		Text:     "如果觉得这个库对您有用",
		FontPath: "./assets/msyh.ttf",
	}
	//绘制文字
	textHandler2 := &handler.TextHandler{
		Next:     handler.Next{},
		X:        180,
		Y:        150,
		Size:     22,
		R:        255,
		G:        241,
		B:        250,
		Text:     "请随意赞赏~~",
		FontPath: "./assets/msyh.ttf",
	}
	//结束绘制，把前面的内容合并成一张图片
	endHandler := &handler.EndHandler{
		Output: "./build/poster_" + xid.New().String() + ".png",
	}

	// 链式调用绘制过程
	nullHandler.
		SetNext(backgroundHandler).
		SetNext(imageCircleHandler).
		SetNext(textHandler1).
		SetNext(textHandler2).
		SetNext(imageLocalHandler).
		SetNext(qrCodeHandler).
		SetNext(endHandler)

	// 开始执行业务
	if err := nullHandler.Run(ctx); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
	// 成功
	fmt.Println("Success")
	return

}
