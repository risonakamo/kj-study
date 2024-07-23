// check if certain words are in a target data
// run using go run

package main

import (
	"fmt"
	jisho_ws "kj-study/lib/jisho/word_sentence"
	"kj-study/lib/utils"
	"path/filepath"
)

func main() {
	// --- config
	// name of data dir in data/jisho-data
	var dataDirName string="worddata2"

	// list of words to check for
	var checkWords []string=[]string{
		"カラー",
		"算盤",
		"仲直り",
		"懐かしい",
		"賞金",
		"お菜",
		"正味",
		"外れる",
		"一段", // 33
		"膨大",
		"混合",
		"破片",
		"特殊",
		"問い合わせ",
		"用語",
		"外れる",
		"方針",
		"薄暗い",
	}

	// set 1
	// // name of data dir in data/jisho-data
	// var dataDirName string="worddata1"

	// // list of words to check for
	// var checkWords []string=[]string{
	// 	"けど",
	// 	"押さえる",
	// 	// "当たり前", // on page 21
	// 	"柔らかい",
	// 	"直す",
	// 	"売買",
	// 	"交流",
	// }

	// set 3
	// // name of data dir in data/jisho-data
	// var dataDirName string="worddata3"

	// // list of words to check for
	// var checkWords []string=[]string{
	// 	"一段", // 33
	// 	"加速",
	// }
	// --- end config


	var here string=utils.GetHereDirRun()
	var datafileDirPath string=filepath.Join(here,"../../data/jisho-data",dataDirName)

	var data jisho_ws.WordSentenceDict=jisho_ws.ReadSentences(datafileDirPath)

	var checkWord string
	for _,checkWord = range checkWords {
		var in bool
		_,in=data[checkWord]

		if !in {
			fmt.Println("missing:",checkWord)
		}
	}

	// for key := range data {
	// 	fmt.Println(key)
	// }

	fmt.Println("done")
}