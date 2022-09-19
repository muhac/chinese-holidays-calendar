package read

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/samber/lo"

	"main/parse/data"
)

func NewReader(dir, file string) data.Reader {
	return dataReader{Dir: "./" + dir + "/", File: file}
}

type dataReader struct {
	Dir  string
	File string
}

type fileInfo struct {
	Name string
	Year int
}

func (dw dataReader) Read() (result data.Input) {
	resultChan := make(chan data.InputRaw)
	wg := new(sync.WaitGroup)

	for _, f := range dw.fileList() {
		wg.Add(1)

		go func(file fileInfo) {
			defer wg.Done()
			raw, err := dw.load(file.Name)
			if err != nil {
				log.Printf("Error loading %s: %s\n", file.Name, err)
				return
			}

			res := data.InputRaw{
				Year: file.Year,
				Data: lines(raw),
			}
			if len(res.Data) == 0 {
				log.Printf("No data in %s\n", file.Name)
				return
			}

			resultChan <- res
		}(f)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for content := range resultChan {
		result = append(result, content)
	}
	return result
}

func (dw dataReader) fileList() (result []fileInfo) {
	files, err := os.ReadDir(dw.Dir)
	if err != nil {
		log.Fatal(err)
	}

	return lo.FilterMap(files, func(file os.DirEntry, _ int) (fileInfo, bool) {
		yr, e := year(file.Name(), dw.File)
		isFile := e == nil && !file.IsDir()
		return fileInfo{Name: file.Name(), Year: yr}, isFile
	})
}

func (dw dataReader) load(filename string) (result string, err error) {
	content, err := os.ReadFile(dw.Dir + filename)
	if err != nil {
		return result, err
	}
	return string(content), nil
}

func year(filename, format string) (result int, err error) {
	regex := regexp.MustCompile(format)
	if !regex.MatchString(filename) {
		return 0, fmt.Errorf("%s is not a valid filename", filename)
	}
	return strconv.Atoi(filename[:4])
}

func lines(data string) (result []string) {
	var (
		dateSingle = `(\d?\d.\d?\d)`
		dateRange  = fmt.Sprintf(`(%s-%s)`, dateSingle, dateSingle)
		dateFormat = fmt.Sprintf(`(%s|%s)`, dateSingle, dateRange)
		dateInputs = fmt.Sprintf(`(%s,)*%s`, dateFormat, dateFormat)
		dateAccept = fmt.Sprintf(`(|%s)`, dateInputs)
		dateRegex  = regexp.MustCompile(fmt.Sprintf(`^[^;]+;%s;%s$`, dateAccept, dateAccept))
	)

	return lo.FilterMap(strings.Split(data, "\n"),
		func(line string, _ int) (string, bool) {
			line = strings.Split(line, "//")[0]
			line = strings.TrimSpace(line)
			return line, dateRegex.MatchString(line)
		},
	)
}
