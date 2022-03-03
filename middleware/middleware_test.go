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
	"testing"

	"github.com/go-the-way/anoweb/context"
	"github.com/stretchr/testify/require"
)

type _middleware struct {
}

func (_ *_middleware) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {}
}

func TestMiddleware(t *testing.T) {
	var m Middleware
	m = &_middleware{}
	require.NotNil(t, &m)
}
