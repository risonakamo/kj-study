// more api functions, but with different calling interfaces to be better than
// the raw api interface

package jisho

import (
	"fmt"

	"github.com/imroc/req/v3"
)

// get N level words words from range of pages. if get a page where the result is empty, immediately stops
func getNLevelWordsMulti(
	nLevel int,
	pageStart int,
	pageEnd int,
	client *req.Client,
) []string {
	var collected []string

	var page int
	for page=pageStart; page<=pageEnd; page++ {
        fmt.Println("getting")
        var newWords []string=getNLevelWords(nLevel,page,client)

        // got an empty page. immediately end
        if len(newWords)==0 {
            fmt.Println("empty")
            return collected
        }

		collected=append(collected,newWords...)
	}

	return collected
}

// get sentences for a word from multiple pages. immediately end on empty page
func getSentencesMulti(
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