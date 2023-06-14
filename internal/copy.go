package internal

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const dirPermissions = 0755
const filePermissions = 0644

func AddGoPathAndDir(inputPackagePath, inputDirPath *string) (fullPath string, err error) {
	// Remove the "./" from the path. This is necessary to get the correct path.
	// Example: "./github.com" -> "github.com"
	*inputPackagePath = filepath.Clean(*inputPackagePath)

	// Get the GOPATH
	gopath := os.Getenv("GOPATH")
	// If GOPATH is not defined, return
	if gopath == "" {
		fmt.Println("GOPATH not defined.")
		return "", errors.New("GOPATH not defined")
	}
	// Join the GOPATH with the path
	fullPath = filepath.Join(gopath, "pkg", "mod", *inputPackagePath, *inputDirPath)
	return fullPath, nil
}

func CheckIfPathExistsAndIsDir(path string) error {
	// Check if the path exists
	pathInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s not exists", path)
	} else if !pathInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	} else if err != nil {
		return err
	}
	return nil
}

func CopyDirectory(srcDirAbs, destDirAbs string) error {
	// Create the destination directory as a subdirectory of the specified destination directory
	destDirAbs = filepath.Join(destDirAbs, filepath.Base(srcDirAbs))

	err := os.MkdirAll(destDirAbs, dirPermissions)
	if err != nil {
		return err
	}

	// Walk the source directory recursively and copy each file and directory
	err = filepath.Walk(srcDirAbs, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDirAbs, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDirAbs, relPath)
		if info.IsDir() {
			// Create the directory in the destination if it does not exist
			return os.MkdirAll(destPath, dirPermissions)
		} else {
			// Copy the file to the destination
			err := copyFileContents(path, destPath)
			if err != nil {
				return err
			}
		}

		return nil
	})
	return err
}

func copyFileContents(src, dst string) error {
	// Open the source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(srcFile)

	// Open the destination file for writing
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(dstFile)

	// Copy the contents of the source file into the destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Set the mode of the destination file to match the source file
	err = os.Chmod(dst, filePermissions)
	if err != nil {
		return err
	}

	return nil
}
