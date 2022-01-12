package handlers

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Path struct {
	Result string `json:"Result"`
	CyGt   string `json:"CyGt"`
	CyLt   string `json:"CyLt"`
	CyRt   string `json:"CyRt"`
	CyDb   string `json:"CyDb"`
}

func runPython(csvFile, outDir string) (Path, error) {
	cmd := exec.Command("./scripts/main.py", "-f", csvFile, "-s", outDir)
	stdout, err := cmd.Output()

	resultPath := Path{}
	if err != nil {
		fmt.Println(err.Error())
		return resultPath, err
	}

	json.Unmarshal(stdout, &resultPath)

	return resultPath, err
}
