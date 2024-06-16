package jisho_ws

import (
	"testing"

	"github.com/imroc/req/v3"
	"github.com/k0kubun/pp/v3"
)

func Test_splitDict(t *testing.T) {
    wordDict:=GetWordSentences_mt(GetWordSentencesMtOptions{
        NLevel: 2,

        WordPageStart: 1,
        WordPageEnd: 5,
        SentencePageLimit: 3,

        Client: req.C(),

        PagesPerWorker: 0,
        Workers: 20,
    })

    result:=splitDict(wordDict,10)

    pp.Println(result)
}