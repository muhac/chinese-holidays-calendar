package read

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/parse/core"
	"regexp"
	"strconv"
	"sync"
)

type dataLoader struct {
	Dir string
}

type fileInfo struct {
	Name string
	Year int
}

func (l dataLoader) Read() (result core.Raw) {
	resultChan := make(chan core.RawInfo)
	wg := new(sync.WaitGroup)

	for _, f := range l.fileList() {
		wg.Add(1)

		go func(file fileInfo) {
			defer wg.Done()
			raw, err := l.load(file.Name)
			if err != nil {
				fmt.Printf("Error loading %s: %s", file.Name, err)
				return
			}
			resultChan <- core.RawInfo{Year: file.Year, Data: raw}
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

func (l dataLoader) fileList() (result []fileInfo) {
	files, err := ioutil.ReadDir(l.Dir)
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

func (l dataLoader) load(filename string) (result string, err error) {
	data, err := ioutil.ReadFile(l.Dir + filename)
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
