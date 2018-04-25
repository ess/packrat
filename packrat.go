package packrat

import (
	"html/template"
)

type storage interface {
	MustString(string) (string, error)
}

type Packrat struct {
	box storage
}

func New(box storage) *Packrat {
	return &Packrat{box: box}
}

func (p *Packrat) ParseFiles(filenames ...string) (*template.Template, error) {
	var t *template.Template

	if len(filenames) == 0 {
		return nil, fmt.Errorf("packrat: no files named in call to ParseFiles")
	}

	for _, file := range filenames {
		if t == nil {
			t = template.New(file)
		}

		if t.Name() != file {
			t = t.New(file)
		}

		content, err := p.box.MustString(file)
		if err != nil {
			return nil, err
		}

		if _, err := t.Parse(content); err != nil {
			return nil, err
		}
	}

	return t, nil
}
