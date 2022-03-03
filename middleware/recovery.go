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
	"fmt"
	"net/http"
	"runtime"

	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/mime"
)

type recovery struct {
	handler func(ctx *context.Context)
}

// Recovery return new recovery
func Recovery(handlers ...func(ctx *context.Context)) Middleware {
	return RecoveryWithConfig("code", 500, "message", handlers...)
}

// RecoveryWithConfig return new recovery
func RecoveryWithConfig(codeName string, codeVal int, msgName string, handlers ...func(ctx *context.Context)) Middleware {
	var handler func(ctx *context.Context)
	if len(handlers) > 0 && handlers[0] != nil {
		handler = handlers[0]
	} else {
		handler = func(ctx *context.Context) {
			defer func() {
				if re := recover(); re != nil {
					_, _ = fmt.Println(fmt.Sprintf("Recovered: %v", re))
					var buf [4096]byte
					n := runtime.Stack(buf[:], false)
					fmt.Printf("%s\n", string(buf[:n]))
					message := ""
					switch re.(type) {
					case string:
						message = fmt.Sprintf("%v", re)
					case error:
						message = re.(error).Error()
					default:
						message = fmt.Sprintf("%v", re)
					}
					ctx.Write(context.Builder().
						Data([]byte(fmt.Sprintf(`{"%s":%d,"%s":"%s"}`, codeName, codeVal, msgName, message))).
						ContentType(mime.JSON).
						Status(http.StatusInternalServerError).
						Build())
				}
			}()
			ctx.Chain()
		}
	}
	return &recovery{handler}
}

// Handler implements
func (r *recovery) Handler() func(ctx *context.Context) {
	return r.handler
}
