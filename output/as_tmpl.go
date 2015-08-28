package output

const asTmpl = `
{{if eq .AS.StartAutnum .AS.EndAutnum}}\
aut-num:     {{.AS.StartAutnum}}
{{else}}\
aut-num:     {{.AS.StartAutnum}} - {{.AS.EndAutnum}}
{{end}}\
{{if ne .AS.Type ""}}\
type:        {{.AS.Type}}
{{end}}\
{{if ne .AS.Country ""}}\
country:     {{.AS.Country}}
{{end}}\
created:     {{formatDate .CreatedAt}}
changed:     {{formatDate .UpdatedAt}}
{{range .IPNetworks}}\
inetnum:     {{.}}
{{end}}\
{{range .AS.RoutingPolicy}}\
{{if gt .Cost 0}}\
as-in:       from AS{{.Traffic}} {{.Cost}} accept {{.Policy}}
{{else}}\
as-out:      to AS{{.Traffic}} announce {{.Autnum}}
{{end}}\
{{end}}\

` + contactTmpl
