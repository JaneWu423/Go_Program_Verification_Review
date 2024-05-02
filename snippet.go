package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type Snippet struct {
	Name    string
	Content string
	Type    int
}

func extractSnippets(snippets_folder string) []Snippet {
	snippets := []Snippet{}

	files, err := ioutil.ReadDir(snippets_folder)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(snippets_folder + "/" + file.Name())

		if err != nil {
			panic(err)
		}

		snippets = append(snippets, initSnippet(file, string(content)))

	}

	return snippets
}

func initSnippet(file os.FileInfo, content string) Snippet {
	s := Snippet{
		Name:    strings.Replace(file.Name(), ".go", "", -1),
		Content: content,
		Type:    extractSnippetType(content)}

	return s
}

func extractSnippetType(content string) int {
	t := CH

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if strings.Contains(line, "type =") {
			ty := strings.Trim(strings.Split(line, "=")[1], " ")

			switch ty {
			case "CH":
				t = CH
			case "WG":
				t = WG
			case "MUTEX":
				t = MUTEX
			case "RWMUTEX":
				t = RWMUTEX
			}
		}
	}

	return t
}
