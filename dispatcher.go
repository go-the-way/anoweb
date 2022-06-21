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
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/headers"
	"github.com/go-the-way/anoweb/middleware"
	"net/http"
)

type dispatcher struct {
	*App
}

func (a *App) newDispatcher() *dispatcher {
	return &dispatcher{a}
}

func (d *dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d.dispatch(r, w)
}

func (d *dispatcher) dispatch(r *http.Request, w http.ResponseWriter) {
	ctx := d.ctxPool.Get().(*context.Context)
	ctx.Allocate(r, d.App.Config.Template)
	d.addChains(ctx, d.App.parsedRouters.Handler(ctx), d.App.Middlewares())
	ctx.Chain()
	d.writeDone(ctx.Response, w)
}

func (d *dispatcher) addChains(ctx *context.Context, handler func(ctx *context.Context), mws []middleware.Middleware) {
	for _, m := range mws {
		ctx.Add(m.Handler())
	}
	ctx.Add(handler)
}

func (d *dispatcher) writeDone(r *context.Response, w http.ResponseWriter) {
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	if r.ContentType != "" {
		w.Header().Set(headers.MIME, r.ContentType)
	}

	for _, cookie := range r.Cookies {
		w.Header().Add(headers.SetCookie, cookie.String())
	}

	if r.Status != 0 {
		w.WriteHeader(r.Status)
	}

	if r.Data != nil {
		_, _ = w.Write(r.Data)
	}
}
