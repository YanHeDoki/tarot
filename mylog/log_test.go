package mylog

import (
	"fmt"
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
