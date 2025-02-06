package website

import (
	"os"
	"path/filepath"

	. "github.com/gregoryv/web"
)

func NewWebsite() *Website {
	title := "Software Engineering"
	author := "Gregory Vin&ccaron;i&cacute;"
	site := Website{
		title:  title,
		author: author,
	}
	site.ToSaver = &saveAll{&site}
	site.AddThemes(a4(), theme())

	article := Article(Class("toc"),
		H1(title, " - Skills &amp; Drills"),
		Img(Src("img/office.jpg")),
		P("Notes by ", author),

		H2("Preface"),

		Blockquote("Practice makes perfect"),

		P(`In software engineering there are many skills required to
        become professional. Here I present some of those skills and
        ways to practice them.`),

		H2("Start"),
		Ul(
			site.AddPage("Start", gettingStartedWithProgramming()),
		),

		H2("Plan"),

		P(`Skill of presenting a problem domain with a scoped solution
		in mind.`),

		H2("Design"),

		H3("System design"),

		P(`Skill of communicating with fellow engineers on what makes
		up a system and why.`),

		Ul(
			site.AddPage("Design", componentsDiagram()),
			site.AddPage("Design", roleBasedService()),
		),

		H3("Software design"),

		P(`Skill of writing software to grow gracefully over time.`),

		Ul(
			site.AddPage("Design", projectLayout()), // todo
			site.AddPage("Design", purposeOfFuncMain()),
			site.AddPage("Design", nexusPattern()),
			site.AddPage("Design", gracefulServerShutdown()),
			site.AddPage("Design", strictMode()),
		),

		H2("Verify"),
		Ul(
			site.AddPage("Verify", inlineTestHelpers()),
			site.AddPage("Verify", alternateDesign()),
			site.AddPage("Verify", setupTeardown()),
		),

		H2("Deliver"),
		Ul(
			site.AddPage("Deliver", embedVersionAndRevision()),
		),

		H2("Drills"),

		P(`Drills are short examples for practicing often used
		concepts.`),

		H3("Command line"),
		Ul(
			site.AddDrill("Command line", "", "drill/flag_types.go"),
			site.AddDrill("Command line", "", "drill/flag_names.go"),
			site.AddDrill("Command line", "", "drill/cmdline_basic.go"),
		),

		H3("Reading files"),
		Ul(
			site.AddDrill("Reading files", "", "drill/openfile.go"),
			site.AddDrill("Reading files", "", "drill/slurp_file.go"),
			site.AddDrill("Reading files", "", "drill/readfile_byline.go"),
		),

		H3("Logging"),
		Ul(
			site.AddDrill("Logging", "", "drill/logging.go"),
			site.AddDrill("Logging", "", "drill/level_logs.go"),
		),

		H3("Encoding"),
		Ul(
			site.AddDrill("Encoding", "", "drill/json_encode.go"),
		),

		H3("Methods"),
		Ul(
			site.AddDrill("Methods", "", "drill/pointer_receiver.go"),
			site.AddDrill("Methods", "", "drill/getters_and_setters.go"),
		),

		H2("References"),
		Ul(
			site.AddPage("References", packageRefs()),
		),
	)

	site.add(NewFile("index.html",
		Html(Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=1.0"),
				),
				stylesheet("theme.css"),
				stylesheet("a4.css"),
				Title(site.title),
			),
			Body(
				Header(Code(
					A(Href("changelog.html"), versionField()),
				)),
				article,
				Footer(),
			),
		),
	))

	site.AddPage("", Changelog)

	return &site
}

type Website struct {
	ToSaver

	title  string
	author string
	pages  []*Page
	themes []*CSS
	drills []*Page
}

// AddPage creates a new page and returns a link to it
func (me *Website) AddPage(right string, article *Element) *Element {
	title := MustQueryOne(article, "h1").Text()
	filename := filenameFrom(title) + ".html"

	backlink := A(Href("index.html"), me.title)
	if right != "" {
		backlink = Code(
			right, " - ", A(Href("index.html"), me.title),
		)
	}
	page := NewFile(filename,
		Html(Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=1.0"),
				),
				stylesheet("theme.css"),
				stylesheet("a4.css"),
				Title(stripTags(title)+" - "+me.title)),
			Body(
				Header(backlink),
				article,
				Footer(me.author),
			),
		),
	)
	me.add(page)

	return linkToPage(page)
}

func linkTo(article *Element) *Element {
	title := MustQueryOne(article, "h1").Text()
	filename := filenameFrom(title) + ".html"
	return A(Href(filename), title)
}

func (me *Website) AddDrill(right, args string, filename string) *Element {
	article := Article(
		loadExample(filename),
		example(args, filename),
	)
	page := NewFile(filepath.Base(toHtmlFile(filename)),
		Html(Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=1.0"),
				),
				stylesheet("../theme.css"),
				stylesheet("../a4.css"),
				Title(right, " - drill"),
			),
			Body(
				Header(
					Header(Code(
						right,
						" - ",
						A(Href("../index.html"), me.title).String(),
					)),
				),
				article,
				Footer(me.author),
			),
		),
	)
	me.drills = append(me.drills, page)
	return linkDrill(filename)
}

func (me *Website) AddThemes(v ...*CSS) {
	me.themes = append(me.themes, v...)
}

func (me *Website) add(p *Page) {
	me.pages = append(me.pages, p)
}

// ----------------------------------------
// Website behaviors

type ToSaver interface {
	SaveTo(base string) error
}

type saveAll struct {
	*Website
}

func (me *saveAll) SaveTo(base string) error {
	p := &savePagesOnly{me.Website}
	if err := p.SaveTo(base); err != nil {
		return err
	}

	for _, theme := range me.themes {
		theme.SaveTo(base)
	}
	drills := filepath.Join(base, "drill")
	os.MkdirAll(drills, 0755)
	for _, drill := range me.drills {
		if err := drill.SaveTo(drills); err != nil {
			return err
		}
	}
	return nil
}

type savePagesOnly struct {
	*Website
}

func (me *savePagesOnly) SaveTo(base string) error {
	for _, page := range me.pages {
		if err := page.SaveTo(base); err != nil {
			return err
		}
	}
	return nil
}
