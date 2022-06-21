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
	"errors"
	"net/http"
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"

	"github.com/stretchr/testify/require"
)

func TestRecovery(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	pass := false
	var panicObj interface{}
	r := Recovery(func(ctx *context.Context) {
		defer func() {
			if re := recover(); re != nil {
				pass = true
				panicObj = re
			}
		}()
		ctx.Chain()
	})
	ctx := context.New()
	ctx.Allocate(req, &config.Template{})
	ctx.Add(r.Handler())
	ctx.Add(func(ctx *context.Context) { panic(`try panic`) })
	ctx.Chain()
	require.Equal(t, true, pass)
	require.Equal(t, `try panic`, panicObj)
}

func TestRecoveryDefault(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	r := Recovery()
	{
		ctx := context.New()
		ctx.Allocate(req, &config.Template{})
		ctx.Add(r.Handler())
		ctx.Add(func(ctx *context.Context) { panic(`try panic`) })
		ctx.Chain()
	}
	{
		ctx := context.New()
		ctx.Allocate(req, &config.Template{})
		ctx.Add(r.Handler())
		ctx.Add(func(ctx *context.Context) { panic(errors.New(`try panic`)) })
		ctx.Chain()
	}
	{
		ctx := context.New()
		ctx.Allocate(req, &config.Template{})
		ctx.Add(r.Handler())
		ctx.Add(func(ctx *context.Context) { panic(100) })
		ctx.Chain()
	}
}
