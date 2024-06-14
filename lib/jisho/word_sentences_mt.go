package jisho

import (
	"fmt"
	"maps"
	"sync"

	"github.com/imroc/req/v3"
)

// func args for get word sentences mt
type GetWordSentencesMtOptions struct {
	nLevel int

	wordPageStart int
	wordPageEnd int
	sentencePageLimit int

	client *req.Client

    // give 0 for each worker to do 1 page
    pagesPerWorker int
    workers int
}

// job for word worker
type GetWordsJob struct {
	wordPageStart int
	wordPageEnd int
}

// multithread version of get word sentences
func getWordSentences_mt(
	options GetWordSentencesMtOptions,
) WordSentenceDict {
    // word jobs submitted to be acted upon by word workers.
    // close this ch to end word workers and finish all jobs
    var wordJobsCh chan GetWordsJob=make(chan GetWordsJob)

    // word workers their results to this ch.
    // close this channel to cause result collector to submit to final dict ch
    var sentenceDictResultsCh chan WordSentenceDict=make(chan WordSentenceDict)

    // 1 time ch. collector worker submits to here once
    var finalDictCh chan WordSentenceDict=make(chan WordSentenceDict)

    // signal ch. recv item if collector worker detected an empty dict
    var foundEmptyDictSigCh chan struct{}=make(chan struct{})

    var wordWorkersWg sync.WaitGroup

    // spawn word workers
    for range options.workers {
        wordWorkersWg.Add(1)
        go wordWorker(
            options.nLevel,
            options.sentencePageLimit,
            options.client,

            wordJobsCh,
            sentenceDictResultsCh,
            &wordWorkersWg,
        )
    }

    // spawn collector
    go dictMergeWorker(
        sentenceDictResultsCh,
        foundEmptyDictSigCh,
        finalDictCh,
    )

    // main thread worker - continuously submit jobs until hit the limit, or, found empty dict
    // signal triggered.
    var currentPage int=options.wordPageStart
    var currentEndPage int=currentPage+options.pagesPerWorker
    jobSubmit:
    for {
        fmt.Println("current page:",currentPage)

        // if over the page end, done
        if currentPage>options.wordPageEnd {
            fmt.Println("page end")
            break
        }

        // check if found an empty dict. if found, end job submission
        select {
            case <-foundEmptyDictSigCh:
                fmt.Println("found empty dict")
                break jobSubmit

            default:
        }

        fmt.Println("submit job:",currentPage,currentEndPage)
        wordJobsCh<-GetWordsJob{
            wordPageStart: currentPage,
            wordPageEnd: currentEndPage,
        }
        fmt.Println("done submit job")

        currentPage=currentEndPage+1
        currentEndPage=currentPage+options.pagesPerWorker
    }


    // done submitting jobs. close the jobs ch to kill workers
    close(wordJobsCh)

    fmt.Println("all jobs submitted. waiting for workers to complete")
    // wait for worker finish jobs
    wordWorkersWg.Wait()

    // close sentence submission ch to cause collector to
    // return final result
    fmt.Println("all workers done")
    close(sentenceDictResultsCh)

    // wait for final result to come in
    return <-finalDictCh
}

// recvs word jobs, converts into word sentence dict, and submits.
// closes when word jobs ch closes
// todo: getWordSentencesFromApi could be multithreaded for sentence retrieval
func wordWorker(
    nLevel int,
    sentencePageLimit int,
    client *req.Client,

    wordJobs <-chan GetWordsJob,
    submitCh chan<- WordSentenceDict,

    wg *sync.WaitGroup,
) {
    var job GetWordsJob
    for job = range wordJobs {
        var gotWordsDict WordSentenceDict=getWordSentencesFromApi(
            nLevel,
            job.wordPageStart,
            job.wordPageEnd,
            sentencePageLimit,
            client,
        )

        fmt.Println("worker trying to submit")
        submitCh<-gotWordsDict
        fmt.Println("submitted")
    }

    fmt.Println("worker done")
    wg.Done()
}

// recvs word sentences and continuously merges into a collected dict.
// upon sentence dict ch closing, submits the collected dict into final
// submit ch.
// additionally, if it encounters an empty dict, it will send a signal on found empty sentence dict
// channel
func dictMergeWorker(
    sentenceDictsCh <-chan WordSentenceDict,
    foundEmptySentenceDict chan<- struct{},
    finalSubmitCh chan<- WordSentenceDict,
) {
    var collectedDict WordSentenceDict=make(WordSentenceDict)

    var collectedCount int=0
    var sentenceDict WordSentenceDict
    for sentenceDict = range sentenceDictsCh {
        fmt.Println("collector got words",len(sentenceDict))
        collectedCount++
        fmt.Println("collected:",collectedCount)
        if len(sentenceDict)==0 {
            fmt.Println("triggering empty dict")
            foundEmptySentenceDict<-struct{}{}
        }

        maps.Copy(collectedDict,sentenceDict)
    }

    finalSubmitCh<-collectedDict
}