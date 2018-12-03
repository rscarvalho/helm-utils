package helm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

// GetReleaseStatus gets a status for a release
func GetReleaseStatus(releaseName string) (*ReleaseStatus, error) {
	statusBytes, err := runHelm("status", releaseName)

	if err != nil {
		return nil, err
	}

	status := &ReleaseStatus{}

	if err := json.Unmarshal(statusBytes, status); err == nil {
		return status, nil
	}

	return nil, err
}

// InstallOrUpgrade executes the helm command by passing args
func InstallOrUpgrade(chartPath string, releaseName string, dryRun bool, extraArgs []string) error {
	return nil
}

func runHelm(extraArgs ...string) ([]byte, error) {
	helmPath, err := exec.LookPath("helm")

	if err != nil {
		return nil, err
	}

	var args []string

	for _, arg := range extraArgs {
		args = append(args, arg)
	}

	args = append(args, "--output", "json")

	cmd := exec.Command(helmPath, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	if err != nil {
		errorMessage := string(stderr.Bytes())
		fmt.Println("helm command failed:\n", errorMessage)
		return nil, NewError(errorMessage, &err)
	}

	output := stdout.Bytes()
	return output, err
}
