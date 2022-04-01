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

package anoweb

import (
	"embed"
	"net/http"
	"testing"

	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/session"
	"github.com/go-the-way/anoweb/session/memory"
)

var (
	_middlewareMessage = `middleware`
	_handlerMessage    = `hello world`
)

type (
	_middleware struct {
		beforeChain bool
		afterChain  bool
	}
)

func (m *_middleware) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		if m.beforeChain {
			ctx.Chain()
		}
		ctx.Text(_middlewareMessage)
		if m.afterChain {
			ctx.Chain()
		}
	}
}

func TestMiddlewareOnly(t *testing.T) {
	testHTTP(t, New().Use(&_middleware{}),
		// test for middleware GET Method
		&testHTTPCase{http.MethodGet, "/", "", nil, _middlewareMessage},
		// test for middleware POST Method
		&testHTTPCase{http.MethodPost, "/", "", nil, _middlewareMessage},
		// test for middleware PUT Method
		&testHTTPCase{http.MethodPut, "/", "", nil, _middlewareMessage},
		// test for middleware DELETE Method
		&testHTTPCase{http.MethodDelete, "/", "", nil, _middlewareMessage},
		// test for middleware PATCH Method
		&testHTTPCase{http.MethodPatch, "/", "", nil, _middlewareMessage})

	testHTTP(t, New().Use(&_middleware{beforeChain: true}),
		// test for middleware GET Method
		&testHTTPCase{http.MethodGet, "/", "", nil, _middlewareMessage},
		// test for middleware POST Method
		&testHTTPCase{http.MethodPost, "/", "", nil, _middlewareMessage},
		// test for middleware PUT Method
		&testHTTPCase{http.MethodPut, "/", "", nil, _middlewareMessage},
		// test for middleware DELETE Method
		&testHTTPCase{http.MethodDelete, "/", "", nil, _middlewareMessage},
		// test for middleware PATCH Method
		&testHTTPCase{http.MethodPatch, "/", "", nil, _middlewareMessage})

	testHTTP(t, New().Use(&_middleware{afterChain: true}),
		// test for middleware GET Method
		&testHTTPCase{http.MethodGet, "/", "", nil, _middlewareMessage},
		// test for middleware POST Method
		&testHTTPCase{http.MethodPost, "/", "", nil, _middlewareMessage},
		// test for middleware PUT Method
		&testHTTPCase{http.MethodPut, "/", "", nil, _middlewareMessage},
		// test for middleware DELETE Method
		&testHTTPCase{http.MethodDelete, "/", "", nil, _middlewareMessage},
		// test for middleware PATCH Method
		&testHTTPCase{http.MethodPatch, "/", "", nil, _middlewareMessage})
}

func TestMiddlewareWithRoute(t *testing.T) {
	testHTTP(t, New().Request("/", func(ctx *context.Context) {
		ctx.Text(_handlerMessage)
	}).Use(&_middleware{}),
		// test for middleware GET Method
		&testHTTPCase{http.MethodGet, "/", "", nil, _middlewareMessage},
		// test for middleware POST Method
		&testHTTPCase{http.MethodPost, "/", "", nil, _middlewareMessage},
		// test for middleware PUT Method
		&testHTTPCase{http.MethodPut, "/", "", nil, _middlewareMessage},
		// test for middleware DELETE Method
		&testHTTPCase{http.MethodDelete, "/", "", nil, _middlewareMessage},
		// test for middleware PATCH Method
		&testHTTPCase{http.MethodPatch, "/", "", nil, _middlewareMessage})

	testHTTP(t, New().Request("/", func(ctx *context.Context) {
		ctx.Text(_handlerMessage)
	}).Use(&_middleware{beforeChain: true}),
		// test for middleware GET Method
		&testHTTPCase{http.MethodGet, "/", "", nil, _middlewareMessage},
		// test for middleware POST Method
		&testHTTPCase{http.MethodPost, "/", "", nil, _middlewareMessage},
		// test for middleware PUT Method
		&testHTTPCase{http.MethodPut, "/", "", nil, _middlewareMessage},
		// test for middleware DELETE Method
		&testHTTPCase{http.MethodDelete, "/", "", nil, _middlewareMessage},
		// test for middleware PATCH Method
		&testHTTPCase{http.MethodPatch, "/", "", nil, _middlewareMessage})

	testHTTP(t, New().Request("/", func(ctx *context.Context) {
		ctx.Text(_handlerMessage)
	}).Use(&_middleware{afterChain: true}),
		// test for middleware GET Method
		&testHTTPCase{http.MethodGet, "/", "", nil, _handlerMessage},
		// test for middleware POST Method
		&testHTTPCase{http.MethodPost, "/", "", nil, _handlerMessage},
		// test for middleware PUT Method
		&testHTTPCase{http.MethodPut, "/", "", nil, _handlerMessage},
		// test for middleware DELETE Method
		&testHTTPCase{http.MethodDelete, "/", "", nil, _handlerMessage},
		// test for middleware PATCH Method
		&testHTTPCase{http.MethodPatch, "/", "", nil, _handlerMessage})
}

var fs2 embed.FS

func TestAppUses(t *testing.T) {
	{
		New().UseRecovery().UseLogger().
			UseSession(memory.Provider(), &session.Config{}, nil).
			UseStatic(true, "", "/static").
			UseFavicon().
			FaviconFile("a.ico").
			FaviconFS(&fs2).
			FaviconRoute("/favicon.ico").
			useDefaultMWs()
	}
	{
		New().UseRecovery().
			RecoveryHandler(func(ctx *context.Context) {}).
			RecoveryConfig("code", 5000, "message").
			useDefaultMWs()
	}
}
