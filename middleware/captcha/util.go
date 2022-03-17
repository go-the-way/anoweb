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

package captcha

import (
	"strings"

	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/middleware"
)

// Current return current session captcha text
func Current(cc *captcha, ctx *context.Context) (str string) {
	if s := middleware.GetSession(ctx); s != nil {
		if ss := s.Get(cc.SessionKey); ss != nil {
			if str2, isOk := ss.(string); isOk {
				return str2
			}
		}
	}
	return ""
}

// Equals return true if current session captcha text equals val
func Equals(cc *captcha, ctx *context.Context, val string, ignoreCase bool) bool {
	current := Current(cc, ctx)
	if current == "" {
		return false
	}
	if val == "" {
		return false
	}
	return current == val || (ignoreCase && strings.EqualFold(current, val))
}

// Match return true if current session captcha text equals text in captcha
func Match(cc *captcha, ctx *context.Context, ignoreCase bool) bool {
	return Equals(cc, ctx, ctx.Param(cc.SessionKey), ignoreCase)
}

// Clear delete captcha text in current session
func Clear(cc *captcha, ctx *context.Context) {
	if s := middleware.GetSession(ctx); s != nil {
		s.Del(cc.SessionKey)
	}
}
