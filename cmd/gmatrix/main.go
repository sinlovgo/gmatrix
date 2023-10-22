//go:build !test

package main

import (
	"github.com/gookit/color"
	"github.com/sinlovgo/gmatrix"
	"github.com/sinlovgo/gmatrix/cmd/cli"
	"github.com/sinlovgo/gmatrix/internal/pkgJson"
	os "os"
)

const (
// exitCodeCmdArgs = 2
)

func main() {
	pkgJson.InitPkgJsonContent(gmatrix.PackageJson)

	app := cli.NewCliApp()

	args := os.Args
	//if len(args) < 2 {
	//	fmt.Printf("%s %s --help\n", color.Yellow.Render("please see help as:"), app.Name)
	//	os.Exit(exitCodeCmdArgs)
	//}
	if err := app.Run(args); nil != err {
		color.Redf("cli err at %v\n", err)
	}
}
