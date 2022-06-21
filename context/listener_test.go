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
	"testing"

	"github.com/go-the-way/anoweb/config"

	"github.com/stretchr/testify/require"
)

func TestAddListeners(t *testing.T) {
	var apple int
	AddListeners(&Listener{Created: func(ctx *Context) { apple++ }},
		&Listener{Created: func(ctx *Context) { apple++ }},
		&Listener{Created: func(ctx *Context) { apple++ }})
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	require.Equal(t, 3, apple)
}

func TestListenerOnCreated(t *testing.T) {
	{
		var apple int
		ClearListeners()
		AddListeners(&Listener{Created: func(ctx *Context) { apple++ }})
		ctx := New()
		ctx.Allocate(buildReq(""), &config.Template{})
		ctx.onCreated()
		ctx.onCreated()
		ctx.onCreated()
		require.Equal(t, 4, apple)
	}
	{
		var apple int
		ClearListeners()
		AddListeners(&Listener{Created: func(ctx *Context) { apple++ }})
		New()
		New()
		New()
		New()
		require.Equal(t, 4, apple)
	}
}

func TestListenerOnDestroyed(t *testing.T) {
	{
		var apple int
		ClearListeners()
		AddListeners(&Listener{Destroyed: func(ctx *Context) { apple++ }})
		ctx := New()
		ctx.Allocate(buildReq(""), &config.Template{})
		ctx.onDestroyed()
		ctx.onDestroyed()
		ctx.onDestroyed()
		require.Equal(t, 3, apple)
	}
	{
		var apple int
		ClearListeners()
		AddListeners(&Listener{Destroyed: func(ctx *Context) { apple++ }})
		ctx := New()
		ctx.Allocate(buildReq(""), &config.Template{})
		ctx.Add(func(ctx *Context) { ctx.Chain() })
		ctx.Add(func(ctx *Context) { ctx.Chain() })
		ctx.Add(func(ctx *Context) { ctx.Chain() })
		ctx.Chain()
		require.Equal(t, 1, apple)
	}
}
