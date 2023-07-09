package tarot

import (
	"encoding/json"
	"io"
	"math/rand"
	"strconv"
	"tarot/mylog"
	"tarot/tarots"
)

var tarotsMap map[string]*TarotCard

type TarotCard struct {
	Id       string `json:"id"`
	Tarot    string `json:"tarot"`
	Detailed string `json:"detailed"`
}
type tarotCardList struct {
	TarotCards []TarotCard `json:"tarots"`
}

func Json_to_map() {
	file, err := tarots.TarotJson.Open("json/tarot.json")
	if err != nil {
		panic("open tarot card err:" + err.Error())
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic("Read tarot card err:" + err.Error())
	}

	ts := tarotCardList{TarotCards: make([]TarotCard, 0, 52)}
	err = json.Unmarshal(bytes, &ts)
	if err != nil {
		panic("unmarshal json err:" + err.Error())
	}
	tarotsMap = make(map[string]*TarotCard, 52)
	for _, v := range ts.TarotCards {
		v := v
		tarotsMap[v.Id] = &TarotCard{
			Id:       v.Id,
			Tarot:    v.Tarot,
			Detailed: v.Detailed,
		}
	}

}

func GetTarotCard() (*TarotCard, bool) {

	if cl, ok := mylog.CheckLog(); !ok {
		randomNum := rand.Intn(51) + 1
		id := strconv.Itoa(randomNum)
		s := "tarot" + id + ".jpg"
		return tarotsMap[s], true
	} else {
		return &TarotCard{
			Id:       cl.CardId,
			Tarot:    tarotsMap[cl.CardId].Tarot,
			Detailed: tarotsMap[cl.CardId].Detailed,
		}, false
	}
}
