package timedan

import (
	"fmt"
	"math/rand"
	"time"
)

func GetTime() (int64, int64) {
	now := time.Now()

	// 设置为中国时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return 0, 0
	}
	chinaTime := now.In(location)

	// 将时间转换为时间戳
	QDTime := chinaTime.Unix() // 秒级时间戳

	rand.Seed(time.Now().UnixNano())

	// 生成 20 到 60 的随机数
	randomNumber := rand.Intn(41) + 20

	return QDTime, QDTime + int64(randomNumber)
}

func UnixToTime(unix int64) string {
	return time.Unix(unix, 0).Format("2006-01-02 15:04:05")
}
