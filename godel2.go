package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Godel2Verifier struct {
}

func (g Godel2Verifier) Verify(model Model) Verification {

	verification := Verification{
		ToolName:  "Godel2",
		ModelName: model.GetName(),
		Result:    NO_BUG,
		Type:      model.Type,
		Content:   model.Content,
	}
	// Put the model inside a .go file
	err := ioutil.WriteFile("./test.go", []byte(model.Content), 0644)

	if err != nil {
		panic(err)
	}

	// verify using Gomela
	var err_output bytes.Buffer

	fl, _ := os.Create("./example.cgo")

	defer func() {
		os.Remove("test.go")
		os.Remove("example.cgo")
	}()
	path, _ := filepath.Abs(".")
	writer := bufio.NewWriter(fl)
	command := exec.Command("docker", "run", "-i", "--rm", "-v", path+":/root", "jgabet/godel2:latest", "migoinfer", "test.go")
	command.Dir = "."
	command.Stdout = writer
	command.Stderr = &err_output
	command.Run()
	writer.Flush()

	if err != nil {
		panic(err)
	}

	var output bytes.Buffer

	command = exec.Command("docker", "run", "-i", "--rm", "-v", path+":/root", "jgabet/godel2:latest", "Godel", "example.cgo")
	command.Stdout = &output
	command.Stderr = &err_output

	pre := time.Now()
	err = command.Run()
	after := time.Now()

	if err != nil {
		verification.Result = CRASH
		verification.Feedback = err_output.String()
		return verification
	}

	verification.Time = after.Sub(pre).Milliseconds()
	verification.Feedback = output.String() + err_output.String()

	// Parse output and print if it found the bug
	verification.Result = parseGodel2Result(output.String())
	// Delete program

	// Delete results folder
	removeGodel2Data()

	return verification
}

func removeGodel2Data() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".mcf") {
			os.Remove(path)
		}

		if strings.HasSuffix(path, ".lps") {
			os.Remove(path)
		}
		if strings.HasSuffix(path, ".crl2") {
			os.Remove(path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func parseGodel2Result(result string) VerificationResult {
	splitted_lines := strings.Split(result, "\n")

	for _, line := range splitted_lines {
		if strings.Contains(line, "Finite Control:") && strings.Contains(line, "False") {
			return NO_SUPPORT
		}
		if strings.Contains(line, "(exit 1): failed") {
			return CRASH
		}
		if strings.Contains(line, "No global deadlock:") && strings.Contains(line, "False") {
			return BUG
		}
		if strings.Contains(line, "Liveness:") && strings.Contains(line, "False") {
			return BUG
		}
		if strings.Contains(line, "Safety:") && strings.Contains(line, "False") {
			return BUG
		}
	}
	return NO_BUG
}
