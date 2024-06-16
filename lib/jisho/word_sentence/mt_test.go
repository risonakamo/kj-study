package jisho_ws

import (
	"fmt"
	"kj-study/lib/utils"
	"testing"

	"github.com/imroc/req/v3"
)

func TestMain(m *testing.M) {
    utils.ConfigureDefaultZeroLogger()
    // zerolog.SetGlobalLevel(zerolog.Disabled)

    m.Run()
}


func Test_getWordsMt(t *testing.T) {
    result:=GetWordSentences_mt(GetWordSentencesMtOptions{
        NLevel: 2,

        WordPageStart: 1,
        WordPageEnd: 5,
        SentencePageLimit: 3,

        Client: req.C(),

        PagesPerWorker: 0,
        Workers: 20,
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
    result:=GetWordSentences_mt(GetWordSentencesMtOptions{
        NLevel: 2,

        WordPageStart: 1,
        WordPageEnd: 20,
        SentencePageLimit: 3,

        Client: req.C(),

        PagesPerWorker: 0,
        Workers: 4,
    })

    fmt.Println("got",len(result),"words")
    fmt.Println("got",countSentences(result),"sentences")
}