# anoweb

```
                                    _
                                   | |
 _____  ____    ___   _ _ _  _____ | |__
(____ ||  _ \  / _ \ | | | || ___ ||  _ \ 
/ ___ || | | || |_| || | | || ____|| |_) )
\_____||_| |_| \___/  \___/ |_____)|____/ 

::anoweb:: 

The lightweight and powerful web framework using the new way for Go. Another go the way.

{{ Version @VER }}

{{ Powered by go-the-way }}

{{ https://github.com/go-the-way/anoweb }}

```

[![CircleCI](https://circleci.com/gh/go-the-way/anoweb/tree/main.svg?style=shield)](https://circleci.com/gh/go-the-way/anoweb/tree/main)
[![codecov](https://codecov.io/gh/go-the-way/anoweb/branch/main/graph/badge.svg?token=8MAR3J959H)](https://codecov.io/gh/go-the-way/anoweb)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-the-way/anoweb)](https://goreportcard.com/report/github.com/go-the-way/anoweb)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-the-way/anoweb?status.svg)](https://pkg.go.dev/github.com/go-the-way/anoweb?tab=doc)
[![Release](https://img.shields.io/github/release/go-the-way/anoweb.svg?style=flat-square)](https://github.com/go-the-way/anoweb/releases)

## Overview

- [Documents](https://github.com/go-the-way/anoweb/wikis)
- [Features](#Features)
- [Install](#Install)
- [Quickstart](#Quickstart)
- [Releases](https://github.com/go-the-way/anoweb/releases)
- [Todo](https://github.com/go-the-way/anoweb/blob/main/TODO.md)
- [Pull Request](https://github.com/go-the-way/anoweb/pulls)
- [Issues](https://github.com/go-the-way/anoweb/issues)
- [Thanks](#thanks)

## Features

- Pure native, no third dependencies
- Basic & Variables & Group router
- REST-ful controllers
- Binding & validation
- Middleware supports
- Session supports
- Rich Response supports

## Install

```
require github.com/go-the-way/anoweb latest
```

## Quickstart

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	
	"github.com/go-the-way/anoweb"
	"github.com/go-the-way/anoweb/context"
)

func main() {
	go func() {
		time.AfterFunc(time.Second, func() {
			response, _ := http.Get("http://localhost:9494")
			resp, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(resp))
		})
	}()
	anoweb.Default.Get("/", func(ctx *context.Context) {
		ctx.Text("Hello world")
	}).Run()
}
```

### Thanks
* [JetBrains OpenSource](https://jb.gg/OpenSource)
