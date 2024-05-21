package main

import (
	_ "embed"
	"os"
	"strings"
	"text/template"
)

//go:embed help.md.tmpl
var helpTextTmpl string

type HelpContext struct {
	CommandName string
	BuildInfo   *BuildInfo
}

func HelpText() string {

	tmpl, err := template.New("help").Parse(helpTextTmpl)
	if err != nil {
		return ""
	}

	var helpContext HelpContext
	helpContext.CommandName = os.Args[0]
	helpContext.BuildInfo = GetBuildInfo()

	var result strings.Builder
	err = tmpl.Execute(&result, helpContext)
	if err != nil {
		return ""
	}

	return result.String()
}
