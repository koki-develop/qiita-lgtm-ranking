package report

import (
	"bytes"
	"text/template"
	"time"

	"github.com/koki-develop/qiita-lgtm-ranking/pkg/infrastructures/qiita"
	"github.com/pkg/errors"
)

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build(from, to time.Time, items qiita.Items) (string, error) {
	tpl, err := template.New("daily.template.md").Funcs(template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}).ParseFiles("./templates/daily.template.md")
	if err != nil {
		return "", errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	items.Sort()
	if err := tpl.Execute(buf, map[string]interface{}{
		"from":  from,
		"to":    to,
		"items": items,
	}); err != nil {
		return "", errors.WithStack(err)
	}

	return buf.String(), nil
}
