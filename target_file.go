package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	spdxPrefix = "SPDX-License-Identifier:"
)

// TargetFile represents a potential writable file
type TargetFile struct {
	data []string
	path string
}

// NewTargetFile creates and returns an instance of TargetFile
func NewTargetFile(path string) *TargetFile {
	return &TargetFile{
		path: path,
	}
}

/* -------------------- Exported Functions -------------------- */

// AddLicense prepends the SPDX license entry to the file
func (tFile *TargetFile) AddLicense(license string) error {
	tFile.load()

	if tFile.HasLicense() {
		return nil
	}

	err := tFile.write(license)
	return err
}

// HasLicense returns TRUE if this file already contains a license entry,
// FALSE if it does not
func (tFile *TargetFile) HasLicense() bool {
	return strings.Contains(tFile.data[0], spdxPrefix)
}

/* -------------------- Unexported Functions -------------------- */

func (tFile *TargetFile) licenseString(license string) string {
	return fmt.Sprintf("// %s %s\n", spdxPrefix, license)
}

func (tFile *TargetFile) load() error {
	f, err := os.Open(tFile.path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		row := scanner.Text()
		tFile.data = append(tFile.data, row)
	}

	return nil
}

func (tFile *TargetFile) write(license string) error {
	if len(tFile.data) == 0 {
		return fmt.Errorf("file is empty: %s", tFile.path)
	}

	// Jam the license string to the front of the file data
	licStr := tFile.licenseString(license)

	tmp := []string{licStr}
	tmp = append(tmp, tFile.data...)

	tFile.data = tmp

	f, err := os.OpenFile(tFile.path, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, line := range tFile.data {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", line))
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
