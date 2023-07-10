package mylog

import (
	"fmt"
	"github.com/go-toast/toast"
	"log"
	"tarot/tarots"
	"testing"
)

func TestLog(t *testing.T) {
	//file, err := tarots.TarotCards.Open("tarotCards/tarot1.jpg")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(file)
	_, err := tarots.TarotCards.ReadFile("tarotCards/tarot1.jpg")
	if err != nil {
		fmt.Println(err)
	}

}

func TestNotify(t *testing.T) {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   "下班",
		Message: "到点了该打卡下班了",
		Icon:    "C:\\path\\to\\your\\logo.png", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "打卡!", "https://www.dingtalk.com/"},
			//{"protocol", "按钮2", "https://github.com/"},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
