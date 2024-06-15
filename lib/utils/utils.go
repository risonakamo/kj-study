// general utils

package utils

import (
	"os"
	"path/filepath"
	"strings"

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