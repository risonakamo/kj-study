// more api functions, but with different calling interfaces to be better than
// the raw api interface

package jisho

import (
	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"
)

// get N level words words from range of pages. if get a page where the result is empty, immediately stops
func GetNLevelWordsMulti(
	nLevel int,
	pageStart int,
	pageEnd int,
	client *req.Client,
) []string {
	var collected []string

	var page int
	for page=pageStart; page<=pageEnd; page++ {
        var newWords []string=getNLevelWords(nLevel,page,client)

        // got an empty page. immediately end
        if len(newWords)==0 {
            log.Warn().Msgf("got page with no words: %d",page)
            return collected
        }

		collected=append(collected,newWords...)
	}

	return collected
}

// get sentences for a word from multiple pages. immediately end on empty page
func GetSentencesMulti(
    word string,
    pageStart int,
    pageEnd int,
    client *req.Client,
) []string {
    var collected []string

    var page int
    for page=pageStart; page<=pageEnd; page++ {
        var newSentences []string=getSentencesForWord(word,page,client)

        if len(newSentences)==0 {
            return collected
        }

        collected=append(collected,newSentences...)
    }

    return collected
}