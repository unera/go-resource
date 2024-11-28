package main

import (
	"encoding/base64"
	"fmt"
)

// R - return resource (as blob) by name
func R(name string) []byte {

	const r1 = "" +
		"cGFja2FnZSB7eyAuQ29uZmlnLlBhY2thZ2UgfX0KCmltcG9ydC" +
		"AoCgkiZW5jb2RpbmcvYmFzZTY0IgoJImZtdCIKKQoKLy8gUiAt" +
		"IHJldHVybiByZXNvdXJjZSAoYXMgYmxvYikgYnkgbmFtZQpmdW" +
		"5jIFIobmFtZSBzdHJpbmcpIHt7IGlmIC5Db25maWcuRXJyb3Ig" +
		"fX0oW11ieXRlLCBlcnJvcil7e2Vsc2V9fVtdYnl0ZXt7ZW5kfX" +
		"0gewoJe3sgcmFuZ2UgLkZpZWxkcyB9fQoJY29uc3Qgcnt7IC5O" +
		"byB9fSA9ICIiICsKCQl7ey0gcmFuZ2UgLlBWYWx1ZSB9fQoJCS" +
		"J7eyAuIH19IiArCgkJe3stIGVuZCB9fQoJCSIiCgl7eyBlbmQg" +
		"fX0KCXZhciAoCgkJZGVjb2RlZCBbXWJ5dGUKCQllcnIgICAgIG" +
		"Vycm9yCgkpCgoJc3dpdGNoIG5hbWUgewoJe3stIHJhbmdlIC5G" +
		"aWVsZHMgfX0KCWNhc2UgInt7IC5OYW1lIH19IjogCgkJZGVjb2" +
		"RlZCwgZXJyID0gYmFzZTY0LlN0ZEVuY29kaW5nLkRlY29kZVN0" +
		"cmluZyhye3sgLk5vIH19KQoJe3sgZW5kIH19CglkZWZhdWx0Og" +
		"oJCWVyciA9IGZtdC5FcnJvcmYoIlJlc291cmNlICclcycgaXMg" +
		"bm90IGZvdW5kIiwgbmFtZSkKCX0KCXt7IGlmIC5Db25maWcuRX" +
		"Jyb3IgfX0KCXJldHVybiBkZWNvZGVkLCBlcnIKICAgICAgICB7" +
		"eyBlbHNlIH19CglpZiBlcnIgIT0gbmlsIHsKCQlwYW5pYyhlcn" +
		"IuRXJyb3IoKSkKCX0KCXJldHVybiBkZWNvZGVkCgl7eyBlbmQg" +
		"fX0KfQoKLy8gUnMgLSByZXR1cm4gcmVzb3VyY2UgKGFzIHN0cm" +
		"luZykgYnkgbmFtZQpmdW5jIFJzKG5hbWUgc3RyaW5nKSAoc3Ry" +
		"aW5ne3sgaWYgLkNvbmZpZy5FcnJvciB9fSwgZXJyb3J7e2VuZH" +
		"19KSB7CiAgICB7eyBpZiAuQ29uZmlnLkVycm9yIH19CglyZXMs" +
		"IGVyciA6PSBSKG5hbWUpCglpZiBlcnIgIT0gbmlsIHsKCSAgIC" +
		"ByZXR1cm4gIiIsIGVycgoJfQoJcmV0dXJuIHN0cmluZyhyZXMp" +
		"LCBuaWwKICAgIHt7IGVsc2UgfX0KCXJldHVybiBzdHJpbmcoUi" +
		"huYW1lKSkKICAgIHt7IGVuZCB9fQp9Cg==" +
		""

	const r2 = "" +
		"LS0tCiMgUGFja2FnZSBuYW1lIGZvciByZXN1bHQgLmdvCnBhY2" +
		"thZ2U6IG1haW4KCiMgaWdub3JlIHBhdHRlcm5zIG9yIG5hbWVz" +
		"Cmlnbm9yZToKIC0gXnJlc291cmNlc1wueWFtbCQKCiMgQ29kZW" +
		"dlbmVyYXRlZCBmdW5jdGlvbnMgcmV0dXJucyBlcnJvciAob3Ig" +
		"cGFuaWMpCnVzZV9lcnJvcjogZmFsc2UKLi4uCg==" +
		""

	var (
		decoded []byte
		err     error
	)

	switch name {
	case "pkg.go.tpl":
		decoded, err = base64.StdEncoding.DecodeString(r1)

	case "resources.yaml.tpl":
		decoded, err = base64.StdEncoding.DecodeString(r2)

	default:
		err = fmt.Errorf("Resource '%s' is not found", name)
	}

	if err != nil {
		panic(err.Error())
	}
	return decoded

}

// Rs - return resource (as string) by name
func Rs(name string) string {

	return string(R(name))

}
