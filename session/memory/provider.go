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
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	se "github.com/go-the-way/anoweb/session"
)

type provider struct {
	mu       *sync.Mutex
	sessions map[string]se.Session
}

// Provider return session provider
func Provider() *provider {
	return &provider{&sync.Mutex{}, map[string]se.Session{}}
}

// CookieName return cookie name
func (p *provider) CookieName() string {
	return "GOSESSID"
}

// GetId get session id
func (p *provider) GetId(r *http.Request) string {
	cookie, err := r.Cookie(p.CookieName())
	if err == nil && cookie != nil {
		return cookie.Value
	}
	return ""
}

// Exists session
func (p *provider) Exists(id string) bool {
	_, have := p.sessions[id]
	return have
}

// Get session
func (p *provider) Get(id string) se.Session {
	value, have := p.sessions[id]
	if have {
		return value.(se.Session)
	}
	return nil
}

// Del session
func (p *provider) Del(id string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.sessions, id)
}

// GetAll sessions
func (p *provider) GetAll() map[string]se.Session {
	return p.sessions
}

// Clear sessions
func (p *provider) Clear() {
	for k := range p.GetAll() {
		p.Del(k)
	}
}

func tmd5(text string) string {
	hashMd5 := md5.New()
	_, _ = io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func newSID() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	return strings.ToUpper(tmd5(tmd5(strconv.FormatInt(nano, 10)) + tmd5(strconv.FormatInt(rndNum, 10))))
}

// New return new session
func (p *provider) New(config *se.Config, listener *se.Listener) se.Session {
	sessionId := newSID()
	currentSession := newSession(sessionId, config.Valid)
	p.mu.Lock()
	p.sessions[sessionId] = currentSession
	p.mu.Unlock()
	go func(currentSession se.Session) {
		if listener != nil && listener.Created != nil {
			listener.Created(currentSession)
		}
	}(currentSession)
	time.Sleep(time.Nanosecond)
	return currentSession
}

// Refresh session
func (p *provider) Refresh(session se.Session, config *se.Config, listener *se.Listener) {
	session.Renew(config.Valid)
	go func() {
		if listener != nil && listener.Refreshed != nil {
			listener.Refreshed(session)
		}
	}()
}

// Clean session
func (p *provider) Clean(_ *se.Config, listener *se.Listener) {
	p.cleanSession(listener)
	time.AfterFunc(time.Second, func() { p.Clean(nil, listener) })
}

func (p *provider) cleanSession(listener *se.Listener) {
	for _, currentSession := range p.GetAll() {
		nu := time.Now().Unix()
		cu := currentSession.(*session).expiresTime.Unix()
		invalidated := false
		if cu <= nu {
			invalidated = true
		}
		if invalidated {
			currentSession.Invalidate()
			go func(currentSession se.Session) {
				if listener != nil && listener.Invalidated != nil {
					listener.Invalidated(currentSession)
				}
			}(currentSession)
		}
		if currentSession.Invalidated() {
			p.Del(currentSession.Id())
			go func(currentSession se.Session) {
				if listener != nil && listener.Destroyed != nil {
					listener.Destroyed(currentSession)
				}
			}(currentSession)
		}
	}
}
