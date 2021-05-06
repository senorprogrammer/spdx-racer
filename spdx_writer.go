// SPDX-License-Identifier: MPL-2.0

package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var (
	// excludableFileNames defines naming conventions for files that should be excluded from
	// having a license entry added to them (like auto-generated files, etc.)
	excludableFileNames = []string{
		"_mock",
	}
)

// SPDXWriter is responsible for writing the license into the files
type SPDXWriter struct {
	fileTypes []string
	license   string
}

// NewSPDXWriter creates and returns an instance of SPDXWriter
func NewSPDXWriter() *SPDXWriter {
	return &SPDXWriter{}
}

/* -------------------- Exported Functions -------------------- */

// Delete deletes the specified license line from the files
func (sw *SPDXWriter) Delete(license string, fileTypes []string) error {
	sw.Set(license, fileTypes)

	currDir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = filepath.WalkDir(currDir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		isLicensable := sw.isLicensableFile(path)
		if !isLicensable {
			return nil
		}

		// This file does indeed have an extension we want to prepend a license to
		tFile := NewTargetFile(path)

		err = tFile.RemoveLicense(sw.license)
		return err
	})

	return err
}

// Set stores the data needed to operate on files
func (sw *SPDXWriter) Set(license string, fileTypes []string) {
	sw.fileTypes = fileTypes
	sw.license = license
}

// Write recursively loops over all the files in the directory looking for ones with extensions
// that are in fileTypes and writing the license string into the top of the file if it
// does not already exist
func (sw *SPDXWriter) Write(license string, fileTypes []string) error {
	sw.Set(license, fileTypes)

	currDir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = filepath.WalkDir(currDir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		isLicensable := sw.isLicensableFile(path)
		if !isLicensable {
			return nil
		}

		// This file does indeed have an extension we want to prepend a license to
		tFile := NewTargetFile(path)

		err = tFile.AddLicense(sw.license)
		return err
	})

	return err
}

/* -------------------- Unexported Functions -------------------- */

// isLicensableFile checks to see if this is the kind of file we want to prepend a license to
func (sw *SPDXWriter) isLicensableFile(path string) bool {
	if !sw.hasExtension(filepath.Ext(path)) {
		return false
	}

	if sw.hasExcludableName(filepath.Base(path)) {
		return false
	}

	return true
}

func (sw *SPDXWriter) hasExcludableName(fileName string) bool {
	for _, fn := range excludableFileNames {
		if strings.Contains(fileName, fn) {
			return true
		}
	}

	return false
}

// hasExtension returns TRUE if the incoming extension is one we're looking to operate
// against, FALSE if it is not
func (sw *SPDXWriter) hasExtension(ext string) bool {
	for _, ft := range sw.fileTypes {
		if ft == strings.TrimPrefix(ext, ".") {
			return true
		}
	}

	return false
}
