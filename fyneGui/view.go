package fyneGui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-toast/toast"
	"tarot/TimeTick"
	"tarot/mylog"
	"tarot/tarot"
	"tarot/tarots"
	"time"
)

var once = 0
var tarotfile string

func GuiStart() {
	myapp := app.New()
	myapp.Settings().SetTheme(theme.LightTheme())
	w := myapp.NewWindow("tarot")
	w.Resize(fyne.NewSize(250, 200))
	f := binding.NewFloat()
	f.Set(1.0)

	progress := widget.NewProgressBarWithData(f)

	button := widget.NewButton("tarot", func() {
		w2 := myapp.NewWindow("tarot")
		if card, ok := tarot.GetTarotCard(); ok {
			tarotfile = "tarotCards/" + card.Id
			mylog.AppendLog(time.Now(), card.Id)
		} else {
			tarotfile = "tarotCards/" + card.Id
		}
		file, err := tarots.TarotCards.Open(tarotfile)
		if err != nil {
			fmt.Println(err)
		}
		//canvas.NewImageFromResource(fyne.NewStaticResource())
		img := canvas.NewImageFromReader(file, "tarot1.jpg")
		img.FillMode = canvas.ImageFillOriginal
		w2.SetContent(img)
		w2.Show()
	})

	button2 := widget.NewButton("start", func() {
		if once == 0 {
			TimeTick.TimeStart()
			//开始计时
			go start(f)
			once = 1
		}
	})
	w.SetMaster()
	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayoutWithRows(3), button2, progress, button))
	w.ShowAndRun()
}

func start(f binding.Float) {
	for range time.Tick(1 * time.Minute) {
		now := time.Now().Unix()
		since := TimeTick.GoHomeTime.Unix() - now
		num := float64(since) / float64(TimeTick.GoHomeTime.Unix()-TimeTick.NowTime.Unix())
		f.Set(num)
		if num < 0.00 {
			Notify()
			once = 0
			break

		}
	}

}

func Notify() {
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
	notification.Push()
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
