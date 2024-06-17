// word sentence fs funcs

package jisho_ws

import (
	"fmt"
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
	return readSentenceDict(filepath.Join(dirpath,"data.gob"))
}

// write a single sentence dict to gob file. give filename with extension
func writeSentenceDict(filename string,dict WordSentenceDict) {
	var e error=utils.WriteGob(filename,&dict)

	if e!=nil {
		panic(e)
	}
}

// read single sentence dict
func readSentenceDict(filename string) WordSentenceDict {
	var data WordSentenceDict=make(WordSentenceDict)
	var e error
	data,e=utils.ReadGob[WordSentenceDict](
		filename,
	)

	if e!=nil {
		panic(e)
	}

	return data
}

// save array of split dicts into a folder full of numbered gob files
func SaveSplitDicts(dirpath string,sentenceDicts []WordSentenceDict) {
	var e error=os.MkdirAll(dirpath,0755)

	if e!=nil {
		panic(e)
	}

	// for all sentence dicts, write a file named after the index
	var i int
	var sentenceDict WordSentenceDict
	for i,sentenceDict = range sentenceDicts {
		var dictFileName string=filepath.Join(
			dirpath,
			fmt.Sprintf("%d.gob",i+1),
		)

		writeSentenceDict(dictFileName,sentenceDict)
	}
}

// read a single split dict from a split dict dir. target a dir containing multiple split dict
// files, and the index name to target
func ReadSingleSplitDict(dirpath string,index int) WordSentenceDict {
	var splitDictFileName string=filepath.Join(dirpath,fmt.Sprintf("%d.gob",index))
	return readSentenceDict(splitDictFileName)
}