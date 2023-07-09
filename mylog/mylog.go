package mylog

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type CardLog struct {
	CardTime string `json:"cardTime"`
	CardId   string `json:"cardId"`
}

func CheckLog() (*CardLog, bool) {
	line := getLastLine2()
	if len(line) <= 0 {
		return nil, false
	}
	cl := new(CardLog)
	json.Unmarshal(line, &cl)
	format := time.Now().Format(time.DateOnly)
	format += " 04:00:00"
	parseNow, err := time.Parse(time.DateTime, format)
	if err != nil {
		fmt.Println("parse Now err:", err)
	}
	parseLog, err := time.Parse(time.DateTime, cl.CardTime)
	if err != nil {
		fmt.Println("parse Log err:", err)
	}
	if parseLog.Before(parseNow) {
		return nil, false
	}
	return cl, true
}

func AppendLog(Logtime time.Time, cardId string) {
	cl := CardLog{
		CardTime: Logtime.Format(time.DateTime),
		CardId:   cardId,
	}
	bytes, _ := json.Marshal(&cl)
	s := string(bytes) + "\n"
	file, err := os.OpenFile("checkLog.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(s)
	if err != nil {
		fmt.Println(err)
	}
}

func getLastLine2() []byte {
	file, err := os.OpenFile("checkLog.txt", os.O_CREATE|os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	bytes := []byte{}
	for scanner.Scan() {
		bytes = scanner.Bytes()
	}
	fmt.Println(string(bytes))
	return bytes
}

func getLastLine() []byte {
	file, err := os.OpenFile("checkLog.txt", os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	rd := bufio.NewReader(file)
	if info.Size() > 0 {
		index := int64(-1)
		for {
			index--
			file.Seek(index, io.SeekEnd)
			readByte, err := rd.ReadByte()
			if readByte == '\n' {
				file.Seek(0, io.SeekEnd)
				break
			}
			if err != nil {
				if err == io.EOF {
					//file.Seek(0, io.SeekEnd)
					break
				}
				fmt.Println(err)
			}
		}
		//fmt.Println(index, info.Size())
		line, _, err := rd.ReadLine()
		if err != nil {
			fmt.Println("get last line err:", err)
		}
		return line
	}
	return []byte{}
}

func getLastErr() {
	file, err := os.OpenFile("checkLog.txt", os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	if info.Size() > 0 {
		index := int64(-1)
		r := bufio.NewReader(file)
		for {
			index--
			file.Seek(index, io.SeekEnd)
			readByte, err := r.ReadByte()
			if readByte == '\n' {
				//file.Seek(0, io.SeekEnd)
				break
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
			}
		}

		line, _, _ := r.ReadLine()
		fmt.Println(string(line))
	}

}

func getLastOk() []byte {
	file, err := os.OpenFile("checkLog.txt", os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	info, _ := file.Stat()
	if info.Size() > 0 {
		index := int64(-1)
		b := []byte{0}
		for {
			index--
			file.Seek(index, io.SeekEnd)
			_, err := file.Read(b)
			if b[0] == '\n' {
				break
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
			}
		}
		rd := bufio.NewReader(file)
		line, _, _ := rd.ReadLine()
		return line
	}
	return []byte{}

}
