// SPDX-License-Identifier: MPL-2.0

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	deleteFlag  bool
	filesFlag   string
	licenseFlag string
)

func init() {
	flag.BoolVar(&deleteFlag, "d", false, "delete the license from each source file (short-hand)")
	flag.BoolVar(&deleteFlag, "delete", false, "delete the license from each source file")

	flag.StringVar(&filesFlag, "f", "", "specifies the file types to write the license into (short-hand)")
	flag.StringVar(&filesFlag, "files", "", "specifies the file types to write the license into")

	flag.StringVar(&licenseFlag, "l", "", "specifies the license to include in each source file (short-hand)")
	flag.StringVar(&licenseFlag, "license", "", "specifies the license to include in each source file")
}

func requiredFlags(filesFlag, licenseFlag string) {
	if filesFlag == "" {
		log.Fatal("--files is required")
	}

	if licenseFlag == "" {
		log.Fatal("--license is required")
	}
}

func splitFilesFlag(files string) []string {
	return strings.Split(files, ",")
}

func run() error {
	spdxWriter := NewSPDXWriter()
	fileTypes := splitFilesFlag(filesFlag)

	// We want to delete the license from the files
	if deleteFlag {
		err := spdxWriter.Delete(licenseFlag, fileTypes)
		return err
	}

	// We want to add the license to the files
	err := spdxWriter.Write(licenseFlag, fileTypes)
	return err
}

/* -------------------- Main -------------------- */

func main() {
	fmt.Println("SPDX Racer go!")

	flag.Parse()
	requiredFlags(licenseFlag, filesFlag)

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
