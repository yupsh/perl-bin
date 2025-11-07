package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/perl"
)

const (
	flagScript      = "execute"
	flagModule      = "module"
	flagLibPath     = "include"
	flagEncoding    = "encoding"
	flagInPlace     = "in-place"
	flagPrint       = "print"
	flagLoop        = "loop"
	flagAutoSplit   = "autosplit"
	flagCheckSyntax = "check"
	flagWarnings    = "warnings"
	flagStrict      = "strict"
	flagDebug       = "debug"
	flagTaint       = "taint"
)

func main() {
	app := &cli.App{
		Name:  "perl",
		Usage: "perl command wrapper for yupsh",
		UsageText: `perl [OPTIONS] [FILE...]

   Execute Perl scripts.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagScript,
				Aliases: []string{"e"},
				Usage:   "execute script",
			},
			&cli.StringFlag{
				Name:    flagModule,
				Aliases: []string{"M"},
				Usage:   "use module",
			},
			&cli.StringFlag{
				Name:    flagLibPath,
				Aliases: []string{"I"},
				Usage:   "directory for library modules",
			},
			&cli.StringFlag{
				Name:  flagEncoding,
				Usage: "specify encoding",
			},
			&cli.BoolFlag{
				Name:    flagInPlace,
				Aliases: []string{"i"},
				Usage:   "edit files in place",
			},
			&cli.BoolFlag{
				Name:    flagPrint,
				Aliases: []string{"p"},
				Usage:   "assume loop like -n but print line also",
			},
			&cli.BoolFlag{
				Name:    flagLoop,
				Aliases: []string{"n"},
				Usage:   "assume loop around script",
			},
			&cli.BoolFlag{
				Name:    flagAutoSplit,
				Aliases: []string{"a"},
				Usage:   "autosplit mode with -n or -p",
			},
			&cli.BoolFlag{
				Name:    flagCheckSyntax,
				Aliases: []string{"c"},
				Usage:   "check syntax only (runs BEGIN and END blocks)",
			},
			&cli.BoolFlag{
				Name:    flagWarnings,
				Aliases: []string{"w"},
				Usage:   "enable warnings",
			},
			&cli.BoolFlag{
				Name:  flagStrict,
				Usage: "enable strict mode",
			},
			&cli.BoolFlag{
				Name:    flagDebug,
				Aliases: []string{"d"},
				Usage:   "run program under debugger",
			},
			&cli.BoolFlag{
				Name:    flagTaint,
				Aliases: []string{"T"},
				Usage:   "enable taint checks",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "perl: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments
	for i := 0; i < c.NArg(); i++ {
		params = append(params, yup.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.IsSet(flagScript) {
		params = append(params, Script(c.String(flagScript)))
	}
	if c.IsSet(flagModule) {
		params = append(params, Module(c.String(flagModule)))
	}
	if c.IsSet(flagLibPath) {
		params = append(params, LibPath(c.String(flagLibPath)))
	}
	if c.IsSet(flagEncoding) {
		params = append(params, Encoding(c.String(flagEncoding)))
	}
	if c.Bool(flagInPlace) {
		params = append(params, InPlace)
	}
	if c.Bool(flagPrint) {
		params = append(params, Print)
	}
	if c.Bool(flagLoop) {
		params = append(params, Loop)
	}
	if c.Bool(flagAutoSplit) {
		params = append(params, AutoSplit)
	}
	if c.Bool(flagCheckSyntax) {
		params = append(params, CheckSyntax)
	}
	if c.Bool(flagWarnings) {
		params = append(params, Warnings)
	}
	if c.Bool(flagStrict) {
		params = append(params, Strict)
	}
	if c.Bool(flagDebug) {
		params = append(params, Debug)
	}
	if c.Bool(flagTaint) {
		params = append(params, Taint)
	}

	// Create and execute the perl command
	cmd := Perl(params...)
	return yup.Run(cmd)
}
