package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type practicals struct {
	Practical []struct {
		Command string `json:"command"`
	} `json:"practical"`
}

func main() {
	jsonFile, err := os.Open("commands.json")
	defer jsonFile.Close()
	must(err)

	var listOfCommands practicals

	bytes, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(bytes, &listOfCommands)
	must(err)

	output := NewWriter("OS.md")

	for _, command := range listOfCommands.Practical {
		cmd := exec.Command(os.Getenv("SHELL"), "-c", command.Command)
		stdout, stderr, err := checkCommand(cmd)
		must(err)
		output.PrintCommand(command.Command)
		output.PrintOutput(stdout, stderr)
	}

	output.Save()
}

func checkCommand(cmd *exec.Cmd) (string, string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", "", err
	}
	return out.String(), stderr.String(), nil
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
