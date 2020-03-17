package sogvin

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gregoryv/sogvin/internal"
	. "github.com/gregoryv/web"
)

type Book struct {
	Title string
	pages []*Page
}

// Saves all pages and table of contents
func (book *Book) SaveTo(base string) error {
	for _, page := range book.pages {
		page.SaveTo(base)
	}
	return nil
}

func findH1(article *Element) string {
	var buf bytes.Buffer
	w := NewHtmlWriter(&buf)
	w.WriteHtml(article)
	from := bytes.Index(buf.Bytes(), []byte("<h1>")) + 4
	to := bytes.Index(buf.Bytes(), []byte("</h1>"))
	return strings.TrimSpace(string(buf.Bytes()[from:to]))
}

func (book *Book) AddPage(right string, article *Element) *Element {
	title := findH1(article)
	// todo strip title from tags
	filename := filenameFrom(title) + ".html"

	page := newPage(
		filename,
		stripTags(title)+" - "+book.Title,
		PageHeader(right+" - "+A(Href("index.html"), "Software Engineering").String()),
		article,
		footer,
	)
	book.pages = append(book.pages, page)
	return linkToPage(page)
}

func linkToPage(page *Page) *Element {
	return Li(A(Href(page.Filename), findH1(page.Element)))
}

func newPage(filename, title string, header, article, footer *Element) *Page {
	return NewPage(filename,
		Html(en,
			Head(utf8, viewport, theme, a4, Title(title)),
			Body(header, article, footer),
		),
	)
}

func stripTags(in string) string {
	var buf bytes.Buffer
	var inside bool
	for _, r := range in {
		switch r {
		case '<':
			inside = true
		case '>':
			inside = false
		default:
			if inside {
				continue
			}
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func filenameFrom(in string) string {
	tidy := bytes.NewBufferString("")
	var inside bool
	for _, c := range in {
		switch c {
		case '(', ')':
			continue
		case ' ':
			tidy.WriteRune('_')
		case '<':
			inside = true
		case '>':
			inside = false
		default:
			if inside {
				continue
			}
			tidy.WriteString(strings.ToLower(string(c)))
		}
	}
	return tidy.String()
}

var (
	en       = Lang("en")
	utf8     = Meta(Charset("utf-8"))
	viewport = Meta(
		Name("viewport"),
		Content("width=device-width, initial-scale=1.0"),
	)
	theme  = Stylesheet("theme.css")
	a4     = Stylesheet("a4.css")
	footer = Footer(myname)
	myname = "Gregory Vin&ccaron;i&cacute;"
)

func PageHeader(right string) *Element {
	h := Header()
	if right != "" {
		h = h.With(Code(right))
	}
	return h
}

// Stylesheet returns a link web element
func Stylesheet(href string) *Element {
	return Link(Rel("stylesheet"), Type("text/css"), Href(href))
}

// Boxnote returns a small box aligned to the left with given top
// margin in cm.
func Sidenote(txt string, cm float64) *Element {
	return Div(Class("sidenote"),
		&Attribute{
			Name: "style",
			Val:  fmt.Sprintf("margin-top: %vcm", cm),
		},
		Div(Class("inner"), txt),
	)
}

// LoadGoFile returns a pre web element wrapping the contents from the
// given file. If to == -1 all lines to the end of file are returned.
func LoadGoFile(filename string, span ...int) *Element {
	from, to := 0, -1
	if len(span) == 2 {
		from, to = span[0], span[1]
	}
	v := internal.LoadFile(filename, from, to)
	class := "srcfile"
	if from == 0 && to == -1 {
		class += " complete"
	}
	return Pre(Class(class), Code(Class("go"), v))
}

// AnyFile returns a pre web element wrapping the contents from the
// given file. If to == -1 all lines to the end of file are returned.
func AnyFile(filename, showas string, from, to int) *Element {
	v := internal.LoadFile(filename, from, to)
	class := "srcfile"
	if from == 0 && to == -1 {
		class += " complete"
	}
	return Div(
		Div(Class("filename"), showas),
		Pre(Class(class), Code(Class("go"), v)),
	)
}

func gregoryv(name, txt string) *Element {
	return Li(
		fmt.Sprintf(
			`<a href="https://github.com/gregoryv/%s">%s</a> - %s`,
			name, name, txt,
		),
	)
}

// ShellCommand returns a web Element wrapping shell commands
func ShellCommand(v string) *Element {
	return Pre(Class("command"), Code(v))
}
