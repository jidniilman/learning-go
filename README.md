# Learning Go
This repository is a collection of my personal Go source code while doing the learning process.

# Materials
Below is the list of materials for my learning

## 1. Programming with Google Go Specialization
- Code : 1_pwggs
- Publisher : University of California, Irvine
- Provider : Coursera
- Link : https://www.coursera.org/specializations/google-golang
- Learning Start : March 30, 2022
- Learning End : July 29, 2022

## 2. LeetCode
- Code : leetcode

# Commands

## go build
Compile the program
```
go build main.go
```

## go doc
Show the documentation
```
go doc <pkg>
```

## go fmt
Format source code files
```
gofmt -w .       # formats files in current directory and all sub-directories
go fmt ./...     # similar to previous
go fmt .         # formats files in current package
gofmt -w foo/bar # formats files in directory $PWD/foo/bar and sub-dirs
go fmt foo/bar   # formats files in directory $GOPATH/src/foo/bar
gofmt -w         # error, no file or directory specified
go fmt           # formats files in current package
```

## go get
Download packages and install
```
go get <pkg>
go get -u <pkg> # also update all its dependencies
```

## go list
List all installed packages
```
go list
```

## go run
Compile and run the executables
```
go run main.go
```

## go test
Run tests file
```
go test
```