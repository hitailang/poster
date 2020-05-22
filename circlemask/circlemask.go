/**
 * @Author: entere@126.com
 * @Description: 圆型图片遮罩
 * @File:  circle_mask
 * @Version: 1.0.0
 * @Date: 2020/4/22 14:50
 */

package circlemask

import (
	"image"
	"image/color"
	"math"
)

func NewCircleMask(img image.Image, p image.Point, d int) CircleMask {
	return CircleMask{img, p, d}
}

type CircleMask struct {
	image    image.Image
	point    image.Point
	diameter int
}

func (ci CircleMask) ColorModel() color.Model {
	return ci.image.ColorModel()
}

func (ci CircleMask) Bounds() image.Rectangle {
	return image.Rect(0, 0, ci.diameter, ci.diameter)
}

func (ci CircleMask) At(x, y int) color.Color {
	d := ci.diameter
	dis := math.Sqrt(math.Pow(float64(x-d/2), 2) + math.Pow(float64(y-d/2), 2))
	if dis > float64(d)/2 {
		return ci.image.ColorModel().Convert(color.RGBA{255, 255, 255, 0})
	} else {
		return ci.image.At(ci.point.X+x, ci.point.Y+y)
	}
}
