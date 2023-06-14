package main

import (
	"flag"
	"fmt"
	cp "github.com/steficalde/go-package-installer/internal"
	"os"
	"path/filepath"
)

func main() {
	// Set the custom usage function
	flag.Usage = func() {
		fmt.Println("Usage: go run install.go [-i inputDir] [-o outputDir] <package>")
		fmt.Println("Example: go run install.go -i internal -o install github.com/user/package")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	outputDirPath := flag.String("o", "", "Specify the output directory.")
	inputDirPath := flag.String("i", "", "Specify the input package directory.")
	help := flag.Bool("h", false, "Show the help message")

	// Parse the flags
	flag.Parse()

	// Show the help message if the flag h is set
	if *help {
		flag.Usage()
		return
	}

	// Get package path from command-line arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: package path is required")
		return
	}
	// Get the package path
	inputPackagePath := args[0]

	// Create the absolute path
	fullPath, err := cp.AddGoPathAndDir(&inputPackagePath, inputDirPath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Get the current path
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error to get current path: ", err)
		return
	}

	//add outputDirPath to the current path
	if *outputDirPath != "" {
		currentPath = filepath.Join(currentPath, *outputDirPath)
	}

	// Check if the input full path exists
	err = cp.CheckIfPathExistsAndIsDir(fullPath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// Check if the output path exists
	err = cp.CheckIfPathExistsAndIsDir(currentPath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Copy the package to the current path
	err = cp.CopyDirectory(fullPath, currentPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the success message
	fmt.Println("The package has been installed in: ", currentPath)

}
