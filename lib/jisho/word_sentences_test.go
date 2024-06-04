package jisho

import (
	"fmt"
	"kj-study/lib/utils"
	"os"
	"testing"

	"github.com/imroc/req/v3"
	"github.com/k0kubun/pp/v3"
)

func Test_getWordSentences(t *testing.T) {
    result:=getWordSentencesFromApi(2,20,3,req.C())

    if len(result)==0 {
        t.Error("result was empty")
    }

    pp.Println(result)
}

func Test_saveWordSentences(t *testing.T) {
    result:=getWordSentencesFromApi(2,3,1,req.C())

    if len(result)==0 {
        t.Error("result was empty")
    }

    e:=os.MkdirAll("test",0755)

    if e!=nil {
        panic(e)
    }

    e=utils.WriteGob("test/words.gob",&result)

    if e!=nil {
        panic(e)
    }
}

func Test_readWordSentences(t *testing.T) {
    result,e:=utils.ReadGob[WordSentenceDict]("test/words.gob")

    if e!=nil {
        panic(e)
    }

    fmt.Println("read back words:",len(result))
}