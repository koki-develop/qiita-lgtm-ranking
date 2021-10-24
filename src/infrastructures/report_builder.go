package infrastructures

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/koki-develop/qiita-lgtm-ranking/src/entities"
	"github.com/pkg/errors"
)

type ReportBuilder struct{}

func NewReportBuilder() *ReportBuilder {
	return &ReportBuilder{}
}

func (b *ReportBuilder) Daily(from time.Time, items entities.Items) (*entities.Report, error) {
	tpl, err := template.New("daily.template.md").Funcs(template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}).ParseFiles("./src/static/daily.template.md")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	items.Sort()
	if err := tpl.Execute(buf, map[string]interface{}{
		"from":  from,
		"to":    from.AddDate(0, 0, 1),
		"items": items,
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &entities.Report{
		Title: "Qiita デイリー LGTM 数ランキング【自動更新】",
		Body:  buf.String(),
		Tags:  entities.Tags{{Name: "Qiita"}, {Name: "lgtm"}, {Name: "ランキング"}},
	}, nil
}

func (b *ReportBuilder) DailyByTag(from time.Time, items entities.Items, tag string) (*entities.Report, error) {
	tpl, err := template.New("dailyByTag.template.md").Funcs(template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}).ParseFiles("./src/static/dailyByTag.template.md")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	items.Sort()
	if err := tpl.Execute(buf, map[string]interface{}{
		"from":  from,
		"to":    from.AddDate(0, 0, 1),
		"items": items,
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &entities.Report{
		Title: fmt.Sprintf("【%s】Qiita デイリー LGTM 数ランキング【自動更新】", tag),
		Body:  buf.String(),
		Tags:  entities.Tags{{Name: "Qiita"}, {Name: "lgtm"}, {Name: "ランキング"}, {Name: tag}},
	}, nil
}

func (b *ReportBuilder) Weekly(from time.Time, items entities.Items) (*entities.Report, error) {
	tpl, err := template.New("weekly.template.md").Funcs(template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}).ParseFiles("./src/static/weekly.template.md")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	items.Sort()
	if err := tpl.Execute(buf, map[string]interface{}{
		"min_stock": 10,
		"from":      from,
		"to":        from.AddDate(0, 0, 7),
		"items":     items,
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &entities.Report{
		Title: "Qiita 週間 LGTM 数ランキング【自動更新】",
		Body:  buf.String(),
		Tags:  entities.Tags{{Name: "Qiita"}, {Name: "lgtm"}, {Name: "ランキング"}},
	}, nil
}

func (b *ReportBuilder) WeeklyByTag(from time.Time, items entities.Items, tag string) (*entities.Report, error) {
	tpl, err := template.New("weeklyByTag.template.md").Funcs(template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}).ParseFiles("./src/static/weeklyByTag.template.md")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	items.Sort()
	if err := tpl.Execute(buf, map[string]interface{}{
		"min_stock": 2,
		"from":      from,
		"to":        from.AddDate(0, 0, 7),
		"items":     items,
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return &entities.Report{
		Title: fmt.Sprintf("【%s】Qiita 週間 LGTM 数ランキング【自動更新】", tag),
		Body:  buf.String(),
		Tags:  entities.Tags{{Name: "Qiita"}, {Name: "lgtm"}, {Name: "ランキング"}, {Name: tag}},
	}, nil
}
