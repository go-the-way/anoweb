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

package context

import (
	"embed"

	"github.com/go-the-way/anoweb/mime"
)

// TextFSFile Response text FS file
func (ctx *Context) TextFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.TEXT)
}

// JSONFSFile Response JSON FS file
func (ctx *Context) JSONFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.JSON)
}

// XMLFSFile Response XML FS file
func (ctx *Context) XMLFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.XML)
}

// ImageFSFile Response image FS file
func (ctx *Context) ImageFSFile(fs *embed.FS, name string) {
	ctx.JPEGFSFile(fs, name)
}

// ICOFSFile Response ico FS file
func (ctx *Context) ICOFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.ICO)
}

// BMPFSFile Response bmp FS file
func (ctx *Context) BMPFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.BMP)
}

// JPGFSFile Response jpg FS file
func (ctx *Context) JPGFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.JPG)
}

// JPEGFSFile Response jpeg FS file
func (ctx *Context) JPEGFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.JPG)
}

// PNGFSFile Response png FS file
func (ctx *Context) PNGFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.PNG)
}

// GIFFSFile Response gif FS file
func (ctx *Context) GIFFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.GIF)
}

// HTMLFSFile Response html FS file
func (ctx *Context) HTMLFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.HTML)
}

// CSSFSFile Response css FS file
func (ctx *Context) CSSFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.CSS)
}

// JSFSFile Response js FS file
func (ctx *Context) JSFSFile(fs *embed.FS, name string) {
	ctx.FSFile(fs, name, mime.JS)
}

// FSFile Response file
func (ctx *Context) FSFile(fs *embed.FS, name, contentType string) {
	bytes, err := fs.ReadFile(name)
	if err != nil {
		panic(err)
	}
	ctx.Binary(bytes, contentType)
}
