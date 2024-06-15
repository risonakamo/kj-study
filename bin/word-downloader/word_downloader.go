// program that downloads sentences to the word-data dir.
// configure with several options

package main

import (
	"fmt"
	jisho_ws "kj-study/lib/jisho/word_sentence"
	"kj-study/lib/utils"
	"path/filepath"

	"github.com/imroc/req/v3"
)

func main() {
	utils.ConfigureDefaultZeroLogger()


	// --- config
	var outputName string="worddata1"

	var nLevel int=2
	var pageStart int=1
	var pageEnd int=20
	var sentencePagesPerWord int=3

	var workers int=30
	var pagesPerWorker int=0
	// --- end config


	var here string=utils.GetHereDirExe()
	var dataPath string=filepath.Join(here,"word-data",outputName)

	fmt.Println("getting data")
	var gotWords jisho_ws.WordSentenceDict=jisho_ws.GetWordSentences_mt(
		jisho_ws.GetWordSentencesMtOptions{
			NLevel: nLevel,
			WordPageStart: pageStart,
			WordPageEnd: pageEnd,
			SentencePageLimit: sentencePagesPerWord,
			Client: req.C(),
			PagesPerWorker: pagesPerWorker,
			Workers: workers,
		},
	)

	fmt.Println("writing data:",dataPath)
	jisho_ws.WriteSentences(
		dataPath,
		gotWords,
		jisho_ws.WordSentenceInfo{
			PageStart: pageStart,
			PageEnd: pageEnd,
			SentencesPagesPerWord: sentencePagesPerWord,
		},
	)
}