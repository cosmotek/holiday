package holiday

import (
	"errors"
	"fmt"
	"os/exec"
)

type ShellDependency struct {
	ExeName     string
	ProgramName string
	Required    bool
	Description string
	*InstallationInfo

	VersionFunc VersionFunc
}

type InstallationInfo struct {
	Installed      bool
	Version        string
	ExecutablePath string
}

func (s *ShellDependency) LoadInstallationInfo() error {
	path, err := exec.LookPath(s.ExeName)
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			s.InstallationInfo = &InstallationInfo{Installed: false}
			return nil
		}
	}

	version, err := s.VersionFunc()
	if err != nil {
		return err
	}

	s.InstallationInfo = &InstallationInfo{
		Installed:      true,
		ExecutablePath: path,
		Version:        version,
	}

	return nil
}

func (s ShellDependency) Summarize() string {
	if !s.Installed {
		var statusItem DoctorItemStatus = FAIL
		if !s.Required {
			statusItem = WARN
		}

		return SprintStatusMessage(statusItem, "", fmt.Sprintf("%s\n    %s✗ %s%s executable is missing from $PATH%s", s.ProgramName, red, white, s.ExeName, reset))
	}

	shortSummary := s.SummarizeShort()

	return fmt.Sprintf("%s\n    %s•%s %s at %s", shortSummary, green, reset, s.ExeName, s.ExecutablePath)
}

func (s ShellDependency) SummarizeShort() string {
	if !s.Installed {
		var statusItem DoctorItemStatus = FAIL
		if !s.Required {
			statusItem = WARN
		}

		return SprintStatusMessage(statusItem, "", fmt.Sprintf("%s\n    %s✗ %s%s executable is missing from $PATH%s", s.ProgramName, red, white, s.ExeName, reset))
	}

	return SprintStatusMessage(PASS, "", fmt.Sprintf("%s (v%s)", s.ProgramName, s.Version))
}
