package excel2conf

var genJsonTmpl = `{
  "{{.Name}}s": 
{{- range $i,$v := .Datas}}
{{- if eq $i 0}}
	{
{{- end}}
    "{{$v.PriKey}}": 
		{
{{- range $j,$vv := $v.Fields}}
		  "{{$vv.FieldName}}": {{$vv.FieldVal}}
{{- if lt $j $v.FieldsLen}}
				,
{{- end}}
{{- end}}
    	}
{{- if lt $i $.DatasLen}}
	,
{{- end}}
{{- end}}
	}
}
`
