package image

import (
	"EK/timedan"
	"bytes"
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"github.com/skip2/go-qrcode"
	"log"
)

func generateImage(content string) *canvas.Image {
	// 生成二维码字节流
	qrCode, err := qrcode.Encode(content, 3, 256)
	if err != nil {
		log.Fatalf("Failed to generate QR code: %v", err)
	}
	// 将字节流包装为 Reader，供 Fyne 使用
	reader := bytes.NewReader(qrCode)

	img := canvas.NewImageFromReader(reader, "qrcode.png")
	return img
}

// 将字节流包装为 Reader，供 Fyne 使用

func GenerateQRCodeImage(id string) (string, string, *canvas.Image, *canvas.Image) {
	//获取时间戳
	QdTimeUnix, QtTimeUnix := timedan.GetTime()

	//生成QD和QT信息
	QdInfo := fmt.Sprintf("%s|%d|QD", id, QdTimeUnix)
	QtInfo := fmt.Sprintf("%s|%d|QT", id, QtTimeUnix)

	// 创建 Fyne 图片组件
	QdImage := generateImage(QdInfo)
	QtImage := generateImage(QtInfo)

	QdImage.FillMode = canvas.ImageFillOriginal
	QtImage.FillMode = canvas.ImageFillOriginal

	QdTime := timedan.UnixToTime(QdTimeUnix)
	QtTime := timedan.UnixToTime(QtTimeUnix)

	return QdTime, QtTime, QdImage, QtImage
}
