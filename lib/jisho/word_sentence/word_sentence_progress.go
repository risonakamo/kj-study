// progress printer for mt

package jisho_ws

import (
	"fmt"

	"github.com/gosuri/uilive"
)

// class to print out progress
type WordSentenceMtProgress struct {
    currentJobs int
    collected int

	writer *uilive.Writer
}

// new sentence progress. while exists, prints are captured by this writer
func newWordSentenceMtProgress() *WordSentenceMtProgress {
    var progress WordSentenceMtProgress=WordSentenceMtProgress{
        writer: uilive.New(),
    }

    progress.writer.Start()

    return &progress
}

// end progress writer.
func (progress *WordSentenceMtProgress) end() {
    progress.writer.Stop()
}

// register a job began
func (progress *WordSentenceMtProgress) addJob() {
    progress.currentJobs+=1
    progress.render()
}

// register a job completed. adds to collected and removes job
func (progress *WordSentenceMtProgress) completeJob() {
    progress.currentJobs-=1
    progress.collected+=1
    progress.render()
}

// render progress
func (progress *WordSentenceMtProgress) render() {
    fmt.Fprintf(
        progress.writer,
        "current jobs: %d\ncollected: %d\n",
        progress.currentJobs,
        progress.collected,
    )
}