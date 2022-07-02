package data

import (
	"os"
	"path"
)

var DoraemonDocumentTsvPath = path.Join(os.Getenv("PROJECT_ROOT"), "cmd/seeds/data/ja/doraemon.tsv")
