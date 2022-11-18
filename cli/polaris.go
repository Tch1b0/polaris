package main

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

var (
	//go:embed solution.template
	solutionTemplate []byte

	//go:embed test.template
	testTemplate []byte

	//go:embed readme.template
	readmeTemplate []byte
)

func main() {
	dayNum, err := strconv.Atoi(requireArg(1))
	if err != nil {
		panic(err)
	}
	day := fmt.Sprintf("%02d", dayNum)
	dayDir := "./days/" + day
	dayVars := map[string]string{
		"DAY_NAME": day,
		"DAY_NUM":  fmt.Sprint(dayNum),
		"DAY_PATH": dayDir,
	}

	testTemplate = replaceVars(testTemplate, dayVars)
	readmeTemplate = replaceVars(readmeTemplate, dayVars)

	os.MkdirAll(dayDir, 0777)
	os.WriteFile(path.Join(dayDir, "solution.go"), solutionTemplate, 0644)
	os.WriteFile(path.Join(dayDir, "solution_test.go"), testTemplate, 0644)
	os.WriteFile(path.Join(dayDir, "README.md"), readmeTemplate, 0644)
}

func help() {
	fmt.Println(`Polaris CLI
    usage: polaris [DAY]`)
	os.Exit(0)
}

func replaceVars(text []byte, vals map[string]string) []byte {
	newTxt := string(text)
	for k, v := range vals {
		newTxt = strings.ReplaceAll(newTxt, fmt.Sprintf("{{%s}}", k), v)
	}

	return []byte(newTxt)
}

func requireArg(pos int) string {
	if len(os.Args) < pos {
		help()
		return ""
	} else {
		return os.Args[pos]
	}
}
