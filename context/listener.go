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

import "sync"

// Listener interface
type Listener struct {
	// Created Listener
	Created func(c *Context)
	// Destroyed Listener
	Destroyed func(c *Context)
}

var (
	mu        = &sync.Mutex{}
	listeners = make([]*Listener, 0)
)

// AddListeners add listeners
func AddListeners(ls ...*Listener) {
	mu.Lock()
	defer mu.Unlock()
	listeners = append(listeners, ls...)
}

// ClearListeners clear listeners
func ClearListeners() {
	mu.Lock()
	defer mu.Unlock()
	listeners = make([]*Listener, 0)
}

func (ctx *Context) onCreated() {
	for _, listener := range listeners {
		if listener.Created != nil {
			listener.Created(ctx)
		}
	}
}

func (ctx *Context) onDestroyed() {
	for _, listener := range listeners {
		if listener.Destroyed != nil {
			listener.Destroyed(ctx)
		}
	}
}
