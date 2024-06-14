// funcs dealing with the word-sentence data struct, which combines words and sentences
// into one data structure

package jisho

import "github.com/imroc/req/v3"

// word-sentence dict. contains multiple words, and each word has a list of sentences
// key: a word
// val: sentences of that word
type WordSentenceDict map[string][]string

// get words-sentences dict. give limits to how much should try to get
func getWordSentencesFromApi(
    nLevel int,

    wordPageStart int,
	wordPageEnd int,
	sentencePageLimit int,

    client *req.Client,
) WordSentenceDict {
    var wordsDict WordSentenceDict=make(WordSentenceDict)

    var words []string=getNLevelWordsMulti(nLevel,wordPageStart,wordPageEnd,client)

    var word string
    for _,word = range words {
        var sentences []string=getSentencesMulti(word,1,sentencePageLimit,client)

        wordsDict[word]=sentences
    }

    return wordsDict
}
