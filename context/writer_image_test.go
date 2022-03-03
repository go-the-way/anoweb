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
	"io/ioutil"
	"os"
	"testing"

	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

func TestImage(t *testing.T) {
	ctx := New(buildReq(""), nil)
	ctx.Image([]byte{100, 200, 150})
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.JPG, r.ContentType)
}

func TestImageFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte{100, 200, 150}, 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New(buildReq(""), nil)
	ctx.ImageFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.JPG, r.ContentType)
}

func TestICO(t *testing.T) {
	ctx := New(buildReq(""), nil)
	ctx.ICO([]byte{100, 200, 150})
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.ICO, r.ContentType)
}

func TestICOFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte{100, 200, 150}, 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New(buildReq(""), nil)
	ctx.ICOFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.ICO, r.ContentType)
}

func TestBMP(t *testing.T) {
	ctx := New(buildReq(""), nil)
	ctx.BMP([]byte{100, 200, 150})
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.BMP, r.ContentType)
}

func TestBMPFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte{100, 200, 150}, 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New(buildReq(""), nil)
	ctx.BMPFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.BMP, r.ContentType)
}

func TestJPG(t *testing.T) {
	ctx := New(buildReq(""), nil)
	ctx.JPG([]byte{100, 200, 150})
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.JPG, r.ContentType)
}

func TestJPGFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte{100, 200, 150}, 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New(buildReq(""), nil)
	ctx.JPGFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.JPG, r.ContentType)
}

func TestJPEG(t *testing.T) {
	ctx := New(buildReq(""), nil)
	ctx.JPEG([]byte{100, 200, 150})
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.JPG, r.ContentType)
}

func TestJPEGFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte{100, 200, 150}, 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New(buildReq(""), nil)
	ctx.JPEGFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.JPG, r.ContentType)
}

func TestPNG(t *testing.T) {
	ctx := New(buildReq(""), nil)
	ctx.PNG([]byte{100, 200, 150})
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.PNG, r.ContentType)
}

func TestPNGFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte{100, 200, 150}, 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New(buildReq(""), nil)
	ctx.PNGFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.PNG, r.ContentType)
}

func TestGIF(t *testing.T) {
	ctx := New(buildReq(""), nil)
	ctx.GIF([]byte{100, 200, 150})
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.GIF, r.ContentType)
}

func TestGIFFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte{100, 200, 150}, 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New(buildReq(""), nil)
	ctx.GIFFile("file.txt")
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.GIF, r.ContentType)
}
