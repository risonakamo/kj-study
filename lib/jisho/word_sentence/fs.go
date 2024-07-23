// word sentence fs funcs

package jisho_ws

import (
	"errors"
	"fmt"
	"io/fs"
	"kj-study/lib/utils"
	"maps"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
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

// read single sentence dict. performs deduplication
func readSentenceDict(filename string) WordSentenceDict {
	var data WordSentenceDict=make(WordSentenceDict)
	var e error
	data,e=utils.ReadGob[WordSentenceDict](
		filename,
	)

	if e!=nil {
		if errors.Is(e,fs.ErrNotExist) {
			log.Warn().Msgf("failed to read sentence file: %s",filename)
			return data
		}

		panic(e)
	}

	data=deduplicateWordSentences(data)

	return data
}

// save array of split dicts into a folder full of numbered gob files
func SaveSplitDicts(
	dirpath string,
	sentenceDicts []WordSentenceDict,
	initialNumber int,
) {
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
			fmt.Sprintf("%d.gob",i+initialNumber),
		)

		writeSentenceDict(dictFileName,sentenceDict)
	}
}

// read a single split dict from a split dict dir. target a dir containing multiple split dict
// files, and the index name to target
func ReadSingleSplitDict(dirpath string,splitDictName string) WordSentenceDict {
	var splitDictFileName string=filepath.Join(dirpath,fmt.Sprintf("%s.gob",splitDictName))
	return readSentenceDict(splitDictFileName)
}

// for a target split dict folder, get the available file names in the folder.
// filenames will not have file extension. expects single level.
func GetSplitDictFilesList(dirpath string) []string {
	var files []fs.DirEntry
	var e error
	files,e=os.ReadDir(dirpath)

	if e!=nil {
		panic(e)
	}

	var collectedFileNames []string

	var file fs.DirEntry
	for _,file = range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name())!=".gob" {
			continue
		}

		collectedFileNames=append(collectedFileNames,utils.RemoveFileExtension(file.Name()))
	}

	return collectedFileNames
}

// read all split dict gobs in a target split dict folder. combines all into a single
// word sentence dict
func ReadAllSplitDicts(dirpath string) WordSentenceDict {
	var datafiles []string=GetSplitDictFilesList(dirpath)

	var collectedDict WordSentenceDict=make(WordSentenceDict)

	var datafileName string
	for _,datafileName = range datafiles {
		var singleData WordSentenceDict=ReadSingleSplitDict(dirpath,datafileName)

		maps.Copy(collectedDict,singleData)
	}

	return collectedDict
}