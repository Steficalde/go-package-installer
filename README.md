# Go Package Installer

Version: **0.0.1**

## What is it?
Go Package Installer is a tool for installing Go packages on your project.

## Installation
Get the package from the GitHub repository.
```
go get -u github.com/steficalde/go-package-installer
```

## Usage
### Command
```
go run install.go [-i inputDir] [-o outputDir] <package>
```

### Example
```
go run install.go -i internal -o install github.com/user/package"
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