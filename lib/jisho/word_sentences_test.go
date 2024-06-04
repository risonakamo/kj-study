package jisho

import (
	"testing"

	"github.com/imroc/req/v3"
	"github.com/k0kubun/pp/v3"
)

func Test_getWordSentences(t *testing.T) {
    result:=getWordSentencesFromApi(20,3,2,req.C())

    if len(result)==0 {
        t.Error("result was empty")
    }

    pp.Println(result)
}