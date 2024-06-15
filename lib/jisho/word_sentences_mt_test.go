package jisho

import (
	"fmt"
	"kj-study/lib/utils"
	"testing"

	"github.com/imroc/req/v3"
)

func TestMain(m *testing.M) {
    utils.ConfigureDefaultZeroLogger()

    m.Run()
}


func Test_getWordsMt(t *testing.T) {
    result:=getWordSentences_mt(GetWordSentencesMtOptions{
        nLevel: 2,

        wordPageStart: 1,
        wordPageEnd: 5,
        sentencePageLimit: 3,

        client: req.C(),

        pagesPerWorker: 0,
        workers: 20,
    })

    fmt.Println("got",len(result),"words")

    // var words []string=maps.Keys(result)
    // pp.ColoringEnabled=false
    // pp.Println(words)

    var expectedWords []string=[]string{
        "看板", // page 5
        "牧場", // page 5
        "中世", // page 4
        "長い", // page 2
        // "家屋", // page 6
    }

    var expectedWord string
    for _,expectedWord = range expectedWords {
        var in bool
        _,in=result[expectedWord]

        if !in {
            t.Errorf("missing expected word: %s",expectedWord)
        }
    }
}

// bigger mt test
func Test_getWordsMt2(t *testing.T) {
    result:=getWordSentences_mt(GetWordSentencesMtOptions{
        nLevel: 2,

        wordPageStart: 1,
        wordPageEnd: 20,
        sentencePageLimit: 3,

        client: req.C(),

        pagesPerWorker: 0,
        workers: 4,
    })

    fmt.Println("got",len(result),"words")
    fmt.Println("got",countSentences(result),"sentences")
}