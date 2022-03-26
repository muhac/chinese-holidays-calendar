package read

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/parse/data"
	"regexp"
	"strconv"
	"strings"
	"sync"
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
				fmt.Printf("Error loading %s: %s", file.Name, err)
				return
			}

			result := data.InputRaw{
				Year: file.Year,
				Data: lines(raw),
			}
			if len(result.Data) == 0 {
				fmt.Printf("No data in %s", file.Name)
				return
			}

			resultChan <- result
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
	files, err := ioutil.ReadDir(dw.Dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if yr, err := year(file.Name(), dw.File); err == nil && !file.IsDir() {
			result = append(result, fileInfo{Name: file.Name(), Year: yr})
		}
	}
	return result
}

func (dw dataReader) load(filename string) (result string, err error) {
	content, err := ioutil.ReadFile(dw.Dir + filename)
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

	for _, line := range strings.Split(data, "\n") {
		line = strings.Split(line, "//")[0]
		line = strings.TrimSpace(line)
		if dateRegex.MatchString(line) {
			result = append(result, line)
		}
	}
	return result
}
