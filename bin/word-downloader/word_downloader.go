// program that downloads sentences to the jisho-data dir.
// configure with several options
// run with go run

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
	var outputName string="worddata3"

	var nLevel int=2
	var pageStart int=21
	var pageEnd int=38
	var sentencePagesPerWord int=4

	var workers int=10
	var pagesPerWorker int=0

	// number of ms to wait between worker collections.
	// used as rate limiter
	var collectorDelay int=4000
	// --- end config


	var here string=utils.GetHereDirRun()
	var dataPath string=filepath.Join(here,"../../data/jisho-data",outputName)

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
			CollectorDelay: collectorDelay,
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