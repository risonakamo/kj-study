package jisho

import (
	"testing"

	"github.com/imroc/req/v3"
	"github.com/k0kubun/pp/v3"
)

func Test_getWords1(t *testing.T) {
	client:=req.C()
	result:=getNLevelWords(2,1,client)

	pp.Println(result)
}

func Test_getSentences(t *testing.T) {
	result:=getSentencesForWord(
		"直線",
		1,
		req.C(),
	)

	pp.Println(result)
}