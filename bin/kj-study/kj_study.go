package main

import (
	"fmt"
	jisho_ws "kj-study/lib/jisho/word_sentence"
	"kj-study/lib/kj_study"
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
    // name of word data file inside of the data dir to use. eventually, will want to be able
    // to select this from ui
    var selectedFile string="1"

    var sentencesPerWordMin int=1
    var sentencesPerWordMax int=2
    // --- end config


    // --- more variables
    var here string=utils.GetHereDirExe()
    splitDictsDataSrc=filepath.Join(here,"data/split-data",splitDictsDataSrc)
    var sessionFile string=filepath.Join(here,"data/session.yml")


    // --- app states
    var session kj_study.KjStudySession=initialiseSession(
        sessionFile,
        splitDictsDataSrc,
        selectedFile,
        sentencesPerWordMin,
        sentencesPerWordMax,
    )


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

    // get the current session
    app.Get("/get-session",func(c fiber.Ctx) error {
        return c.JSON(session)
    })

    // update the session with change in a word sentence's state
    app.Post("/set-sentence-state",func(c fiber.Ctx) error {
        var wordSentenceUpdate kj_study.WordSentencePair
        var e error=c.Bind().JSON(&wordSentenceUpdate)

        if e!=nil {
            panic(e)
        }

        e=kj_study.SetPairState(
            &session,
            wordSentenceUpdate.Word,
            wordSentenceUpdate.Sentence,
            wordSentenceUpdate.Status,
        )

        if e!=nil {
            log.Warn().
                AnErr("error",e).
                Msg("error while setting pair state")

            c.SendStatus(fiber.StatusInternalServerError)
        }

        return c.SendStatus(fiber.StatusOK)
    })



    // --- static
    app.Static("/",filepath.Join(here,"kj-study-web/build"))




    // --- running
    var e error=utils.OpenTargetWithDefaultProgram("http://localhost:4200")

    if e!=nil {
        log.Err(e).Msg("failed to open webpage with default program")
    }

    app.Listen(":4200")
}

// get initial session object by reading file. if file didnt exist, then make new one
func initialiseSession(
    sessionFile string,
    splitDictsDataSrc string,
    selectedData string,
    sentencesPerWordMin int,
    sentencesPerWordMax int,
) kj_study.KjStudySession {
    var session kj_study.KjStudySession=kj_study.GetSession(sessionFile)

    // session was empty. create a new session and write it
    if len(session.WordSentences)==0 {
        log.Info().Msg("creating new session")
        session=kj_study.GenerateNewSession(
            splitDictsDataSrc,
            selectedData,
            sentencesPerWordMin,
            sentencesPerWordMax,
        )

        kj_study.WriteSession(sessionFile,&session)
    }

    return session
}