// general utils

package utils

import (
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// trim common whitespace from target string
func TrimWhitespace(text string) string {
	return strings.Trim(text," \n\r")
}

// set zerolog global logger default options
func ConfigureDefaultZeroLogger() {
    log.Logger=log.Output(zerolog.ConsoleWriter{
        Out:os.Stdout,
        TimeFormat: "2006/01/02 15:04:05",
    })
}

// give folder location of the exe that calls this func
func GetHereDirExe() string {
    var exePath string
    var e error
    exePath,e=os.Executable()

    if e!=nil {
        panic(e)
    }

    return filepath.Dir(exePath)
}

// when called, gives the location of the file that called this function
// only works when using go run
func GetHereDirRun() string {
    var selfFilepath string
    _,selfFilepath,_,_=runtime.Caller(1)

    return filepath.Dir(selfFilepath)
}


// shuffle an array (in place)
func ShuffleArray[T any](array []T) {
    rand.Shuffle(len(array),func (i int,j int) {
        (array)[i],(array)[j]=(array)[j],(array)[i]
    })
}

// random pick from array. does NOT mutate original array
func RandomSliceArray[T any](array []T,size int) []T {
    var arrayCopy []T=array[:]

    ShuffleArray(arrayCopy)

    return arrayCopy[0:size]
}

// get current date as a string, but with special condition. if the current time is before 8am,
// then the date becomes the previous date
func GetCurrentDateSpecial() time.Time {
    var now time.Time=time.Now()

    if now.Hour()<8 {
        now=now.Add(-24*time.Hour)
    }

    return now
}

// random pick from array, but with a daily seed. gives same thing each day
// using special current date
// func RandomSliceDaily[T any](array []T,size int) []T {
//     var now time.Time=GetCurrentDateSpecial()
//     var dailySeed int64=now.UnixNano()/int64(time.Millisecond)

//     pcg:=rand.NewPCG(uint64(dailySeed),uint64(dailySeed))
// }

// try to open web url or file with default program.
// essentially runs program like it was double clicked
func OpenTargetWithDefaultProgram(url string) error {
    var cmd *exec.Cmd=exec.Command("cmd","/c","start",url)
    var e error=cmd.Run()

    if e!=nil {
        return e
    }

    return nil
}

// remove extension from a path
func RemoveFileExtension(path string) string {
    return strings.TrimSuffix(path,filepath.Ext(path))
}

// rand int between range. max included
func RandIntRange(min int,max int) int {
    return rand.IntN(max+1-min)+min
}

// deduplicate items in array. returns new array. items are considered by their key fn, which
// must return a string
func DeduplicateBy[T any](
    array []T,
    keyFunc func(item *T) string,
) []T {
    var seenItems map[string]struct{}=make(map[string]struct{})
    var deduplicatedResult []T

    var itemI int
    for itemI = range array {
        var item *T=&array[itemI]

        var itemKey string=keyFunc(item)

        // if seen the item before, skip and do nothing
        var in bool
        _,in=seenItems[itemKey]

        if in {
            continue
        }

        // otherwise, mark seen it, and add to result
        seenItems[itemKey]=struct{}{}
        deduplicatedResult=append(deduplicatedResult,*item)
    }

    return deduplicatedResult
}