package ui

import "embed"

//comment directive, this instructs Go to store the files from
// ui/html and ui/static folder in an embed.FS embedded filesystem
// referenced by the global var Files

//go:embed "html" "static"
var Files embed.FS
