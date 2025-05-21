package secret

import _ "embed"

//go:embed password.txt
var password string

func Password() string {
	return password
}
