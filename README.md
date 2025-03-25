## EK-GO-fyne停用

### 环境

测试环境：

go version go1.23.2 windows/amd64

fyne：fyne.io/fyne/v2 v2.5.2

### 使用

源码+打包exe文件

-   可直接食用exe文件
-   其他平台自行打包

使用`go mod tidy` 会自动下载缺失的依赖库

```
go mod tity
```

打包

icon为图标，可省略

```
fyne package -os darwin -icon myapp.png
fyne package -os linux -icon myapp.png
fyne package -os windows -icon myapp.png
fyne package -os android -appID com.example.myapp -icon mobileIcon.png
fyne package -os ios -appID com.example.myapp -icon mobileIcon.png
```

安装

```
android:
adb install myapp.apk

IOS:
xcrun simctl install booted myapp.app
```

大菜鸡第一次做，希望大佬多有包含。谢谢体谅
