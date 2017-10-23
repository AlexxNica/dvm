// +build !windows

package main

import (
	"os"
	"path/filepath"
)

const binaryFileExt string = ""

func upgradeSelf(version string) {
	binaryURL := buildDvmReleaseURL(version, dvmOS, dvmArch, "dvm-helper")
	binaryPath := filepath.Join(dvmDir, "dvm-helper", "dvm-helper")
	downloadFileWithChecksum(binaryURL, binaryPath)

	scriptURL := buildDvmReleaseURL(version, "dvm.sh")
	scriptPath := filepath.Join(dvmDir, "dvm.sh")
	downloadFile(scriptURL, scriptPath)
}

func getCleanPathRegex() string {
	versionDir := getVersionsDir()
	return versionDir + `/[^:]+:`
}

func validateShellFlag() {
	// we don't care about the shell flag on non-Windows platforms
}

func getUserHomeDir() string {
	return os.Getenv("HOME")
}
