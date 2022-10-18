<div align="center">

# Barad-dûr

</div>

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/baraddur.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/baraddur/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/baraddur/actions/workflows/go.yml) [![CodeQL](https://github.com/gowizzard/baraddur/actions/workflows/codeql.yml/badge.svg)](https://github.com/gowizzard/baraddur/actions/workflows/codeql.yml) [![CompVer](https://github.com/gowizzard/baraddur/actions/workflows/compver.yml/badge.svg)](https://github.com/gowizzard/baraddur/actions/workflows/compver.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/gowizzard/baraddur.svg)](https://pkg.go.dev/github.com/gowizzard/baraddur) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/baraddur)](https://goreportcard.com/report/github.com/gowizzard/baraddur) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/baraddur)](https://github.com/gowizzard/baraddur/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/baraddur)](https://github.com/gowizzard/baraddur/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/baraddur)](https://github.com/gowizzard/baraddur/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/baraddur)](https://github.com/gowizzard/baraddur/blob/master/LICENSE)

This is a small and very lightweight file watcher that can also be used in docker projects. With this file watcher you can watch different files, set the interval of the check and define in a function what should happen with a new mod time of the file.

## Install

First you have to install the package. You can do this as follows:

```console
go get github.com/gowizzard/baraddur
```

## How to use

Here you can find an example of how to use the function. You can define multiple files. For each file you can define the path, the interval of the check and the function to be executed in case of a new modification. It is also possible to define what should happen if an error occurs.

```go
c := baraddur.Config{
    Files: []baraddur.File{
        {
            Path:     "config.go",
            Interval: 1 * time.Second,
            Fault: func(err error) {
                log.Fatalln(err)
            },
            Execute: func() {
                log.Println("Update!")
            },
        },
    },
}

c.Watch()
```

## Special thanks

Thanks to [JetBrains](https://github.com/JetBrains) for supporting me with this and other [open source projects](https://www.jetbrains.com/community/opensource/#support).