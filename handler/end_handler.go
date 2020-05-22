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

// EndHandler 结束，写在最后，把图片合并到一张图上
type EndHandler struct {
	// 合成复用Next
	Next
	Output string // "/tmp/xxx.png"
}

// Do 地址逻辑
func (h *EndHandler) Do(c *Context) (err error) {
	// 新建文件载体
	//fileName := "poster-" + xid.New().String() + ".png"
	merged, err := core.NewMerged(h.Output)
	if err != nil {
		fmt.Errorf("core.NewMerged err：%v", err)
	}
	// 合并
	err = core.Merge(c.PngCarrier, merged)
	if err != nil {
		fmt.Errorf("core.Merge err：%v", err)
	}
	return
}
