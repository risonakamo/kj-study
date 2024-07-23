package jisho_ws

import (
	"maps"
	"sync"

	"github.com/fatih/color"
	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"
)

// func args for get word sentences mt
type GetWordSentencesMtOptions struct {
	NLevel int

	WordPageStart int
	WordPageEnd int
	SentencePageLimit int

	Client *req.Client

    // give 0 for each worker to do 1 page
    PagesPerWorker int
    Workers int
}

// job for word worker
type GetWordsJob struct {
	wordPageStart int
	wordPageEnd int
}

// multithread version of get word sentences
func GetWordSentences_mt(
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

    var wordWorkersWg sync.WaitGroup

    // var progressPrinter *WordSentenceMtProgress=newWordSentenceMtProgress()

    // spawn word workers
    for range options.Workers {
        wordWorkersWg.Add(1)
        go wordWorker(
            options.NLevel,
            options.SentencePageLimit,
            options.Client,

            wordJobsCh,
            sentenceDictResultsCh,
            &wordWorkersWg,
        )
    }

    // spawn collector
    go dictMergeWorker(
        // progressPrinter,
        sentenceDictResultsCh,
        finalDictCh,
    )

    // main thread worker - continuously submit jobs until hit the limit, or, found empty dict
    // signal triggered.
    var currentPage int=options.WordPageStart
    var currentEndPage int=currentPage+options.PagesPerWorker
    for {
        // if over the page end, done
        if currentPage>options.WordPageEnd {
            log.Info().Msgf(color.YellowString("reached end of page jobs"))
            break
        }

        log.Info().Msgf("getting page: %d - %d",currentPage,currentEndPage)
        // progressPrinter.addJob()
        wordJobsCh<-GetWordsJob{
            wordPageStart: currentPage,
            wordPageEnd: currentEndPage,
        }

        currentPage=currentEndPage+1
        currentEndPage=currentPage+options.PagesPerWorker
    }


    // done submitting jobs. close the jobs ch to kill workers
    close(wordJobsCh)

    log.Info().Msgf(color.GreenString("waiting for workers end"))
    // wait for worker finish jobs
    wordWorkersWg.Wait()

    // close sentence submission ch to cause collector to
    // return final result
    close(sentenceDictResultsCh)

    // progressPrinter.end()
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

        submitCh<-gotWordsDict
    }

    wg.Done()
}

// recvs word sentences and continuously merges into a collected dict.
// upon sentence dict ch closing, submits the collected dict into final
// submit ch.
func dictMergeWorker(
    // progressPrint *WordSentenceMtProgress,
    sentenceDictsCh <-chan WordSentenceDict,
    finalSubmitCh chan<- WordSentenceDict,
) {
    var collectedDict WordSentenceDict=make(WordSentenceDict)

    var collectedCount int=0
    var sentenceDict WordSentenceDict
    for sentenceDict = range sentenceDictsCh {
        collectedCount++
        log.Info().Msgf("total sentence jobs collected: %d",collectedCount)
        // progressPrint.completeJob()

        if len(sentenceDict)==0 {
            log.Info().Msgf("recved empty dict from worker")
            continue
        }

        maps.Copy(collectedDict,sentenceDict)

        log.Info().Msgf("collected words: %d",len(collectedDict))
    }

    log.Info().Msg("dict merge worker done")
    finalSubmitCh<-collectedDict
}