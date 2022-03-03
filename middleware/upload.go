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
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-the-way/anoweb/context"
)

type upload struct {
	Root       string
	Size       int
	Mimes      []string
	Extensions []string
	Domain     string
	Prefix     string
	DateDir    bool
}

// Upload return upload
func Upload() *upload {
	return &upload{
		Root:       os.TempDir(),
		Size:       100 * 1024 * 1024,
		Mimes:      []string{"text/plain", "image/jpeg", "image/jpg", "image/png", "image/gif", "application/octet-stream"},
		Extensions: []string{".txt", ".jpg", ".png", ".gif", ".xlsx"},
		Domain:     "http://localhost:9494",
		Prefix:     "/download/file",
		DateDir:    true,
	}
}

func (u *upload) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		defer removeTmpFiles(ctx)
		type jd struct {
			Msg  string `json:"msg"`
			Code int    `json:"code"`
		}
		type jdd struct {
			Msg  string                 `json:"msg"`
			Code int                    `json:"code"`
			Data map[string]interface{} `json:"data"`
		}
		getJson := func(msg string, code int) *jd {
			return &jd{
				Msg:  msg,
				Code: code,
			}
		}
		getJsonData := func(data map[string]interface{}) *jdd {
			return &jdd{
				Msg:  "success",
				Code: 0,
				Data: data,
			}
		}
		var err error
		ctx.ParseMultipart(int64(u.Size))
		files := ctx.MultipartFiles("file")
		for _, file := range files {
			err = verifyFile(file, u)
			if err != nil {
				ctx.JSON(getJson(err.Error(), 1))
				return
			}
		}
		uploadFiles := make([]uFile, 0)
		saveFiles := make([]string, 0)
		dateDir := ""
		if u.DateDir {
			dateDir = time.Now().Format("20060102")
		}
		parentPath := filepath.Join(u.Root, dateDir)
		_ = os.MkdirAll(parentPath, 0700)
		for _, file := range files {
			extIndex := strings.LastIndexByte(file.FileHeader.Filename, '.')
			ext := file.FileHeader.Filename[extIndex+1:]
			rand.Seed(time.Now().UnixNano())
			saveFile := fmt.Sprintf("%x.%s", rand.Int63(), ext)
			saveFilePath := filepath.Join(parentPath, saveFile)
			fileUrl := fmt.Sprintf("%s%s?file=%s", u.Domain, u.Prefix, saveFile)
			uploadFiles = append(uploadFiles, uFile{
				File: saveFile,
				Url:  fileUrl,
			})
			file.Copy(saveFilePath)
			saveFiles = append(saveFiles, saveFilePath)
		}

		ctx.JSON(getJsonData(map[string]interface{}{
			"total": len(files),
			"files": uploadFiles,
		}))

	}
}

func (u *upload) Create(file string, buf []byte) error {
	dateDir := ""
	if u.DateDir {
		dateDir = time.Now().Format("20060102")
	}
	fileName := filepath.Join(u.Root, dateDir, file)
	_ = os.MkdirAll(filepath.Dir(fileName), 0700)
	return os.WriteFile(fileName, buf, 0700)
}

func (u *upload) Remove(file string) error {
	dateDir := ""
	if u.DateDir {
		dateDir = time.Now().Format("20060102")
	}
	fileName := filepath.Join(u.Root, dateDir, file)
	return os.Remove(fileName)
}

type uFile struct {
	File string `json:"file"`
	Url  string `json:"url"`
}

func verifyFile(file *context.MultipartFile, u *upload) error {
	err := verifySize(file, u)
	if err != nil {
		return err
	}

	err = verifyMime(file, u)
	if err != nil {
		return err
	}

	err = verifyExt(file, u)
	if err != nil {
		return err
	}

	return nil
}

func verifySize(file *context.MultipartFile, u *upload) error {
	if file.FileHeader.Size > int64(u.Size) {
		return fmt.Errorf("the file size exceed limit[max:%d, current:%d]", u.Size, file.FileHeader.Size)
	}
	return nil
}

func verifyMime(file *context.MultipartFile, u *upload) error {
	ms := strings.Join(u.Mimes, "|")
	mimeAll := fmt.Sprintf("|%s|", ms)
	if !strings.Contains(mimeAll, fmt.Sprintf("|%s|", file.ContentType)) {
		return fmt.Errorf("the file mime type not supported[supports:%s, current:%s]", ms, file.ContentType)
	}
	return nil
}

func verifyExt(file *context.MultipartFile, u *upload) error {
	es := strings.Join(u.Extensions, "|")
	extAll := fmt.Sprintf("|%s|", es)
	extIndex := strings.LastIndexByte(file.FileHeader.Filename, '.')
	if extIndex == -1 {
		return fmt.Errorf("the file ext not supported[supports:%s, current:%s]", es, "")
	}
	ext := file.FileHeader.Filename[extIndex:]
	if !strings.Contains(extAll, fmt.Sprintf("|%s|", ext)) {
		return fmt.Errorf("the file ext not supported[supports:%s, current:%s]", es, ext)
	}
	return nil
}

func removeTmpFiles(ctx *context.Context) {
	if form := ctx.Request.MultipartForm; form != nil {
		_ = form.RemoveAll()
	}
}
