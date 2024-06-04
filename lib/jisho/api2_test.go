package jisho

import (
	"testing"

	"github.com/imroc/req/v3"
	"github.com/k0kubun/pp/v3"
)

func Test_getWords2(t *testing.T) {
	client:=req.C()
	result:=getNLevelWordsMulti(2,1,10,client)

	pp.Println(result)
}

// get words over the page limit
func Test_getWords3(t *testing.T) {
	client:=req.C()
	result:=getNLevelWordsMulti(2,200,210,client)

	pp.Println(result)
}