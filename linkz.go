package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	strftime "github.com/hhkbp2/go-strftime"
)

var blogRoot = path.Join(os.Getenv("HOME"), "code/international.github.com/")
var blogFolder = path.Join(blogRoot, "_posts")

func obtainNextNumber(dir string) (int, error) {
	dirHandles, err := ioutil.ReadDir(dir)

	if err != nil {
		return 0, err
	}

	nextNumber := 0

	linksFiles := make([]int, 0)

	for _, handle := range dirHandles {
		fileName := handle.Name()
		if strings.Contains(fileName, "linx") {
			linksFiles = append(linksFiles, extractNum(fileName))
		}
	}

	sort.Ints(linksFiles)
	nextNumber = linksFiles[len(linksFiles)-1]

	return nextNumber, nil
}

func extractNum(name string) int {
	r := regexp.MustCompile(`linx[^0-9]+(\d+)\.md`)
	results := r.FindStringSubmatch(name)

	if results == nil {
		return 0
	}

	matchedNumber := results[1]
	num, err := strconv.ParseInt(matchedNumber, 0, 10)
	if err != nil {
		return 0
	}

	return int(num) + 1
}

func genNextBlogPost(num int) error {
	templatePath := path.Join(blogRoot, "tpl-list.md")
	fileHandle, err := os.Open(templatePath)

	if err != nil {
		return err
	}

	defer fileHandle.Close()

	templateContent, err := ioutil.ReadAll(fileHandle)
	if err != nil {
		return err
	}

	date := time.Now()
	formattedNow := strftime.Format("%Y-%m-%d-%H-%M", date)
	newFileName := fmt.Sprintf("%s-linx-%d.md", formattedNow, num)
	newFileName = path.Join(blogFolder, newFileName)

	return ioutil.WriteFile(newFileName, templateContent, 0644)
}

func main() {
	nextNum, err := obtainNextNumber(blogFolder)
	if err != nil {
		panic(err)
	}
	err = genNextBlogPost(nextNum)
	if err != nil {
		panic(err)
	}
}
