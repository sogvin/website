package website

import (
	"strings"

	. "github.com/gregoryv/web"
)

func Version() string {
	latest := MustQueryOne(Changelog, "h3.latest")
	parts := strings.Split(latest.Text(), " ")
	return parts[0]
}

func Released() string {
	latest := MustQueryOne(Changelog, "h3.latest")
	parts := strings.Split(latest.Text(), " ")
	return parts[1]
}

var Changelog = Article(Class("changelog"),
	H1("Changelog"),

	P(`All notable changes to this project will be documented in
        this file.`),

	H3("0.5.2", Class("latest"), Span("2024-10-05")),
	Ul(
		Li("Fix typo navstart"),
		Li("Update dependencies"),
		Li("Use Go v1.21"),
	),

	H3("0.5.1", Span("2022-05-22")),
	Ul(
		Li("Move module to github.com/sogvin/website"),
	),

	H3("0.5.0", Span("2022-01-20")),
	Ul(
		Li("Link to changelog"),
		Li("Rename sections to match skills"),
		Li("Split design section into system and software design"),
		Li(`Add drill "Pointer receiver or not"`),
	),

	H3("0.4.0", Span("2022-01-16")),
	Ul(
		Li("Add preface with purpose of the website"),
		Li("Add drills",
			Ul(
				Li("Encode struct to json"),
				Li("Read file line by line"),
				Li("Simple level loggers"),
				Li("Basic use of log printer fucs"),
				Li("Open file for reading"),
				Li("Slurp file"),
				Li("Parse options using cmdline.BasicParser"),
				Li("Parse builtin types"),
				Li("Short and long option names"),
			),
		),
		Li("Group packages under references"),
	),

	H3("0.3.0", Span("2022-01-15")),
	Ul(
		Li("Added version and release date to front page"),
		Li(`Update package cmdline and "Purpose of func main" article`),
	),

	H3("0.2.0", Span("2021-04-05")),
	H3("0.1.0", Span("2020-02-09")),
)
