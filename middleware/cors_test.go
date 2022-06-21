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
	"net/http"
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/headers"

	"github.com/stretchr/testify/require"
)

func TestCors(t *testing.T) {
	req1, _ := http.NewRequest(http.MethodGet, "", nil)
	req2, _ := http.NewRequest(http.MethodOptions, "", nil)
	req3, _ := http.NewRequest(http.MethodHead, "", nil)
	for _, req := range []*http.Request{req1, req2, req3} {
		c := Cors()
		ctx := context.New()
		ctx.Allocate(req, &config.Template{})
		ctx.Add(c.Handler())
		ctx.Chain()
		r := ctx.Response
		require.Equal(t, []string{"GET, POST, DELETE, PUT, PATCH, HEAD, OPTIONS"}, r.Header.Values(headers.Allow))
		require.Equal(t, []string{"GET, POST, DELETE, PUT, PATCH, HEAD, OPTIONS"}, r.Header.Values(headers.AccessControlAllowMethods))
		require.Equal(t, []string{"*"}, r.Header.Values(headers.AccessControlAllowOrigin))
		require.Equal(t, "", r.Header.Get(headers.AccessControlAllowHeaders))
	}
}
