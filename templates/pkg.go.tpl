package {{ .Config.Package }}

import (
	"fmt"
)

var resData = map[string][]byte{
    {{ range .Fields }}
	"{{ .Name }}": []byte{ {{"\n\t\t"}}
	    {{- range $i, $v := .BValue }}
		{{- $v | printf "%4d" }},{{ if (isWrap $i) }}{{"\n\t\t"}}{{end -}}
	    {{ end }}
	},
    {{ end }}
}

// R - return resource (as blob) by name
func R(name string) ([]byte{{ if .Config.Error }}, error{{end}}) {

	data, ok := resData[name]

	if !ok {
	    {{ if .Config.Error }}
		return nil, fmt.Errorf("Unknown resource: %s", name)
	    {{ else }}
		panic(fmt.Sprintf("Unknown resource: %s", name))
	    {{ end }}
	}
	
	return data{{ if .Config.Error }}, nil{{end}}
}

// Rs - return resource (as string) by name
func Rs(name string) (string{{ if .Config.Error }}, error{{end}}) {
    {{ if .Config.Error }}
	res, err := R(name)
	if err != nil {
	    return "", err
	}
	return string(res), nil
    {{ else }}
	return string(R(name))
    {{ end }}
}
