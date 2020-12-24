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

	rows = append(rows, "# 他のタグ")
	rows = append(rows, "")

	tags := []string{
		"[`AWS`](https://qiita.com/items/e24b6279326a462d456c)",
		"[`Docker`](https://qiita.com/items/ae11fca7d2eba445b037)",
		"[`iOS`](https://qiita.com/items/e61a29a383d0403e92fc)",
		"[`Java`](https://qiita.com/items/4c3f84836bfdbb137226)",
		"[`JavaScript`](https://qiita.com/items/eaa7ac5b62a0a723edbb)",
		"[`PHP`](https://qiita.com/items/3318cbdbc45c6ebd4014)",
		"[`Python`](https://qiita.com/items/9d7f2ffeafb36cf59a77)",
		"[`Rails`](https://qiita.com/items/93b9e7f7d143e9ce650e)",
		"[`React`](https://qiita.com/items/f9712f8acace22815b99)",
		"[`Ruby`](https://qiita.com/items/72c3d2e896bdc3e1a6b3)",
		"[`Swift`](https://qiita.com/items/e2b6f0645e29f0e2b761)",
		"[`TypeScript`](https://qiita.com/items/25b7c0870afa6d41d19b)",
		"[`Vim`](https://qiita.com/items/f5361177baef95e447d1)",
		"[`Vue.js`](https://qiita.com/items/2774e02c6eea5c830d99)",
	}
	rows = append(rows, strings.Join(tags, " "))
	rows = append(rows, "")

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
		rows = append(rows, "")
	}

	return strings.Join(rows, "\n"), nil
}
