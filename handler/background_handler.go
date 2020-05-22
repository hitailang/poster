/**
 * @Author: entere@126.com
 * @Description:
 * @File:  background_handler.go
 * @Version: 1.0.0
 * @Date: 2020/5/21 12:31
 */

package handler

import (
    "fmt"
    "github.com/hitailang/poster/core"
    "image"
    "image/png"
    "os"
)



// BackgroundHandler
type BackgroundHandler struct {
    // 合成复用Next
    Next
    X int
    Y int
    Path string
}

// Do 地址逻辑
func (h *BackgroundHandler) Do(c *Context) (err error) {
    //获取背景 必须是PNG图
    bgFile, err := os.Open(h.Path)
    if err != nil {
        fmt.Errorf("os.Open err：%s", err)
    }
    bgImage, err := png.Decode(bgFile)
    if err != nil {
        fmt.Errorf("png.Decode err：%v", err)
    }
    bgPoint := image.Point{
        X: h.X,
        Y: h.Y,
    }
    core.MergeImage(c.PngCarrier, bgImage, bgImage.Bounds().Min.Sub(bgPoint))
    return
}