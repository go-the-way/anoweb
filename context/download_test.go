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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/go-the-way/anoweb/headers"
	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

func TestDownload(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte(`hello world`), 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.Download("file.txt", "file")
	r := ctx.Response
	require.Equal(t, Builder().Data([]byte(`hello world`)).Cookies([]*http.Cookie{}).ContentType(mime.BINARY).Header(http.Header{
		headers.ContentDisposition: []string{fmt.Sprintf("attachment; filename=\"%s\"", "file")},
	}).Build(), r)
}

func TestDownloadPanic(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	ctx := New()
	ctx.Allocate(buildReq(""), nil)
	ctx.Download("file.txt", "file")
}
