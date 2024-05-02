package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type GomelaVerifier struct {
}

func (g GomelaVerifier) Verify(model Model) Verification {

	verification := Verification{
		ToolName:  "gomela",
		ModelName: model.GetName(),
		Result:    NO_BUG,
		Type:      model.Type,
		Content:   model.Content,
	}
	// Put the model inside a .go file
	err := ioutil.WriteFile("./source/test.go", []byte(model.Content), 0644)

	if err != nil {
		panic(err)
	}

	// verify using Gomela
	var output bytes.Buffer
	var err_output bytes.Buffer

	command := exec.Command("gomela", "fs", "source")
	command.Stdout = &output
	command.Stderr = &err_output

	pre := time.Now()
	err = command.Run()
	after := time.Now()

	if err != nil {
		panic(err)
	}

	verification.Time = after.Sub(pre).Milliseconds()

	// Parse output and print if it found the bug
	verification.Feedback = output.String()
	if strings.Contains(output.String(), "Error : too many processes.") {
		verification.Result = NO_SUPPORT
	}
	if strings.Contains(output.String(), "true.") {
		verification.Result = BUG
	} else if !strings.Contains(output.String(), "Send on close safety error") && output.String() != "" {
		verification.Result = CRASH
	}
	// Delete program
	os.Remove("./source/test.go")
	// Delete results folder
	removeResultFolder()

	return verification
}

func removeResultFolder() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		if strings.Contains(path, "result2022") {
			os.RemoveAll(path)
		}

		if strings.Contains(path, ".trail") {
			os.Remove(path)
		}

		if strings.HasSuffix(path, "pan") {
			os.Remove(path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
