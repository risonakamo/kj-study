// reads several target jisho data dirs, combines them all, and splits them into
// pieces. the data can be shuffled before splitting.
// check config section for options
// run with go run

package main

import (
	"fmt"
	jisho_ws "kj-study/lib/jisho/word_sentence"
	"kj-study/lib/utils"
	"path/filepath"
)

func main() {
    // --- config
    // name of a jisho data dir present in jisho-data
    var targetData string="worddata1"

    var wordsPerSplit int=10
    // --- end config

    var here string=utils.GetHereDirExe()
    var jishoDataDir string=filepath.Join(here,"data/jisho-data",targetData)
    var splitDataPath=filepath.Join(here,"data/split-data",targetData)


    var fullSentencesDict jisho_ws.WordSentenceDict=jisho_ws.ReadSentences(jishoDataDir)
    fmt.Println("read sentence dict:",jishoDataDir)
    fmt.Println("words:",len(fullSentencesDict))
    fmt.Println("sentences:",jisho_ws.CountSentences(fullSentencesDict))
    fmt.Println()

    fmt.Println("doing split...")
    var splittedDicts []jisho_ws.WordSentenceDict=jisho_ws.SplitDict(
        fullSentencesDict,
        wordsPerSplit,
        true,
    )
    fmt.Println("split into",len(splittedDicts),"dicts")
    fmt.Println()

    fmt.Println("saving to:",splitDataPath)
    jisho_ws.SaveSplitDicts(splitDataPath,splittedDicts)
    fmt.Println()

    fmt.Println("complete")
}