package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	CH int = iota
	WG
	MUTEX
	RWMUTEX
	ALL
	EMPTY
)

type Context struct {
	Name    string // The name of the context
	Content string
	Type    int
	Bounds  []string
}

func extractContexts(contexts_path string) []Context {
	contexts := []Context{}

	files, err := ioutil.ReadDir(contexts_path)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(contexts_path + "/" + file.Name())

		if err != nil {
			panic(err)
		}

		contexts = append(contexts, InitContext(file, string(content)))

	}

	return contexts
}

func InitContext(file os.FileInfo, content string) Context {
	ctx := Context{
		Name:    strings.Replace(file.Name(), ".go", "", -1),
		Content: string(removeComment(content)),
		Type:    extractContextType(content),
		Bounds:  extractBounds(content),
	}

	return ctx
}

func (c Context) Print() {
	fmt.Println(" ==== ======= ===")
	fmt.Println(" NAME : ", c.Name)
	fmt.Println(" Type : ", c.Type)
	fmt.Println(" Bounds : ", c.Bounds)
	fmt.Println(" ==== ======= ===")
}

func extractContextType(content string) int {
	t := ALL

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
			default:
				t = ALL
			}
		}
	}
	return t
}

func removeComment(content string) string {
	lines := strings.Split(content, "\n")

	without_comment := ""
	for _, line := range lines {
		if !strings.Contains(line, "bounds") {
			without_comment += line + "\n"
		}
	}

	return without_comment
}

func extractBounds(content string) []string {
	bounds := []string{}

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if strings.Contains(line, "bounds =") {
			splitted_bounds := strings.Split(strings.Split(line, "=")[1], ",")

			for _, bound := range splitted_bounds {
				bound = strings.Trim(bound, " ") // removing outside spaces

				bounds = append(bounds, bound)
			}
		}
	}

	return bounds
}
