package TimeTick

import (
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
	//8小时后下班
	GoHomeTime = NowTime.Add(8 * time.Hour)
}
