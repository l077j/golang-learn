package main

import (
	"github.com/fogleman/gg"
)

func main() {
	Dot()
	Line()
	dc := gg.NewContext(1000, 1000) // 创建画布 高1000 宽1000
	dc.DrawCircle(500, 500, 400)    // 在(500,500)坐标位置绘制一个半径400的圆
	dc.SetRGB(0, 0, 0)              // 设置颜色黑色
	dc.Fill()                       // 填充
	dc.SavePNG("out.png")           // 保存到图片文件
}

// 描点
func Dot() {
	dc := gg.NewContext(1000, 1000)
	dc.SetRGB(0, 0, 0)
	for i := 1; i < 10; i++ {
		dc.DrawPoint(float64(50*i), float64(50*i), 5) // 设置点的坐标和半径
	}
	dc.Fill() // 填充
	dc.SavePNG("dots.png")
}

// 画线
func Line() {
	dc := gg.NewContext(1000, 1000)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(5) // 设置线宽
	dc.DrawLine(1000, 0, 0, 1000)
	dc.DrawLine(1000, 1000, 0, 0)
	dc.Stroke() // 连线
	dc.SavePNG("lines.png")
}
