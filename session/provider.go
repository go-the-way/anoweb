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

package session

import "net/http"

// Provider interface
type Provider interface {
	// CookieName session cookie name
	CookieName() string
	// Exists session
	Exists(id string) bool
	// GetId session id
	GetId(r *http.Request) string
	// Del session
	Del(id string)
	// Get session id
	Get(id string) Session
	// GetAll session
	GetAll() map[string]Session
	// Clear sessions
	Clear()
	// New return session
	New(config *Config, listener *Listener) Session
	// Refresh session
	Refresh(session Session, config *Config, listener *Listener)
	// Clean session
	Clean(config *Config, listener *Listener)
}
