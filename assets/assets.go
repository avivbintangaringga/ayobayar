package assets

import "embed"

//go:embed static
var static embed.FS

func Static() *embed.FS {
	return &static
}
