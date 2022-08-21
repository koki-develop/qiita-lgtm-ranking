package report

import (
	"bytes"
	"text/template"

	"github.com/koki-develop/qiita-lgtm-ranking/pkg/infrastructures/qiita"
	"github.com/pkg/errors"
)

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

type BuildOptions struct {
	Tags       Tags
	Conditions Conditions
	Items      qiita.Items
}

func (b *Builder) Build(opts *BuildOptions) (string, error) {
	tpl, err := template.New("daily.template.md").Funcs(template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}).ParseFiles("./templates/daily.template.md")
	if err != nil {
		return "", errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	opts.Items.Sort()
	if err := tpl.Execute(buf, map[string]interface{}{
		"tags":       opts.Tags,
		"conditions": opts.Conditions,
		"items":      opts.Items,
	}); err != nil {
		return "", errors.WithStack(err)
	}

	return buf.String(), nil
}

type Tag struct {
	ReportID string `json:"report_id"`
	Name     string `json:"tag"`
}

type Tags []*Tag

type Condition struct {
	Key   string
	Value string
}

type Conditions []*Condition
