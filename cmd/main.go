package main

import (
	"os"

	"github.com/xpzouying/go-clean-code-template/internal/cli"
)

func main() {

	app := cli.CreateCliApp()
	if err := app.Run(os.Args); err != nil {
		println("Run app failed!", err.Error())
	}

}
