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
	"testing"
	"time"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/headers"
	se "github.com/go-the-way/anoweb/session"

	"github.com/stretchr/testify/require"
)

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

type _provider struct {
	cookieName string
	sessions   map[string]se.Session
}

func (p *_provider) CookieName() string {
	return p.cookieName
}

func (p *_provider) Exists(id string) bool {
	return p.sessions[id] != nil
}

func (p *_provider) GetId(r *http.Request) string {
	cookie, err := r.Cookie(p.CookieName())
	if err == nil && cookie != nil {
		return cookie.Value
	}
	return ""
}

func (p *_provider) Del(id string) {
	delete(p.sessions, id)
}

func (p *_provider) Get(id string) se.Session {
	return p.sessions[id]
}

func (p *_provider) GetAll() map[string]se.Session {
	return p.sessions
}

func (p *_provider) Clear() {
	for k := range p.sessions {
		delete(p.sessions, k)
	}
}

func (p *_provider) New(_ *se.Config, _ *se.Listener) se.Session {
	return &_session{id: "fixed-session"}
}

var (
	_refreshed = 0
	_clean     = 0
)

func (p *_provider) Refresh(_ se.Session, _ *se.Config, _ *se.Listener) {
	_refreshed = 1
}

func (p *_provider) Clean(_ *se.Config, _ *se.Listener) {
	_clean = 1
}

func TestSessionName(t *testing.T) {
	require.Equal(t, sessionMWName, Session(&_provider{}, &se.Config{}, nil).Name())
}

func TestSession(t *testing.T) {
	_testFunc := func(req *http.Request, initProvider bool) {
		var currentSession se.Session
		p := _provider{cookieName: "GOSESSID", sessions: map[string]se.Session{}}
		if initProvider {
			p.sessions["fixed-session"] = &_session{id: "fixed-session"}
		}
		s := Session(&p, &se.Config{}, nil)
		ctx := context.New()
		ctx.Allocate(req, &config.Template{})
		ctx.Add(s.Handler())
		ctx.Add(func(ctx *context.Context) {
			currentSession = GetSession(ctx)
		})
		ctx.Chain()
		r := ctx.Response
		require.NotEmpty(t, r.Header[headers.SetCookie])
		require.NotNil(t, currentSession)
	}
	{
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		_testFunc(req, false)
	}
	{
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: "GOSESSID", Value: "fixed-session"})
		_testFunc(req, true)
	}
}

func TestSessionNil(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	ctx := context.New()
	ctx.Allocate(req, &config.Template{})
	ctx.Add(func(ctx *context.Context) { GetSession(ctx) })
	ctx.Chain()
}
