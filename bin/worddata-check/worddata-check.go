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
	}
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

	fmt.Println("done")
}