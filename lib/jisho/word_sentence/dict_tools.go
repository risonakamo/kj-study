// funcs interacting with the word sentence dict data type

package jisho_ws

import (
	"kj-study/lib/utils"

	"golang.org/x/exp/maps"
)

// count number of sentences in word sentence dict
func CountSentences(wordsDict WordSentenceDict) int {
    var count int=0

    var sentences []string
    for _,sentences = range wordsDict {
        count+=len(sentences)
    }

    return count
}

// split a sentence dict into multiple dicts, with randomisation.
// size of each smaller dict set by split dict size.
// give shuffle option to shuffle all words before splitting
func SplitDict(
    wordsDict WordSentenceDict,
    splitDictSize int,
    shuffle bool,
) []WordSentenceDict {
    if splitDictSize==0 {
        panic("bad split dict size")
    }

    var allWords []string=maps.Keys(wordsDict)

    if shuffle {
        utils.ShuffleArray[string](allWords)
    }

    var collectedDicts []WordSentenceDict
    var newMiniDict WordSentenceDict=make(WordSentenceDict)
    var addedWords int=0

    var word string
    for _,word = range allWords {
        // added the target number of words to the split dict. add the split
        // dict to the collection. start new split dict
        if addedWords>=splitDictSize {
            addedWords=0
            collectedDicts=append(collectedDicts,newMiniDict)
            newMiniDict=make(WordSentenceDict)
        }

        newMiniDict[word]=wordsDict[word]
        addedWords++
    }

    // for the last split dict, if it has something, add it to the collection
    if len(newMiniDict)>0 {
        collectedDicts=append(collectedDicts,newMiniDict)
    }

    return collectedDicts
}