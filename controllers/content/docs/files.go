package docs

import (
	"dragonback/lib/models/repo"
)

var files, _ = repo.New(repo.CDragonDocs).Documentation()
