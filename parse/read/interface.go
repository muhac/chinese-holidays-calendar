package read

import "main/parse/base"

type Parser interface {
	Data() []base.Holiday
}

func InDir(path string) Parser {
	return dataLoader{Dir: path}
}
