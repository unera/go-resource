package {{ .Config.Package }}

import (
	"encoding/base64"
	"fmt"
)

// R - return resource (as blob) by name
func R(name string) {{ if .Config.Error }}([]byte, error){{else}}[]byte{{end}} {
	var b64value string

	{{- range .Fields }}
	if name == "{{ .Name }}" {
		b64value = "" +
                        {{- range .PValue }}
			"{{ . }}" +
			{{- end }}
			""
		goto found			
        }
	{{ end }}

	{{ if .Config.Error }}
	return nil, fmt.Errorf("Resource '%s' is not found", name)
        {{- else -}}
	panic(fmt.Sprintf("Resource '%s' is not found", name))
	{{- end -}}
{{ if ne (0) (len .Fields) }}
found:
{{ end }}
	{{ if .Config.Error }}
	return base64.StdEncoding.decode(b64value)
        {{ else }}
	decoded, err := base64.StdEncoding.DecodeString(b64value)
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
