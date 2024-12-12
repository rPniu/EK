package main

import (
	"EK/id"
	"EK/image"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("DAN")

	EKLabel := widget.NewLabel("Ek")
	IDLabel := widget.NewLabel("ID:")

	//输入框，获取元素
	EntryToGetData := widget.NewEntry()
	EntryToGetData.SetPlaceHolder("Enter URL...")

	GetDataContainer := container.NewVBox(
		EKLabel,
		EntryToGetData,
	)

	//主容器加载获取数据容器
	mainContainer := container.NewVBox(GetDataContainer)

	var qrContainer *fyne.Container

	GenerateQRCode := func() {

		//获取id生成二维码
		ID := id.ExtractIDFromURL(EntryToGetData.Text)

		if id.CheckID(ID) {
			IDLabel.SetText("ID is empty or invalid, please enter a valid URL.") // 更新 IDLabel 提示信息
			mainContainer.Refresh()
			return
		}

		IDLabel.SetText(ID)

		QdTime, QtTime, QdImage, QtImage := image.GenerateQRCodeImage(ID)

		QdContainer := container.NewVBox(widget.NewLabel("QD\n"+QdTime), QdImage)
		QtContainer := container.NewVBox(widget.NewLabel("QT\n"+QtTime), QtImage)

		//二维码容器
		ImageContainer := container.NewHBox(layout.NewSpacer(), QdContainer, layout.NewSpacer(), QtContainer, layout.NewSpacer())

		// 如果已有二维码容器，移除它
		if qrContainer != nil {
			mainContainer.Remove(qrContainer)
		}

		// 更新二维码容器变量，并添加到主容器
		qrContainer = ImageContainer
		mainContainer.Add(qrContainer)
		mainContainer.Refresh()
	}

	//确认按钮，展示id和二维码
	ConfirmButton := widget.NewButtonWithIcon("start", theme.ConfirmIcon(), GenerateQRCode)

	// 绑定回车键触发确认按钮
	EntryToGetData.OnSubmitted = func(text string) {
		GenerateQRCode()
	}

	//主容器加载确认按钮
	mainContainer.Add(container.NewHBox(layout.NewSpacer(), ConfirmButton))
	mainContainer.Add(container.NewHBox(layout.NewSpacer(), IDLabel, layout.NewSpacer()))
	mainContainer.Refresh()

	w.SetContent(mainContainer)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
