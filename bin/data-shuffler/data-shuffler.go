// data shuffler program. targets a split data folder (folder full of gob data)
// and generates another split data dir using all data from the initial split data dir.
// the number of words per split data can be adjusted.
//
// usage:
// data-shuffler.exe <input split dir> <output split dir>

package main

import (
	"fmt"
	jisho_ws "kj-study/lib/jisho/word_sentence"
)

func main() {
    // --- config
    var targetSplitdataDir string="data/split-data/review-set1"
    var outputDataDir string="data/split-data/review-set1-shuffled"
    var newWordsPerSplit int=10
    // --- end config


    fmt.Println("reading split datas")
    var collectedSplitData jisho_ws.WordSentenceDict=jisho_ws.ReadAllSplitDicts(targetSplitdataDir)

    fmt.Println("resplitting")
    var splittedDicts []jisho_ws.WordSentenceDict=jisho_ws.SplitDict(
        collectedSplitData,
        newWordsPerSplit,
        true,
    )
    fmt.Println("split into",len(splittedDicts),"dicts")
    fmt.Println()

    fmt.Println("saving to:",outputDataDir)
    jisho_ws.SaveSplitDicts(outputDataDir,splittedDicts)
    fmt.Println()

    fmt.Println("complete")
}