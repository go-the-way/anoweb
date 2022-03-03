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
	"log"
	"os"
	"time"

	"github.com/go-the-way/anoweb/context"
)

type logger struct {
	logger  *log.Logger
	handler func(ctx *context.Context)
}

// Logger return logger
func Logger() *logger {
	return &logger{
		logger: log.New(os.Stdout, "[Logger] ", log.LstdFlags),
	}
}

// Handler implements
func (l *logger) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		start := time.Now().UnixNano()
		ctx.Chain()
		r := ctx.Response
		end := time.Now().UnixNano()
		req := ctx.Request
		l.logger.Println(fmt.Sprintf("\"%s\" \"%s %s %s\" %d %d %dms \"%s\"", req.RemoteAddr, req.Method, req.URL.RequestURI(), req.Proto, r.Status, len(r.Data), (end-start)/1000.0/1000.0, req.UserAgent()))
	}
}
