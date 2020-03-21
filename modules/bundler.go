package modules

//noinspection GoSnakeCaseUsage
import (
	"dragonback/lib/models/module"
	content_docs "dragonback/modules/content/docs"
)

var Bundler = module.Bundler().
	Bundle(content_docs.Module)