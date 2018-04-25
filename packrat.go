package packrat

import (
	"html/template"

	"github.com/gobuffalo/packr"
)

func ParseBox(name string) (*template.Template, error) {
	var t *template.Template
	box := packr.NewBox(name)

	for _, file := range box.List() {
		if t == nil {
			t = template.New(file)
		}

		if t.Name() != file {
			t = t.New(file)
		}

		if _, err := t.Parse(box.String(file)); err != nil {
			return nil, err
		}
	}

	return t, nil
}
