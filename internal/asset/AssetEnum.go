package asset

import "strings"

type String struct {
	Value string
}

type assetEnum struct {
	Thyo  string
	Garan string
	Kchol string
	Sahol string
}

var Types = newAssets()

func newAssets() *assetEnum {
	return &assetEnum{
		Thyo:  "thyao-hisse",
		Garan: "garan-hisse",
		Kchol: "kchol-hisse",
		Sahol: "sahol-hisse",
	}
}

func (str *String) Slice() string {
	replaced := strings.ReplaceAll(str.Value, "\n", "")
	replaced = strings.TrimSpace(replaced)
	return replaced
}
