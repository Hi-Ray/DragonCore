package docs

import (
	repo "dragonback/lib/constants/repositories"
	"dragonback/lib/modules"
)

var files = modules.GitModule().FetchMarkdownFiles(repo.CDragonDocs)
