package assets

import "embed"

//go:embed static
var assets embed.FS

func GetFS() *embed.FS {
	return &assets
}
