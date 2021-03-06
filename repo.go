package website

import (
	"path"

	. "github.com/gregoryv/web"
)

// Repo is used to generate and load files
type Repo struct {
	host  string
	local string
}

func (me *Repo) LinkedLabel(pth string) *Element {
	label := path.Join(path.Base(me.local), pth)
	return Div(Class("filename"),
		A(
			Href(me.host+"/blob/main/"+pth),
			label,
		),
	)
}

func (me *Repo) loadFile(pth string, span ...int) *Element {
	return Wrap(
		me.LinkedLabel(pth),
		loadFile(path.Join(me.local, pth), span...),
	)
}
