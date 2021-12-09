package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/cosmotek/holiday"
)

func forgeGenPluginVersionFunc(pluginExeName string) holiday.VersionFunc {
	return func() (string, error) {
		versionOutput, err := exec.Command(pluginExeName, "version").Output()
		if err != nil {
			return "", err
		}

		versionOutputStr := strings.TrimPrefix(strings.TrimSuffix(string(versionOutput), "\n"), "v")
		return versionOutputStr, nil
	}
}

func main() {
	var deps = []holiday.ShellDependency{
		{
			ExeName:     "go",
			Required:    true,
			ProgramName: "Go Language",
			Description: "Go Language Tools",
			VersionFunc: holiday.GoVersionFunc,
		},
		{
			ExeName:     "forge-plug-gen-go-server",
			ProgramName: "Forge Go Server Generation Plugin",
			Description: "Forge Go Server Generation Plugin",
			VersionFunc: forgeGenPluginVersionFunc("forge-plug-gen-go-server"),
		},
		{
			ExeName:     "forge-plug-gen-go-client",
			ProgramName: "Forge Go Client Generation Plugin",
			Description: "Forge Go Client Generation Plugin",
			VersionFunc: forgeGenPluginVersionFunc("forge-plug-gen-go-client"),
		},
		{
			ExeName:     "forge-plug-gen-dart-client",
			ProgramName: "Forge Dart Client Generation Plugin",
			Description: "Forge Dart Client Generation Plugin",
			VersionFunc: forgeGenPluginVersionFunc("forge-plug-gen-dart-client"),
		},
		{
			ExeName:     "forge-plug-gen-js-client",
			ProgramName: "Forge JS Client Generation Plugin",
			Description: "Forge JS Client Generation Plugin",
			VersionFunc: forgeGenPluginVersionFunc("forge-plug-gen-js-client"),
		},
	}

	fmt.Println("Doctor summary:")
	issueCount, err := holiday.PrintEval(holiday.PrintOpts{
		OutputVerbose: true,
		Output:        nil, // defaults to stdout
	}, deps...)
	if err != nil {
		panic(err)
	}

	if issueCount > 0 {
		suffix := "category"
		if issueCount > 1 {
			suffix = "categories"
		}

		fmt.Printf("\n\n! Doctor found issues in %d %s.\n", issueCount, suffix)
	}
}
