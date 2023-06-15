# Go Package Installer

![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/steficalde/go-package-installer?include_prereleases)
[![Go Report Card](https://goreportcard.com/badge/github.com/steficalde/go-package-installer)](https://goreportcard.com/report/github.com/steficalde/go-package-installer)
[![Go Reference](https://pkg.go.dev/badge/github.com/steficalde/go-package-installer.svg)](https://pkg.go.dev/github.com/steficalde/go-package-installer)


## What is it?
Go Package Installer is a tool for installing Go packages on your project.

## Installation
Get the package from the GitHub repository.
```
go get -u github.com/steficalde/go-package-installer
```

Install the package.
```
go install github.com/steficalde/go-package-installer
```

## Usage
### Command
```
go-package-installer [-i inputDir] [-o outputDir] <package>
```

### Example
```
go-package-installer -i internal -o install github.com/user/package
```

### Options

`-o` Specify the output directory.  

`-i` Specify the input package directory. 

`-h` Show the help message.  


## Permissions
The default permissions for files and directories are as follows:  
File: 644  
Directory: 755  
These permissions can be modified based on your needs.
