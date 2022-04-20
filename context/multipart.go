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
	"io"
	"mime/multipart"
	"os"

	"github.com/go-the-way/anoweb/headers"
)

// MultipartFile struct
type MultipartFile struct {
	// ContentType Content type
	ContentType string
	// FileHeader file headers
	FileHeader *multipart.FileHeader
}

// Open open source file in multiple
func (file *MultipartFile) Open() (f multipart.File, err error) {
	return file.FileHeader.Open()
}

// Copy source file in multiple
func (file *MultipartFile) Copy(distName string) (err error) {
	var (
		f    multipart.File
		dist *os.File
	)
	if f, err = file.Open(); err != nil {
		return err
	} else {
		if dist, err = os.Create(distName); err != nil {
			return err
		} else {
			_, _ = io.Copy(dist, f)
			_ = f.Close()
			_ = dist.Close()
		}
	}
	return err
}

// ParseMultipart parse multiple Request
func (ctx *Context) ParseMultipart(maxMemory int64) error {
	if err := ctx.Request.ParseMultipartForm(maxMemory); err != nil {
		return err
	}
	paramMap := make(map[string][]string, 0)
	if f := ctx.Request.MultipartForm; f != nil {
		for name, values := range f.Value {
			paramMap[name] = values
		}
		ctx.SetParamMap(paramMap, false)
		for name, header := range f.File {
			mfs := make([]*MultipartFile, 0)
			for _, fileHeader := range header {
				mfs = append(mfs, &MultipartFile{fileHeader.Header.Get(headers.MIME), fileHeader})
			}
			ctx.MultipartMap[name] = mfs
		}
	}
	return nil
}

// MultipartFile get multiple file
func (ctx *Context) MultipartFile(name string) *MultipartFile {
	files := ctx.MultipartFiles(name)
	if files != nil && len(files) > 0 {
		return files[0]
	}
	return nil
}

// MultipartFiles get multiple files
func (ctx *Context) MultipartFiles(name string) []*MultipartFile {
	files, have := ctx.MultipartMap[name]
	if have {
		return files
	}
	return nil
}
