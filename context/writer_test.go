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
	"net/http"
	"testing"

	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

func TestWrite(t *testing.T) {
	{
		ctx := New()
		ctx.Allocate(buildReq(""), nil)
		ctx.Write(&Response{})
		require.Equal(t, Builder().DefaultBuild(), ctx.Response)
	}
	{
		ctx := New()
		ctx.Allocate(buildReq(""), nil)
		ctx.Write(&Response{
			Data:        []byte(`hello`),
			Header:      http.Header{"apple": {"100"}},
			Cookies:     []*http.Cookie{{Name: "coco", Value: "100"}},
			Status:      http.StatusOK,
			ContentType: mime.TEXT,
		})
		r := ctx.Response
		require.Equal(t, []byte(`hello`), r.Data)
		require.Equal(t, http.Header{"apple": {"100"}}, r.Header)
		require.Equal(t, []*http.Cookie{{Name: "coco", Value: "100"}}, r.Cookies)
		require.Equal(t, http.StatusOK, r.Status)
		require.Equal(t, mime.TEXT, r.ContentType)
	}
}
