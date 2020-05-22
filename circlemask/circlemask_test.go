/**
 * @Author: entere@126.com
 * @Description:
 * @File:  circle_mask_test
 * @Version: 1.0.0
 * @Date: 2020/4/22 14:50
 */

package circlemask

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"
	"time"
)

// 生成圆形图片
func TestNewCircleMask(t *testing.T) {
	randomStr := time.Unix(time.Now().Unix(), 0).Format("20060102150405")
	file, err := os.Create(randomStr + ".png")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// 目标图片，必须是png
	imageFile, err := os.Open("../assets/background.png")

	if err != nil {
		fmt.Println(err)
	}
	defer imageFile.Close()

	srcImg, _ := png.Decode(imageFile)

	w := srcImg.Bounds().Max.X - srcImg.Bounds().Min.X
	h := srcImg.Bounds().Max.Y - srcImg.Bounds().Min.Y

	d := w
	if w > h {
		d = h
	}

	maskImg := NewCircleMask(srcImg, image.Point{X: 0, Y: 0}, d)

	png.Encode(file, maskImg)

}
