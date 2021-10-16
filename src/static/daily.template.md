# 集計について

- 集計対象日: {{ .from.Format "2006-01-02" }}

# GitHub

<a href="https://github.com/koki-develop/qiita-lgtm-ranking"><img src="https://github-link-card.s3.ap-northeast-1.amazonaws.com/koki-develop/qiita-lgtm-ranking.png" width="460px"></a>
スターをもらえるととっても励みになります :bow:

# LGTM 数ランキング

{{ range $i, $item := .items }}## {{ inc $i }} 位: [{{ $item.Title }}]({{ $item.URL }})

{{ range $item.Tags }}[`{{ .Name }}`](https://qiita.com/tags/{{ .Name }}) {{ end }}

**{{ $item.LikesCount }}** LGTM
[@{{ $item.User.ID }}](https://qiita.com/{{ $item.User.ID }}) さん ( {{ $item.CreatedAt.Format "2006-01-02 15:04" }} に投稿 )
{{ end }}
