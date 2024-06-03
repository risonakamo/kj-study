// jisho api accessing functions

package jisho

import (
	"kj-study/lib/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"
)

// get words of some N level from a target page.
// lower N levels include all words of the higher levels.
func getNLevelWords(
	nLevel int,
	page int,
	client *req.Client,
) []string {
	var resp *req.Response
	var e error
	resp,e=client.R().
		SetPathParam("page",strconv.Itoa(page)).
		SetPathParam("nlevel",strconv.Itoa(nLevel)).
		Get("https://jisho.org/search/%23words%20%23jlpt-n{nlevel}?page={page}")

	if e!=nil {
		panic(e)
	}

	var doc *goquery.Document
	doc,e=goquery.NewDocumentFromReader(resp.Body)

	if e!=nil {
		panic(e)
	}

	var collectedWords []string
	// target words and add them to collection. trim the words.
	doc.Find("span.text").Each(func(i int,element *goquery.Selection) {
		var trimmed string=utils.TrimWhitespace(element.Text())

		if trimmed=="Words" {
			return
		}

		collectedWords=append(collectedWords,trimmed)
	})

	return collectedWords
}

// get N level words words from range of pages
func getNLevelWordsMulti(
	nLevel int,
	pageStart int,
	pageEnd int,
	client *req.Client,
) []string {
	var collected []string

	var page int
	for page=pageStart; page<=pageEnd; page++ {
		collected=append(collected,getNLevelWords(nLevel,page,client)...)
	}

	return collected
}

// get sentences for a word on certain page
func getSentencesForWord(
	word string,
	page int,
	client *req.Client,
) []string {
	var resp *req.Response
	var e error
	resp,e=client.R().
		SetPathParam("word",word).
		SetPathParam("page",strconv.Itoa(page)).
		Get("https://jisho.org/search/{word}%20%23sentences?page={page}")

	if e!=nil {
		panic(e)
	}

	var doc *goquery.Document
	doc,e=goquery.NewDocumentFromReader(resp.Body)

	if e!=nil {
		panic(e)
	}

	var collected []string

	// find all sentence containers
	doc.Find("ul.japanese_sentence").Each(func(_ int,element *goquery.Selection) {
		var sentenceParts []string

		// sentence container is bunch of text nodes intermixed with li nodes. each li node has
		// 2 nodes in it - hiragana and kanji. take the kanji node.
		element.Contents().Each(func(_ int,childElement *goquery.Selection) {
			// li element. handle specially
			if childElement.Is("li") {
				var liSize int=childElement.Children().Length()

				switch liSize {
					// pair of hiragana and not hiragana. take the 2nd
					case 2:
					sentenceParts=append(
						sentenceParts,
						utils.TrimWhitespace(childElement.Children().Eq(1).Text()),
					)

					// just 1 hiragana. take the 1st
					case 1:
					sentenceParts=append(
						sentenceParts,
						utils.TrimWhitespace(childElement.Children().First().Text()),
					)

					default:
					log.Warn().Msgf("strange li size: %d",liSize)
				}

			// text. just take the text
			} else {
				sentenceParts=append(sentenceParts,utils.TrimWhitespace(childElement.Text()))
			}
		})

		collected=append(collected,strings.Join(sentenceParts,""))
	})

	return collected
}