# タグ別

{{ range .tags }}[`{{ .Name }}`](https://qiita.com/items/{{ .ReportID }}) {{ end }}

# 集計について

{{ range .conditions -}}
- {{ .Key }}: {{ .Value }}
{{ end -}}

# GitHub

<a href="https://github.com/koki-develop/qiita-lgtm-ranking"><img src="https://github-link-card.s3.ap-northeast-1.amazonaws.com/koki-develop/qiita-lgtm-ranking.png" width="460px"></a>
スターをもらえるととっても励みになります :bow:

# いいね数ランキング

{{ if .items -}}
{{ range $i, $item := .items -}}
## {{ inc $i }} 位: [{{ $item.Title }}]({{ $item.URL }})

{{ range $item.Tags }}[`{{ .Name }}`](https://qiita.com/tags/{{ .Name }}) {{ end }}
**{{ $item.LikesCount }}** いいね　**{{ $item.StockersCount }}** ストック
[@{{ $item.User.ID }}](https://qiita.com/{{ $item.User.ID }}) さん ( {{ $item.CreatedAt.Format "2006-01-02 15:04" }} に投稿 )
{{ end -}}
{{ else -}}
ランキングに入る記事が見つかりませんでした。
{{ end -}}
