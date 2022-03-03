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

import (
	"time"
)

// Session interface
type Session interface {
	// Id session id
	Id() string
	// Renew session
	Renew(lifeTime time.Duration)
	// Invalidate session
	Invalidate()
	// Invalidated session
	Invalidated() bool
	// Get named val
	Get(name string) interface{}
	// GetAll session's values
	GetAll() map[string]interface{}
	// Set named val
	Set(name string, val interface{})
	// SetAll values
	SetAll(data map[string]interface{}, flush bool)
	// Del named val
	Del(name string)
	// Clear session's val
	Clear()
}
