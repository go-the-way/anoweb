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
	"github.com/stretchr/testify/require"
	"testing"
)

type _controller struct {
}

func (*_controller) Prefix() string {
	return "/_"
}

func (*_controller) Get() func(ctx *context.Context) {
	return func(ctx *context.Context) {}
}

func (*_controller) Gets() func(ctx *context.Context) {
	return func(ctx *context.Context) {}
}

func (*_controller) Post() func(ctx *context.Context) {
	return func(ctx *context.Context) {}
}

func (*_controller) Put() func(ctx *context.Context) {
	return func(ctx *context.Context) {}
}

func (*_controller) Delete() func(ctx *context.Context) {
	return func(ctx *context.Context) {}
}

func TestRestController(t *testing.T) {
	a := New()
	_c := _controller{}
	a.Controller(&_c).routeRestControllers()
	require.Equal(t, a.controllers[0], &_c)
}
