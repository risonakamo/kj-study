package jisho

import (
	"fmt"
	"testing"

	"github.com/imroc/req/v3"
)

func Test_getWordsMt(t *testing.T) {
    result:=getWordSentences_mt(
        2,

        1,
        20,
        3,

        req.C(),

        3,
        10,
    )

    fmt.Println("got",len(result),"words")
}