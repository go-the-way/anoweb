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
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProvider(t *testing.T) {
	p := &_provider{"SID", map[string]Session{}}

	// test for CookieName
	require.Equal(t, "SID", p.CookieName())

	// test for Exists
	require.Equal(t, false, p.Exists("AOO"))

	// test for GetId
	req, _ := http.NewRequest("", "", nil)
	req.Header.Set(p.CookieName(), "hello")
	require.Equal(t, "hello", p.GetId(req))

	// test for Del
	p.sessions["abc"] = &_session{}
	p.Del("abc")
	require.Nil(t, p.Get("abc"))

	// test for Get
	p.sessions["abc"] = &_session{}
	require.NotNil(t, p.Get("abc"))

	// test for GetAll
	p.sessions["abc"] = &_session{}
	require.NotNil(t, p.GetAll())

	// test for Clear
	p.Clear()
	require.Equal(t, 0, len(p.sessions))

	// test for New
	require.NotNil(t, p.New(nil, nil))

	// test for Refresh
	p.Refresh(nil, nil, nil)
	require.Equal(t, 1, _refreshed)

	// test for Clean
	p.Clean(nil, nil)
	require.Equal(t, 1, _clean)
}

type _provider struct {
	cookieName string
	sessions   map[string]Session
}

func (p *_provider) CookieName() string {
	return p.cookieName
}

func (p *_provider) Exists(id string) bool {
	return p.sessions[id] != nil
}

func (p *_provider) GetId(r *http.Request) string {
	return r.Header.Get(p.cookieName)
}

func (p *_provider) Del(id string) {
	delete(p.sessions, id)
}

func (p *_provider) Get(id string) Session {
	return p.sessions[id]
}

func (p *_provider) GetAll() map[string]Session {
	return p.sessions
}

func (p *_provider) Clear() {
	for k := range p.sessions {
		delete(p.sessions, k)
	}
}

func (p *_provider) New(_ *Config, _ *Listener) Session {
	return &_session{}
}

var (
	_refreshed = 0
	_clean     = 0
)

func (p *_provider) Refresh(_ Session, _ *Config, _ *Listener) {
	_refreshed = 1
}

func (p *_provider) Clean(_ *Config, _ *Listener) {
	_clean = 1
}
