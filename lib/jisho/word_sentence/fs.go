// word sentence fs funcs

package jisho_ws

import (
	"kj-study/lib/utils"
	"os"
	"path/filepath"
)

// info about a word sentence dict, for if want to regenerate
type WordSentenceInfo struct {
	PageStart int
	PageEnd int
	SentencesPagesPerWord int
}

// write sentences to a target folder. sentence resources come as 2 files:
// - data.gob
// - info.yml
func WriteSentences(
	dirpath string,
	sentences WordSentenceDict,
	info WordSentenceInfo,
) {
	var e error=os.MkdirAll(dirpath,0755)

	if e!=nil {
		panic(e)
	}

	e=utils.WriteGob(filepath.Join(dirpath,"data.gob"),&sentences)

	if e!=nil {
		panic(e)
	}

	e=utils.WriteYaml(filepath.Join(dirpath,"info.yml"),info)

	if e!=nil {
		panic(e)
	}
}

// read a sentence data dir
func ReadSentences(
	dirpath string,
) WordSentenceDict {
	var data WordSentenceDict=make(WordSentenceDict)
	var e error

	data,e=utils.ReadGob[WordSentenceDict](
		filepath.Join(dirpath,"data.gob"),
	)

	if e!=nil {
		panic(e)
	}

	return data
}

// save array of split dicts into a folder full of numbered gob files
func saveSplitDicts(dirpath string,sentenceDicts []WordSentenceDict) {

}