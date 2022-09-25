package output

import (
	"fmt"
	"hash/crc32"
	"math/rand"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"main/parse/core"
	"main/parse/data"
)

func NewFormatter(name string) data.Formatter {
	return formatter{name}
}

type formatter struct {
	name string
}

func (f formatter) Format(info core.Holidays) (result data.Output) {
	result.Prefix = fmt.Sprintf(icsHead, f.name)
	result.Suffix = icsTail

	uuid.SetRand(rand.New(rand.NewSource(int64(crc32.ChecksumIEEE([]byte(f.name))))))

	result.Body = lo.Map(info, func(day core.Holiday, i int) string {
		return event{
			id:    uuid.NewString(),
			group: day.Group,
			title: getTitle(day),
			date:  day.Date,
			desc:  getDesc(day),
		}.Ics()
	})

	return
}
