package word

import (
	"errors"
	. "github.com/legion-zver/go-docx-templates"
	"strings"
)

// LegionZverDocx is implementation of DocxTemplate interface
//
//	Using "github.com/legion-zver/go-docx-templates"
type LegionZverDocx struct {
	t    *DocxTemplateFile
	path string
}

func (r *LegionZverDocx) OpenTemplate(path string) (DocxTemplate, error) {
	// Validation
	if !strings.HasSuffix(path, ".docx") {
		return nil, errors.New("incorrect file format")
	}
	// Open template
	template, err := OpenTemplate(path)
	if err != nil {
		return nil, err
	}
	return &LegionZverDocx{t: template, path: path}, nil
}

func (r *LegionZverDocx) RenderTemplate(data interface{}) error {
	if r.t == nil {
		return errors.New("template has not initials or already closed")
	}
	return r.t.RenderTemplate(data)
}

func (r *LegionZverDocx) Save(filePath string) error {
	if r.t == nil {
		return errors.New("template has not initials or already closed")
	}
	return r.t.Save(filePath)
}

func (r *LegionZverDocx) CloseTemplate() error {
	if r.t == nil {
		return errors.New("template has not initials or already closed")
	}
	r.t = nil
	return nil
}

func (r *LegionZverDocx) ResetTemplate() error {
	template, err := OpenTemplate(r.path)
	if err != nil {
		return err
	}
	r.t = template
	return nil
}
