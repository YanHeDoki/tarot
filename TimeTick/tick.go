package TimeTick

import (
	"fmt"
	"time"
)

var (
	NowTime    time.Time
	GoHomeTime time.Time
)

// 开始下班计时
func TimeStart() {
	//获取当前时间
	NowTime = time.Now()
	//format := time.Now().Format(time.DateOnly)
	//format += " 09:00:00"
	//NowTime, _ = time.Parse(time.DateTime, format)
	//8小时后下班
	GoHomeTime = NowTime.Add(9 * time.Hour)
	fmt.Println(NowTime, GoHomeTime)
}
