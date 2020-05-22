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
)

// TextHandler
type TextHandler struct {
	// 合成复用Next
	Next
	X        int
	Y        int
	Size     float64
	R        uint8
	G        uint8
	B        uint8
	Text     string
	FontPath string
}

// Do 地址逻辑
func (h *TextHandler) Do(c *Context) (err error) {
	//设置字体切片
	if h.Size == 0 {
		h.Size = 24
	}

	trueTypeFont, err := core.LoadTextType(h.FontPath)
	if err != nil {
		fmt.Errorf("core.LoadTextType err：%v", err)
	}

	dText := core.NewDrawText(c.PngCarrier)
	//设置颜色
	dText.SetColor(h.R, h.G, h.B)
	err = dText.MergeText(h.Text, h.Size, trueTypeFont, h.X, h.Y)
	if err != nil {
		fmt.Errorf("dText.MergeText err：%v", err)
	}
	return
}
