package read

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/parse/data"
	"regexp"
	"strconv"
	"sync"
)

func NewReader(dir string) data.Reader {
	return dataReader{Dir: dir}
}

type dataReader struct {
	Dir string
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
			resultChan <- data.InputRaw{Year: file.Year, Data: raw}
		}(f)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for data := range resultChan {
		result = append(result, data)
	}
	return result
}

func (dw dataReader) fileList() (result []fileInfo) {
	files, err := ioutil.ReadDir(dw.Dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if yr, err := year(file.Name()); err == nil && !file.IsDir() {
			result = append(result, fileInfo{Name: file.Name(), Year: yr})
		}
	}
	return result
}

func (dw dataReader) load(filename string) (result string, err error) {
	data, err := ioutil.ReadFile(dw.Dir + filename)
	if err != nil {
		return result, err
	}
	return string(data), nil
}

func year(filename string) (result int, err error) {
	regex := regexp.MustCompile(`^\d{4}\.txt$`)
	if !regex.MatchString(filename) {
		return 0, fmt.Errorf("%s is not a valid filename", filename)
	}
	return strconv.Atoi(filename[:4])
}
