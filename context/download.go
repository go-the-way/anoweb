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
	"net/http"
	"net/url"
	"os"

	"github.com/go-the-way/anoweb/headers"
	"github.com/go-the-way/anoweb/mime"
)

// Download Response download
func (ctx *Context) Download(file, filename string) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fn := url.PathEscape(filename)
	ctx.Write(Builder().Data(bytes).ContentType(mime.BINARY).Header(http.Header{
		headers.ContentDisposition: []string{fmt.Sprintf("attachment; filename=\"%s\"", fn)},
	}).Build())
}
