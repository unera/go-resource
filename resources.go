package main

import (
	"encoding/base64"
	"fmt"
)

// R - return resource (as blob) by name
func R(name string) []byte {
	var b64value string
	if name == "pkg.go.tpl" {
		b64value = "" +
			"cGFja2FnZSB7eyAuQ29uZmlnLlBhY2thZ2UgfX0KCmltcG9ydC" +
			"AoCgkiZW5jb2RpbmcvYmFzZTY0IgoJImZtdCIKKQoKLy8gUiAt" +
			"IHJldHVybiByZXNvdXJjZSAoYXMgYmxvYikgYnkgbmFtZQpmdW" +
			"5jIFIobmFtZSBzdHJpbmcpIHt7IGlmIC5Db25maWcuRXJyb3Ig" +
			"fX0oW11ieXRlLCBlcnJvcil7e2Vsc2V9fVtdYnl0ZXt7ZW5kfX" +
			"0gewoJdmFyIGI2NHZhbHVlIHN0cmluZwoKCXt7LSByYW5nZSAu" +
			"RmllbGRzIH19CglpZiBuYW1lID09ICJ7eyAuTmFtZSB9fSIgew" +
			"oJCWI2NHZhbHVlID0gIiIgKwogICAgICAgICAgICAgICAgICAg" +
			"ICAgICB7ey0gcmFuZ2UgLlBWYWx1ZSB9fQoJCQkie3sgLiB9fS" +
			"IgKwoJCQl7ey0gZW5kIH19CgkJCSIiCgkJZ290byBmb3VuZAkJ" +
			"CQogICAgICAgIH0KCXt7IGVuZCB9fQoKCXt7IGlmIC5Db25maW" +
			"cuRXJyb3IgfX0KCXJldHVybiBuaWwsIGZtdC5FcnJvcmYoIlJl" +
			"c291cmNlICclcycgaXMgbm90IGZvdW5kIiwgbmFtZSkKICAgIC" +
			"AgICB7ey0gZWxzZSAtfX0KCXBhbmljKGZtdC5TcHJpbnRmKCJS" +
			"ZXNvdXJjZSAnJXMnIGlzIG5vdCBmb3VuZCIsIG5hbWUpKQoJe3" +
			"stIGVuZCAtfX0Ke3sgaWYgbmUgKDApIChsZW4gLkZpZWxkcykg" +
			"fX0KZm91bmQ6Cnt7IGVuZCB9fQoJe3sgaWYgLkNvbmZpZy5Fcn" +
			"JvciB9fQoJcmV0dXJuIGJhc2U2NC5TdGRFbmNvZGluZy5kZWNv" +
			"ZGUoYjY0dmFsdWUpCiAgICAgICAge3sgZWxzZSB9fQoJZGVjb2" +
			"RlZCwgZXJyIDo9IGJhc2U2NC5TdGRFbmNvZGluZy5EZWNvZGVT" +
			"dHJpbmcoYjY0dmFsdWUpCglpZiBlcnIgIT0gbmlsIHsKCSAgIC" +
			"BwYW5pYyhlcnIuRXJyb3IoKSkKCX0KCXJldHVybiBkZWNvZGVk" +
			"Cgl7eyBlbmQgfX0KfQoKLy8gUnMgLSByZXR1cm4gcmVzb3VyY2" +
			"UgKGFzIHN0cmluZykgYnkgbmFtZQpmdW5jIFJzKG5hbWUgc3Ry" +
			"aW5nKSAoc3RyaW5ne3sgaWYgLkNvbmZpZy5FcnJvciB9fSwgZX" +
			"Jyb3J7e2VuZH19KSB7CiAgICB7eyBpZiAuQ29uZmlnLkVycm9y" +
			"IH19CglyZXMsIGVyciA6PSBSKG5hbWUpCglpZiBlcnIgIT0gbm" +
			"lsIHsKCSAgICByZXR1cm4gIiIsIGVycgoJfQoJcmV0dXJuIHN0" +
			"cmluZyhyZXMpLCBuaWwKICAgIHt7IGVsc2UgfX0KCXJldHVybi" +
			"BzdHJpbmcoUihuYW1lKSkKICAgIHt7IGVuZCB9fQp9Cg==" +
			""
		goto found			
        }
	
	if name == "resources.yaml.tpl" {
		b64value = "" +
			"LS0tCiMgUGFja2FnZSBuYW1lIGZvciByZXN1bHQgLmdvCnBhY2" +
			"thZ2U6IG1haW4KCiMgaWdub3JlIHBhdHRlcm5zIG9yIG5hbWVz" +
			"Cmlnbm9yZToKIC0gXnJlc291cmNlc1wueWFtbCQKCiMgQ29kZW" +
			"dlbmVyYXRlZCBmdW5jdGlvbnMgcmV0dXJucyBlcnJvciAob3Ig" +
			"cGFuaWMpCnVzZV9lcnJvcjogZmFsc2UKLi4uCg==" +
			""
		goto found			
        }
	

	panic(fmt.Sprintf("Resource '%s' is not found", name))
found:

	
	decoded, err := base64.StdEncoding.DecodeString(b64value)
	if err != nil {
	    panic(err.Error())
	}
	return decoded
	
}

// Rs - return resource (as string) by name
func Rs(name string) (string) {
    
	return string(R(name))
    
}
