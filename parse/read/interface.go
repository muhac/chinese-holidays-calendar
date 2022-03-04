package read

import (
	"main/parse/core"
)

type Reader interface {
	Read() core.Raw
}

func From(path string) Reader {
	return dataLoader{Dir: path}
}
