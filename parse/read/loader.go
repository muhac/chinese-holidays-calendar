package read

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/parse/base"
	"sync"
)

type dataLoader struct {
	Dir string
}

func (l dataLoader) Data() (result []base.Holiday) {
	resultChan := make(chan base.Holiday)
	wg := new(sync.WaitGroup)
	for _, filename := range l.fileList() {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			data, err := l.load(filename)
			if err != nil {
				fmt.Printf("Error loading %s: %s", filename, err)
				return
			}
			resultChan <- data
		}(filename)
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

func (l dataLoader) fileList() []string {
	files, err := ioutil.ReadDir(l.Dir)
	if err != nil {
		log.Fatal(err)
	}

	var result []string
	for _, file := range files {

		if _, ok := year(file.Name()); ok && !file.IsDir() {
			result = append(result, file.Name())
			fmt.Println("using", file.Name())
		}
	}
	return result
}

func (l dataLoader) load(filename string) (result base.Holiday, err error) {
	data, err := ioutil.ReadFile(l.Dir + filename)
	if err != nil {
		return result, err
	}
	return parse(filename, string(data))
}
