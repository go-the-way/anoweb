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

func TestBuilderBuild(t *testing.T) {
	b := ResponseBuilder{}
	require.Equal(t, b.r, b.Build())
}

func TestRenderBuilder(t *testing.T) {
	require.Equal(t, &ResponseBuilder{r: &Response{Status: http.StatusOK, ContentType: mime.TEXT}}, Builder())
}

func TestDefaultBuild(t *testing.T) {
	rb := Builder()
	r := rb.DefaultBuild()
	require.Equal(t, r, &Response{
		Header:      http.Header{},
		Cookies:     make([]*http.Cookie, 0),
		Status:      http.StatusOK,
		ContentType: mime.TEXT,
	})
}

func TestBuilderBuffer(t *testing.T) {
	response := Builder().Data([]byte(`hello world`)).Build()
	require.Equal(t, []byte(`hello world`), response.Data)
}

func TestBuilderHeader(t *testing.T) {
	response := Builder().Header(http.Header{"apple": []string{"100"}}).Build()
	require.Equal(t, http.Header{"apple": {"100"}}, response.Header)
}

func TestBuilderCookies(t *testing.T) {
	response := Builder().Cookies([]*http.Cookie{{Name: "apple", Value: "100"}, {Name: "orange", Value: "200"}}).Build()
	require.Equal(t, []*http.Cookie{{Name: "apple", Value: "100"}, {Name: "orange", Value: "200"}}, response.Cookies)
}

func TestBuilderCode(t *testing.T) {
	response := Builder().Status(http.StatusGatewayTimeout).Build()
	require.Equal(t, http.StatusGatewayTimeout, response.Status)
}

func TestBuilderContentType(t *testing.T) {
	response := Builder().ContentType(mime.TEXT).Build()
	require.Equal(t, mime.TEXT, response.ContentType)
}
