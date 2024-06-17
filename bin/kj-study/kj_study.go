package main

import (
	"fmt"
	"kj-study/lib/utils"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
)

func main() {
    utils.ConfigureDefaultZeroLogger()

    var here string=utils.GetHereDirExe()

    var app *fiber.App=fiber.New(fiber.Config{
        CaseSensitive: true,
        ErrorHandler: func(c fiber.Ctx, err error) error {
            fmt.Println("fiber error")
            fmt.Println(err)
            return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
        },
    })



    // --- static
    app.Static("/",filepath.Join(here,"kj-study-web/build"))




    // utils.OpenTargetWithDefaultProgram("http://localhost:4200")
    app.Listen(":4200")
}