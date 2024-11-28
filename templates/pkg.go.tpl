package {{ .Config.Package }}

import (
	"encoding/base64"
	"fmt"
)

// R - return resource (as blob) by name
func R(name string) {{ if .Config.Error }}([]byte, error){{else}}[]byte{{end}} {
	{{ range .Fields }}
	const r{{ .No }} = "" +
		{{- range .PValue }}
		"{{ . }}" +
		{{- end }}
		""
	{{ end }}
	var (
		decoded []byte
		err     error
	)

	switch name {
	{{- range .Fields }}
	case "{{ .Name }}": 
		decoded, err = base64.StdEncoding.DecodeString(r{{ .No }})
	{{ end }}
	default:
		err = fmt.Errorf("Resource '%s' is not found", name)
	}
	{{ if .Config.Error }}
	return decoded, err
        {{ else }}
	if err != nil {
		panic(err.Error())
	}
	return decoded
	{{ end }}
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
