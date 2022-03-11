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

func TestFile(t *testing.T) {
	_ = ioutil.WriteFile("file.txt", []byte(`hello world`), 0700)
	defer func() { _ = os.Remove("file.txt") }()
	ctx := New(buildReq(""), nil)
	ctx.File("file.txt", mime.BINARY)
	r := ctx.Response
	require.Equal(t, []byte(`hello world`), r.Data)
	require.Equal(t, mime.BINARY, r.ContentType)
}

func TestFilePanic(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	ctx := New(buildReq(""), nil)
	ctx.File("file.txt", mime.BINARY)
}

func TestBinary(t *testing.T) {
	ctx := New(buildReq(""), nil)
	ctx.Binary([]byte{100, 200, 150}, mime.BINARY)
	r := ctx.Response
	require.Equal(t, []byte{100, 200, 150}, r.Data)
	require.Equal(t, mime.BINARY, r.ContentType)
}