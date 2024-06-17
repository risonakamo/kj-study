package main

import (
	"fmt"
	jisho_ws "kj-study/lib/jisho/word_sentence"
	"kj-study/lib/utils"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
)

func main() {
    utils.ConfigureDefaultZeroLogger()

    // --- config
    // name of word data folder to use. must be present inside data/split-data
    var splitDictsDataSrc string="worddata1"
    // --- end config


    // --- more variables
    var here string=utils.GetHereDirExe()
    splitDictsDataSrc=filepath.Join(here,"data/split-data",splitDictsDataSrc)



    // --- fiber setup
    var app *fiber.App=fiber.New(fiber.Config{
        CaseSensitive: true,
        ErrorHandler: func(c fiber.Ctx, err error) error {
            fmt.Println("fiber error")
            fmt.Println(err)
            return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
        },
    })


    // --- apis
    // get the available kj filenames
    app.Get("/get-kj-files",func(c fiber.Ctx) error {
        var kjFiles []string=jisho_ws.GetSplitDictFilesList(splitDictsDataSrc)

        return c.JSON(kjFiles)
    })

    // get a target kj file. file type is a word sentence dict
    app.Get("/get-kj-file/:filename",func(c fiber.Ctx) error {
        var filename string=c.Params("filename")

        var response jisho_ws.WordSentenceDict=make(jisho_ws.WordSentenceDict)

        if len(filename)==0 {
            log.Warn().Msg("tried to get empty filename")
            return c.JSON(response)
        }

        response=jisho_ws.ReadSingleSplitDict(splitDictsDataSrc,filename)

        return c.JSON(response)
    })


    // --- static
    app.Static("/",filepath.Join(here,"kj-study-web/build"))




    // utils.OpenTargetWithDefaultProgram("http://localhost:4200")
    app.Listen(":4200")
}