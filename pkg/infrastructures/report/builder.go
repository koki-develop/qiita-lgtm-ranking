package report

import (
	"time"

	"github.com/koki-develop/qiita-lgtm-ranking/pkg/infrastructures/qiita"
)

type Builder struct{}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Build(from, to time.Time, items qiita.Items) (string, error) {
	return "report", nil
}
