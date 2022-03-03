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
	"github.com/go-the-way/anoweb/mime"
)

// Image Response image
func (ctx *Context) Image(buffer []byte) {
	ctx.Binary(buffer, mime.JPG)
}

// ImageFile Response image file
func (ctx *Context) ImageFile(imageFile string) {
	ctx.File(imageFile, mime.JPG)
}

// ICO Response ico
func (ctx *Context) ICO(buffer []byte) {
	ctx.Binary(buffer, mime.ICO)
}

// ICOFile Response ico file
func (ctx *Context) ICOFile(icoFile string) {
	ctx.File(icoFile, mime.ICO)
}

// BMP Response bmp
func (ctx *Context) BMP(buffer []byte) {
	ctx.Binary(buffer, mime.BMP)
}

// BMPFile Response bmp file
func (ctx *Context) BMPFile(bmpFile string) {
	ctx.File(bmpFile, mime.BMP)
}

// JPG Response jpg
func (ctx *Context) JPG(buffer []byte) {
	ctx.Binary(buffer, mime.JPG)
}

// JPGFile Response jpg file
func (ctx *Context) JPGFile(jpgFile string) {
	ctx.File(jpgFile, mime.JPG)
}

// JPEG Response jpeg
func (ctx *Context) JPEG(buffer []byte) {
	ctx.Binary(buffer, mime.JPG)
}

// JPEGFile Response jpeg file
func (ctx *Context) JPEGFile(jpegFile string) {
	ctx.JPGFile(jpegFile)
}

// PNG Response png
func (ctx *Context) PNG(buffer []byte) {
	ctx.Binary(buffer, mime.PNG)
}

// PNGFile Response png file
func (ctx *Context) PNGFile(pngFile string) {
	ctx.File(pngFile, mime.PNG)
}

// GIF Response gif
func (ctx *Context) GIF(buffer []byte) {
	ctx.Binary(buffer, mime.GIF)
}

// GIFFile Response gif file
func (ctx *Context) GIFFile(gifFile string) {
	ctx.File(gifFile, mime.GIF)
}
