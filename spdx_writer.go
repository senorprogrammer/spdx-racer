// SPDX-License-Identifier: MPL-2.0

package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// SPDXWriter is responsible for writing the license into the files
type SPDXWriter struct {
	fileTypes []string
	license   string
}

// NewSPDXWriter creates and returns an instance of SPDXWriter
func NewSPDXWriter() *SPDXWriter {
	w := &SPDXWriter{}

	return w
}

// HasExtension returns TRUE if the incoming extension is one we're looking to operate
// against, FALSE if it is not
func (sw *SPDXWriter) HasExtension(ext string) bool {
	for _, ft := range sw.fileTypes {
		if ft == strings.TrimPrefix(ext, ".") {
			return true
		}
	}

	return false
}

// Write recursively loops over all the files in the directory looking for ones with extensions
// that are in fileTypes and writing the license string into the top of the file if it
// does not already exist
func (sw *SPDXWriter) Write(license string, fileTypes []string) error {
	sw.fileTypes = fileTypes
	sw.license = license

	currDir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = filepath.WalkDir(currDir, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Check to see if this is the kind of file we want to prepend
		// a license to
		fileExt := filepath.Ext(path)
		if !sw.HasExtension(fileExt) {
			return nil
		}

		// This file does indeed have an extension we want to prepend a license to
		tFile := NewTargetFile(path)

		err = tFile.AddLicense(sw.license)
		return err
	})

	return err
}
