package holiday

import (
	"io"
	"os"
	"strings"
)

type PrintOpts struct {
	OutputVerbose bool
	Output        io.Writer
}

func PrintEval(opts PrintOpts, deps ...ShellDependency) (uint, error) {
	var writer io.Writer = os.Stdout
	if opts.Output != nil {
		writer = opts.Output
	}

	strs := []string{}
	issueCount := 0

	for _, dep := range deps {
		err := dep.LoadInstallationInfo()
		if err != nil {
			_, ierr := io.WriteString(writer, err.Error())
			if ierr != nil {
				return 0, ierr
			}

			return 1, nil
		}

		if !dep.Installed {
			issueCount++
		}

		var summaryStr string
		if opts.OutputVerbose {
			summaryStr = dep.Summarize()
		} else {
			summaryStr = dep.SummarizeShort()
		}

		strs = append(strs, summaryStr)
	}

	_, err := io.WriteString(writer, strings.Join(strs, "\n"))
	if err != nil {
		return 0, err
	}

	return uint(issueCount), nil
}
