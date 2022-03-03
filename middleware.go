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

	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/middleware"
	"github.com/go-the-way/anoweb/session"
)

type defaultMWState struct {
	header bool
	logger bool

	recovery         bool
	recoveryHandler  func(ctx *context.Context)
	recoveryCodeName string
	recoveryCodeVal  int
	recoveryMsgName  string

	session         bool
	sessionProvider session.Provider
	sessionConfig   *session.Config
	sessionListener *session.Listener

	static       bool
	staticCache  bool
	staticRoot   string
	staticPrefix string

	favicon     bool
	faviconFile string
	faviconFS   *embed.FS
}

// Use Middlewares
func (a *App) Use(middlewares ...middleware.Middleware) *App {
	a.middlewares = append(a.middlewares, middlewares...)
	return a
}

// UseSession Use Session Middleware
func (a *App) UseSession(provider session.Provider, config *session.Config, listener *session.Listener) *App {
	a.defaultMWState.session = true
	a.defaultMWState.sessionProvider = provider
	a.defaultMWState.sessionConfig = config
	a.defaultMWState.sessionListener = listener
	return a
}

// UseRecovery Use Recovery Middleware
func (a *App) UseRecovery() *App {
	a.defaultMWState.recovery = true
	return a
}

// UseLogger Use Logger Middleware
func (a *App) UseLogger() *App {
	a.defaultMWState.logger = true
	return a
}

// RecoveryConfig Sets Recovery config
func (a *App) RecoveryConfig(codeName string, codeVal int, msgName string) *App {
	a.defaultMWState.recoveryCodeName = codeName
	a.defaultMWState.recoveryCodeVal = codeVal
	a.defaultMWState.recoveryMsgName = msgName
	return a
}

// RecoveryHandler Sets Recovery handler
func (a *App) RecoveryHandler(handlers ...func(ctx *context.Context)) *App {
	if len(handlers) > 0 {
		a.defaultMWState.recoveryHandler = handlers[0]
	}
	return a
}

// UseStatic Use Static Resources Middleware
func (a *App) UseStatic(cache bool, root, prefix string) *App {
	a.defaultMWState.static = true
	a.defaultMWState.staticCache = cache
	a.defaultMWState.staticRoot = root
	a.defaultMWState.staticPrefix = prefix
	return a
}

// UseFavicon Use Favicon Middleware
func (a *App) UseFavicon() *App {
	a.defaultMWState.favicon = true
	return a
}

// FaviconFile Use Favicon File
func (a *App) FaviconFile(file string) *App {
	a.defaultMWState.faviconFile = file
	return a
}

// FaviconFS Use Favicon FS
func (a *App) FaviconFS(fs *embed.FS) *App {
	a.defaultMWState.faviconFS = fs
	return a
}

func (a *App) useDefaultMWs() *App {

	if a.defaultMWState.session {
		a.middlewares[0] = middleware.Session(a.defaultMWState.sessionProvider, a.defaultMWState.sessionConfig, a.defaultMWState.sessionListener)
	}

	if a.defaultMWState.header {
		a.middlewares[1] = middleware.Header()
	}

	if a.defaultMWState.logger {
		a.middlewares[2] = middleware.Logger()
	}

	if a.defaultMWState.recovery {
		if a.defaultMWState.recoveryMsgName != "" {
			a.middlewares[3] = middleware.RecoveryWithConfig(a.defaultMWState.recoveryCodeName, a.defaultMWState.recoveryCodeVal, a.defaultMWState.recoveryMsgName, a.defaultMWState.recoveryHandler)
		} else {
			a.middlewares[3] = middleware.Recovery(a.defaultMWState.recoveryHandler)
		}
	}

	if a.defaultMWState.static {
		a.middlewares = append(a.middlewares, middleware.Static(a.defaultMWState.staticCache, a.defaultMWState.staticRoot, a.defaultMWState.staticPrefix))
	}

	return a
}

// Middlewares Filter Middlewares
func (a *App) Middlewares() []middleware.Middleware {
	middlewares := make([]middleware.Middleware, 0)
	for _, m := range a.middlewares {
		if m != nil {
			middlewares = append(middlewares, m)
		}
	}
	return middlewares
}
