package test

import (
	"fmt"
	jisho_ws "kj-study/lib/jisho/word_sentence"
	"testing"

	"github.com/k0kubun/pp/v3"
)

// try to read a single split data file. only works after have run the data-splitter.exe
func Test_readSingle(t *testing.T) {
    data:=jisho_ws.ReadSingleSplitDict("../data/split-data/worddata1","2")
    pp.Print(data)

    fmt.Println("words:",len(data))
}

func Test_getDataList(t *testing.T) {
	result:=jisho_ws.GetSplitDictFilesList("../data/split-data/worddata1")
	pp.Print(result)
}

func Test_sentenceSubset(t *testing.T) {
	data:=jisho_ws.ReadSingleSplitDict("../data/split-data/worddata1","2")

	subset:=jisho_ws.GetSentenceSubset(
		data,
		1,
		3,
	)

	pp.Println(subset)
}