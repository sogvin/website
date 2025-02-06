package website

import (
	"bytes"
	_ "embed"
	"strings"

	. "github.com/gregoryv/web"
	"github.com/russross/blackfriday/v2"
)

func Version() string {
	prefix := "## ["
	from := strings.Index(changelog, prefix) + len(prefix)
	to := from + strings.Index(changelog[from:], "]")
	return changelog[from:to]
}

var Changelog = Article(Class("changelog"),
	H1("Changelog"),
	string(
		bytes.ReplaceAll(
			blackfriday.Run(
				stripFirstLine([]byte(changelog)),
			),
			[]byte("h2"),
			[]byte("h3"),
		),
	),
)

//go:embed changelog.md
var changelog string

func stripFirstLine(txt []byte) []byte {
	i := bytes.Index(txt, []byte{'\n'})
	if i > 0 {
		return txt[i+1:]
	}
	return txt
}
