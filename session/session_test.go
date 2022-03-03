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
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSession(t *testing.T) {
	s := &_session{"SESSION", time.Now(), map[string]interface{}{}}

	// test for Id
	require.Equal(t, "SESSION", s.Id())

	// test for Renew
	s.Renew(time.Minute)
	require.Equal(t, false, s.Invalidated())

	// test for Renew
	s.Renew(-time.Minute)
	require.Equal(t, true, s.Invalidated())

	// test for Invalidate
	s.Invalidate()
	require.Equal(t, true, s.Invalidated())

	// test for Invalidated
	require.Equal(t, true, s.Invalidated())

	// test for Get
	require.Equal(t, nil, s.Get("banana"))

	// test for GetAll
	require.Equal(t, map[string]interface{}{}, s.GetAll())

	// test for Set
	s.Set("banana", "1")
	require.Equal(t, "1", s.Get("banana"))

	// test for SetAll with no flush
	s.SetAll(map[string]interface{}{"orange": "2"}, false)
	require.Equal(t, "1", s.Get("banana"))
	require.Equal(t, "2", s.Get("orange"))

	// test for SetAll with flush
	s.SetAll(map[string]interface{}{"orange": "2"}, true)
	require.Equal(t, nil, s.Get("banana"))
	require.Equal(t, "2", s.Get("orange"))

	// test for Del
	s.Del("banana")
	require.Equal(t, nil, s.Get("banana"))

	// test for Clear
	s.Clear()
	require.Equal(t, 0, len(s.data))

}

type _session struct {
	id       string
	lifeTime time.Time
	data     map[string]interface{}
}

func (s *_session) Id() string {
	return s.id
}

func (s *_session) Renew(lifeTime time.Duration) {
	s.lifeTime = time.Now().Add(lifeTime)
}

func (s *_session) Invalidate() {
	s.lifeTime = time.Now().Add(-time.Minute)
}

func (s *_session) Invalidated() bool {
	return time.Now().After(s.lifeTime)
}

func (s *_session) Get(name string) interface{} {
	return s.data[name]
}

func (s *_session) GetAll() map[string]interface{} {
	return s.data
}

func (s *_session) Set(name string, val interface{}) {
	s.data[name] = val
}

func (s *_session) SetAll(data map[string]interface{}, flush bool) {
	if flush {
		s.data = data
		return
	}
	for k, v := range data {
		s.data[k] = v
	}
}

func (s *_session) Del(name string) {
	delete(s.data, name)
}

func (s *_session) Clear() {
	for k := range s.data {
		delete(s.data, k)
	}
}
