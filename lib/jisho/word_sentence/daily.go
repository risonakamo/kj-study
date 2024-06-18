// provides functions to access a "daily word set"
// note: for now, pivoting program direction, so this is not needed yet.

package jisho_ws

import (
	"kj-study/lib/utils"
	"math"
)

// from a target sentence data dir, return a subset sentence data dict
// func getDailySet(
//     sentenceDataDir string,
// ) WordSentenceDict {
//     var allSentences WordSentenceDict=ReadSentences(sentenceDataDir)
// }

// given a sentence dict, return the sentence dict, but for each word, the number of sentences
// for that word is reduced to a random number between a given range
func GetSentenceSubset(
    wordData WordSentenceDict,
    sentencePerWordMin int,
    sentencePerWordMax int,
) WordSentenceDict {
    var newDict WordSentenceDict=make(WordSentenceDict)

    var word string
    var sentences []string
    for word,sentences = range wordData {
        var randomSize int=utils.RandIntRange(
            int(math.Min(float64(sentencePerWordMin),float64(len(sentences)))),
            int(math.Min(float64(sentencePerWordMax),float64(len(sentences)))),
        )

        var chosenWords []string=utils.RandomSliceArray(sentences,randomSize)

        newDict[word]=chosenWords
    }

    return newDict
}