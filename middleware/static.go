// Copyright 2022 anoweb Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/mime"
	"github.com/go-the-way/anoweb/util"
)

type static struct {
	Cache      bool
	Prefix     string
	Caches     map[string][]byte
	MimeCaches map[string]string
	Root       string
	Mimes      map[string]string
}

// Static new static
func Static(cache bool, root, prefix string) *static {
	root, _ = filepath.Abs(root)
	if prefix == "" {
		prefix = "/static"
	}
	st := &static{
		Cache:      cache,
		Prefix:     prefix,
		Caches:     make(map[string][]byte, 0),
		MimeCaches: make(map[string]string, 0),
		Root:       root,
		Mimes:      defaultMimes(),
	}
	return st
}

func defaultMimes() map[string]string {
	return map[string]string{

		"jpg":  mime.JPG,
		"jpeg": mime.JPG,
		"png":  mime.PNG,
		"gif":  mime.GIF,
		"bmp":  mime.BMP,
		"ico":  mime.ICO,

		"txt":  mime.TEXT,
		"css":  mime.CSS,
		"js":   mime.JS,
		"json": mime.JSON,
		"xml":  mime.XML,
		"html": mime.HTML,

		"zip": mime.ZIP,
		"7z":  mime.ZIP7,
		"tar": mime.TAR,
		"gz":  mime.GZIP,
		"tgz": mime.TGZ,
		"rar": mime.RAR,

		"xls":  mime.XLS,
		"xlsx": mime.XLSX,
		"doc":  mime.DOC,
		"docx": mime.DOCX,
		"ppt":  mime.PPT,
		"pptx": mime.PPTX,

		"": mime.BINARY,
	}
}

// Handler implements
func (s *static) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		urlPath := strings.TrimPrefix(util.ReBuildPath(ctx.Request.URL.Path), s.Prefix)
		realPath := filepath.Join(s.Root, urlPath)
		ext := ""
		extPos := strings.LastIndexByte(urlPath, '.')
		if extPos != -1 {
			ext = urlPath[extPos+1:]
		}
		var (
			buffer   []byte
			have     bool
			mimeType string
		)

		if mimeType, have = s.Mimes[ext]; !have {
			mimeType = mime.BINARY
		}

		if buffer, have = s.Caches[urlPath]; !have {
			bytes, err := ioutil.ReadFile(realPath)
			if err != nil {
				ctx.Chain()
				return
			}
			buffer = bytes
			if s.Cache {
				s.Caches[urlPath] = buffer
			}
		}
		ctx.Write(context.Builder().Data(buffer).ContentType(mimeType).Build())
	}
}

// Add mime
func (s *static) Add(ext, mime string) *static {
	s.Mimes[ext] = mime
	return s
}

// Adds mimes
func (s *static) Adds(m map[string]string) *static {
	if m != nil {
		for k, v := range m {
			s.Add(k, v)
		}
	}
	return s
}
