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

package memory

import (
	"time"

	se "github.com/go-the-way/anoweb/session"
)

type session struct {
	id          string
	expiresTime time.Time
	invalidated bool
	attributes  map[string]interface{}
}

func newSession(id string, valid time.Duration) se.Session {
	return &session{
		id:          id,
		expiresTime: time.Now().Add(valid),
		attributes:  make(map[string]interface{}),
	}
}

// Id return session id
func (s *session) Id() string {
	return s.id
}

// Renew session
func (s *session) Renew(lifeTime time.Duration) {
	s.expiresTime = time.Now().Add(lifeTime)
}

// Invalidate session
func (s *session) Invalidate() {
	s.invalidated = true
}

// Invalidated session
func (s *session) Invalidated() bool {
	return s.invalidated
}

func (s *session) GetAll() map[string]interface{} {
	return s.attributes
}

func (s *session) Get(name string) interface{} {
	val, have := s.attributes[name]
	if have {
		return val
	}
	return nil
}

func (s *session) Set(name string, val interface{}) {
	s.attributes[name] = val
}

func (s *session) SetAll(data map[string]interface{}, flush bool) {
	if data == nil {
		return
	}

	if flush {
		s.attributes = data
		return
	}

	for k, v := range data {
		s.Set(k, v)
	}
}

// Del named val from session
func (s *session) Del(name string) {
	delete(s.attributes, name)
}

func (s *session) Clear() {
	for k := range s.attributes {
		delete(s.attributes, k)
	}
}
