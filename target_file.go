// SPDX-License-Identifier: MPL-2.0

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

	err := tFile.add(license)
	return err
}

// HasLicense returns TRUE if this file already contains a license entry,
// FALSE if it does not
func (tFile *TargetFile) HasLicense() bool {
	if len(tFile.data) == 0 {
		return false
	}

	return strings.Contains(tFile.data[0], spdxPrefix)
}

// HasData returns TRUE if this file has data to operate against,
// FALSE if it has no data to operate against
func (tFile *TargetFile) HasData() bool {
	return len(tFile.data) > 0
}

// RemoveLicense removed the specified license from the file
func (tFile *TargetFile) RemoveLicense(license string) error {
	tFile.load()

	if !tFile.HasLicense() {
		return nil
	}

	err := tFile.remove(license)
	return err
}

/* -------------------- Unexported Functions -------------------- */

func (tFile *TargetFile) add(license string) error {
	if !tFile.HasData() {
		return fmt.Errorf("file is empty: %s", tFile.path)
	}

	licStr := tFile.licenseString(license)

	tmp := []string{licStr}
	prependedData := append(tmp, tFile.data...)

	err := tFile.writeToFile(tFile.path, prependedData)
	return err
}

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

func (tFile *TargetFile) remove(license string) error {
	if !tFile.HasData() {
		return fmt.Errorf("file is empty: %s", tFile.path)
	}

	if !strings.Contains(tFile.data[0], license) {
		return nil
	}

	truncatedData := tFile.data[2:]

	err := tFile.writeToFile(tFile.path, truncatedData)
	return err
}

func (tFile *TargetFile) writeToFile(path string, data []string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range data {
		_, err := writer.WriteString(fmt.Sprintf("%s\n", line))
		if err != nil {
			return err
		}
	}

	err = writer.Flush()
	return err
}
