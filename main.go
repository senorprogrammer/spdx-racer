package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	filesFlag   string
	licenseFlag string
)

func init() {
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
