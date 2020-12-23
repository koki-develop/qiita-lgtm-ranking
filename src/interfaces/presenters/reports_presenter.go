package presenters

import (
	"fmt"
	"strings"
	"time"

	"github.com/kou-pg-0131/qiita-lgtm-ranking/src/domain"
)

// IReportsPresenter .
type IReportsPresenter interface {
	WeeklyPerTag(from time.Time, items *domain.Items, tag string) (string, error)
}

// ReportsPresenter .
type ReportsPresenter struct{}

// NewReportsPresenter .
func NewReportsPresenter() *ReportsPresenter {
	return new(ReportsPresenter)
}

// WeeklyPerTag .
func (p *ReportsPresenter) WeeklyPerTag(from time.Time, items *domain.Items, tag string) (string, error) {
	rows := []string{}

	rows = append(rows, "# 集計期間")
	rows = append(rows, "")
	rows = append(rows, fmt.Sprintf("%s ~ %s", from.Format("2006-01-02"), from.AddDate(0, 0, 7).Format("2006-01-02")))
	rows = append(rows, "")

	rows = append(rows, "# LGTM 数ランキング")
	rows = append(rows, "")

	items.SortByLikesCount()
	for i, item := range *items.HasLGTM() {
		rows = append(rows, fmt.Sprintf("## %d 位: [%s](%s)", i+1, item.Title, item.URL))
		rows = append(rows, "")

		tags := []string{}
		for _, tag := range item.Tags {
			tags = append(tags, fmt.Sprintf("[`%s`](https://qiita.com/tags/%s)", tag.Name, tag.Name))
		}
		rows = append(rows, strings.Join(tags, " "))
		rows = append(rows, "")

		rows = append(rows, fmt.Sprintf("**%d** LGTM", item.LikesCount))
		rows = append(rows, fmt.Sprintf("[@%s](https://qiita.com/%s) さん ( %s に投稿 )", item.User.ID, item.User.ID, item.CreatedAt.Format("2006-01-02 15:04")))
	}

	return strings.Join(rows, "\n"), nil
}
