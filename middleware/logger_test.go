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
	"strings"
	"testing"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"
)

func TestLogger(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	h := Logger()
	var buf strings.Builder
	h.logger.SetOutput(&buf)
	ctx := context.New()
	ctx.Allocate(req, &config.Template{})
	ctx.Add(h.Handler())
	ctx.Add(func(ctx *context.Context) {})
	ctx.Chain()
}
