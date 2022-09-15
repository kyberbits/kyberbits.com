package resources

import (
	"embed"
	"io/fs"
)

// Public is foobar
var Public fs.FS

//go:embed *
var everything embed.FS

func init() {
	x, err := fs.Sub(everything, "public")
	if err != nil {
		panic(err)
	}
	Public = x
}
