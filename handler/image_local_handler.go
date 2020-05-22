/**
 * @Author: entere@126.com
 * @Description:
 * @File:  image_local_handler
 * @Version: 1.0.0
 * @Date: 2020/5/22 08:51
 */

package handler

import (
	"fmt"
	"github.com/hitailang/poster/core"
	"image"
	"image/png"
	"os"
)

// ImageLocalHandler 根据本地PATH设置图片
type ImageLocalHandler struct {
	// 合成复用Next
	Next
	X    int
	Y    int
	Path string //xxx/xx.png
}

// Do 地址逻辑
func (h *ImageLocalHandler) Do(c *Context) (err error) {
	//获取背景 必须是PNG图
	imageFile, err := os.Open(h.Path)
	if err != nil {
		fmt.Errorf("os.Open err：%v", err)
	}
	srcImage, err := png.Decode(imageFile)
	if err != nil {
		fmt.Errorf("png.Decode err：%v", err)
	}
	srcPoint := image.Point{
		X: h.X,
		Y: h.Y,
	}
	core.MergeImage(c.PngCarrier, srcImage, srcImage.Bounds().Min.Sub(srcPoint))
	return
}
