/**
 * @Author: entere@126.com
 * @Description:
 * @File:  qrcode_handler
 * @Version: 1.0.0
 * @Date: 2020/5/21 22:16
 */

package handler

import (
	"fmt"
	"github.com/hitailang/poster/core"
	"github.com/skip2/go-qrcode"
	"image"
)

// BackgroundHandler
type QRCodeHandler struct {
	// 合成复用Next
	Next
	X   int
	Y   int
	URL string
}

// Do 地址逻辑
func (h *QRCodeHandler) Do(c *Context) (err error) {
	//生成二维码
	// qrImage, err := core.DrawQRImage(url, qrcode.Medium, 164)
	qrImage, err := core.DrawQRImage(h.URL, qrcode.Medium, 132)
	if err != nil {
		fmt.Errorf("core.DrawQRImage err：%v", err)
	}
	// 把二维码合并到pngCarrier
	qrPoint := image.Point{X: h.X, Y: h.Y}
	core.MergeImage(c.PngCarrier, qrImage, qrImage.Bounds().Min.Sub(qrPoint))
	return
}
