package output

import (
	"fmt"
	"github.com/google/uuid"
	"hash/crc32"
	"main/parse/base"
	"main/parse/data"
	"math/rand"
)

func NewFormatter(format string) data.Formatter {
	return formatterICS{"节假日"}
}

type formatterICS struct{
	name string
}

func (f formatterICS) Format(info base.Holidays) (result data.Output) {
	result.Prefix = fmt.Sprintf(IcsHead, f.name)
	result.Suffix = IcsTail

	uuid.SetRand(rand.New(rand.NewSource(int64(crc32.ChecksumIEEE([]byte(f.name))))))

	for _, day := range info {
		outputDay := event{
			Id:    uuid.NewString(),
			Group: day.Group,
			Title: getTitle(day),
			Date:  day.Date,
			Desc:  getDesc(day),
		}
		result.Body = append(result.Body, outputDay.Ics())
	}
	return
}
