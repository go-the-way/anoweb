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
	"bytes"
	"github.com/go-the-way/anoweb/config"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func buildReq(body string) *http.Request {
	var buf bytes.Buffer
	buf.WriteString(body)
	req, _ := http.NewRequest("", "", &buf)
	return req
}

func TestContextBind(t *testing.T) {
	type _model struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	m := _model{}
	ctx := New(buildReq(`{"id":100,"name":"hello world"}`), &config.Template{})
	ctx.Bind(&m)
	require.Equal(t, &_model{100, "hello world"}, &m)
}

func TestContextBindPanic(t *testing.T) {
	type _model struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	defer func() {
		if re := recover(); re != nil {
			t.Log(re)
			t.Log("test ok!")
		}
	}()
	m := _model{}
	ctx := New(buildReq(`<xml></xml>`), &config.Template{})
	ctx.Bind(&m)
}
