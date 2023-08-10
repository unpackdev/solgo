package solc

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// RealCommand implements the Commander interface using real solc-select commands.
type RealCommand struct {
	current string
}

// Current returns the current version of solc in use.
func (rc *RealCommand) Current() string {
	return rc.current
}

// Install a specific version of solc.
// Returns a boolean indicating if the version was installed, a slice of output lines from solc-select, and an error if any occurred.
func (rc *RealCommand) Install(version string) (bool, []string, error) {
	cmd := exec.Command("solc-select", "install", version)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false, nil, err
	}

	// Split the output into lines
	outputLines := strings.Split(out.String(), "\n")

	// Check if the version was successfully installed
	isInstalled := false
	for _, line := range outputLines {
		if strings.Contains(line, "Version '"+version+"' installed.") {
			isInstalled = true
			break
		}
	}

	return isInstalled, outputLines, nil
}

// Use a specific version of solc. If the version does not exist, install it.
// Returns a boolean indicating success, the outputs from the install and use commands, and an error if any occurred.
func (rc *RealCommand) Use(version string) (bool, []string, []string, error) {
	// If the desired version is already the current version, do nothing and return
	if rc.Current() == version {
		return true, nil, nil, nil
	}

	// Check if the version exists
	versions, err := rc.Versions()
	if err != nil {
		return false, nil, nil, err
	}

	versionExists := false
	for _, v := range versions {
		if v.Release == version {
			versionExists = true
			break
		}
	}

	// If the version is not installed, install it and capture the output
	var installOutput []string
	if !versionExists {
		success, output, err := rc.Install(version)
		if err != nil || !success {
			return false, output, nil, err
		}
		installOutput = output
	}

	// Switch to the desired version and capture the output
	cmd := exec.Command("solc-select", "use", version)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return false, installOutput, nil, err
	}
	useOutput := strings.Split(out.String(), "\n")

	// Check if the switch was successful
	success := strings.Contains(out.String(), "Switched global version to")

	return success, installOutput, useOutput, nil
}

// Versions lists all available versions of solc using solc-select.
func (rc *RealCommand) Versions() ([]Version, error) {
	cmd := exec.Command("solc-select", "versions")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	// Parse the output to get the list of versions
	versionsOutput := out.String()
	versions := strings.Split(versionsOutput, "\n")

	var versionList []Version
	for _, version := range versions {
		if version != "" {
			isCurrent := strings.Contains(version, "current")
			release := strings.Fields(version)[0]
			versionList = append(versionList, Version{
				Release: release,
				Current: isCurrent,
			})
		}
	}

	if len(versionList) == 0 {
		return nil, errors.New("no solc versions found")
	}

	return versionList, nil
}

// Upgrade to the latest version of solc.
// Returns a bool indicating if solc-select is up to date, and an error if any occurred.
func (rc *RealCommand) Upgrade() (bool, error) {
	cmd := exec.Command("solc-select", "upgrade")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false, err
	}

	// Check the output to see if solc-select is already up to date
	response := out.String()
	if strings.Contains(response, "already up to date") {
		return true, nil
	}

	return false, nil
}
