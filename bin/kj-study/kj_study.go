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

    // --- vars
    var here string=utils.GetHereDirExe()

    var kjStudyConfig kj_study.KjStudyConfig=kj_study.ReadKjStudyConfig(
        filepath.Join(here,"data/config.yml"),
    )

    var sessionFile string=filepath.Join(here,"data/session.yml")
    var dataSrcDir string=filepath.Join(here,"data/split-data",kjStudyConfig.DataDir)


    // --- app states
    // load last session. if not exist, then this is empty
    var session kj_study.KjStudySession=kj_study.GetSession(sessionFile)


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
    // get the available datafiles
    app.Get("/get-kj-files",func(c fiber.Ctx) error {
        var kjFiles []string=jisho_ws.GetSplitDictFilesList(dataSrcDir)

        log.Info().Msgf("got files: %d",len(kjFiles))

        return c.JSON(kjFiles)
    })

    // start a new session on the target data file name
    app.Get("/start-new-session/:datafile",func(c fiber.Ctx) error {
        var targetDataFile string=c.Params("datafile")

        log.Info().Msgf("new session: %s",targetDataFile)

        session=kj_study.GenerateNewSession(
            dataSrcDir,
            targetDataFile,
            kjStudyConfig.SentenceConfig.Min,
            kjStudyConfig.SentenceConfig.Max,
        )

        kj_study.WriteSession(sessionFile,&session)

        return c.JSON(session)
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

        kj_study.WriteSession(sessionFile,&session)

        return c.SendStatus(fiber.StatusOK)
    })

    // get new words for the current session. uses the same data file.
    // if there is no session loaded, returns an empty session
    app.Get("/shuffle-session",func(c fiber.Ctx) error {
        // tried to gen new session but one has not yet been loaded yet. return the same
        // empty session
        if len(session.Datafile)==0 {
            log.Warn().Msg("tried to shuffle session, but session has no datafile")
            return c.JSON(session)
        }

        session=kj_study.GenerateNewSession(
            dataSrcDir,
            session.Datafile,
            kjStudyConfig.SentenceConfig.Min,
            kjStudyConfig.SentenceConfig.Max,
        )

        log.Info().Msgf("new shuffled sentences: %d",len(session.WordSentences))

        kj_study.WriteSession(sessionFile,&session)

        return c.JSON(session)
    })



    // --- static
    app.Static("/",filepath.Join(here,"kj-study-web/build"))




    // --- running
    var e error=utils.OpenTargetWithDefaultProgram(
        fmt.Sprintf("http://localhost:%d",kjStudyConfig.Port),
    )

    if e!=nil {
        log.Err(e).Msg("failed to open webpage with default program")
    }

    app.Listen(
        fmt.Sprintf(":%d",kjStudyConfig.Port),
    )
}