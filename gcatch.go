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

type GCatchVerifier struct {
}

func (g GCatchVerifier) Verify(model Model) Verification {
	verification := Verification{
		ToolName:  "GCatch",
		ModelName: model.GetName(),
		Result:    NO_BUG,
		Type:      model.Type,
		Content:   model.Content,
	}

	// Put the model inside a .go file
	err := ioutil.WriteFile("./source/main.go", []byte(model.Content), 0644)

	if err != nil {
		panic(err)
	}

	// verify using Gomela
	path, _ := filepath.Abs("./source")

	var output bytes.Buffer

	command := exec.Command("GCatch", "-checker=BMOC:double:unlock", "-compile-error", "-path="+path)
	command.Stdout = &output

	pre := time.Now()
	err = command.Run()
	after := time.Now()

	if err != nil {
		panic(err)
	}

	verification.Time = after.Sub(pre).Milliseconds()

	if strings.Contains(output.String(), "----------Bug[1]----------") {
		if strings.Contains(output.String(), "src/sync/once.go") && !strings.Contains(output.String(), "----------Bug[2]----------") {
		} else {
			verification.Result = BUG
		}
	}
	verification.Feedback = output.String()

	os.Remove("./source/main.go")
	return verification
}
