package holiday

import (
	"os/exec"
	"strings"
)

type VersionFunc func() (string, error)

func GoVersionFunc() (string, error) {
	goVersionOutput, err := exec.Command("go", "version").Output()
	if err != nil {
		return "", err
	}
	goVersionOutputStr := strings.TrimSuffix(string(goVersionOutput), "\n")

	var version string
	if strings.HasPrefix(goVersionOutputStr, "go version") {
		segs := strings.Split(strings.TrimPrefix(goVersionOutputStr, "go version go"), " ")
		version = segs[0]
	}

	return version, nil
}
