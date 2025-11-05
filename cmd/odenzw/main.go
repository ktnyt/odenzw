package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/go-gts/flags"
	"github.com/ktnyt/go-perm"
	"github.com/mattn/go-isatty"
)

func run(ctx *flags.Context) error {
	pos, opt := flags.Flags()

	s := pos.String("string", "string to permute")
	g := opt.String('g', "include", "", "pattern to include in output")
	v := opt.String('v', "exclude", "", "pattern to exclude from output")
	if _, err := flags.Parse(pos, opt, ctx.Args); err != nil {
		return err
	}

	include, err := regexp.Compile(*g)
	if err != nil {
		return err
	}

	exclude, err := regexp.Compile(*v)
	if err != nil {
		return err
	}

	a := []rune(*s)
	p := perm.Permute(a)
	ss := make([]string, 0, 120)
	for p.Next() {
		if *v == "" || !exclude.MatchString(string(a)) || (*g != "" && include.MatchString(string(a))) {
			ss = append(ss, string(a))
		}
	}
	for i := 0; i < len(ss); i += 10 {
		sep := "\n"
		if isatty.IsTerminal(os.Stdout.Fd()) {
			sep = " "
		}
		fmt.Println(strings.Join(ss[i:i+10], sep))
	}

	return nil
}

var VERSION = flags.Version{
	Major: 1,
	Minor: 1,
	Patch: 0,
}

func main() {
	os.Exit(flags.Run("odenzw", "odenzw generator", VERSION, run))
}
