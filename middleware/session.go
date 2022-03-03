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
	"time"

	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/headers"
	se "github.com/go-the-way/anoweb/session"
)

const sessionMWName = "Session"

type session struct {
	Listener *se.Listener
	Config   *se.Config
	Provider se.Provider
}

// Session return new session
func Session(provider se.Provider, config *se.Config, listener *se.Listener) *session {
	go provider.Clean(config, listener)
	return &session{
		Provider: provider,
		Config:   config,
		Listener: listener,
	}
}

// Name implements
func (s *session) Name() string {
	return sessionMWName
}

func (s *session) setSession(ctx *context.Context, session se.Session) {
	ctx.SetData("CURRENT_SESSION", session)
}

// GetSession get current session
func GetSession(ctx *context.Context) se.Session {
	currentSession := ctx.GetData("CURRENT_SESSION")
	if currentSession != nil {
		ses, ok := currentSession.(se.Session)
		if ok {
			return ses
		}
	}
	return nil
}

// Handler implements
func (s *session) Handler() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		sessionId := s.Provider.GetId(ctx.Request)
		have := false
		if sessionId != "" {
			have = s.Provider.Exists(sessionId)
		}
		if have {
			currentSession := s.Provider.Get(sessionId)
			s.Provider.Refresh(currentSession, s.Config, s.Listener)
			s.setSession(ctx, currentSession)
			ctx.Response.Header.Add(headers.SetCookie, (&http.Cookie{
				Name:    s.Provider.CookieName(),
				Value:   currentSession.Id(),
				Expires: time.Now().Add(s.Config.Valid),
				Path:    "/",
			}).String())
		} else {
			currentSession := s.Provider.New(s.Config, s.Listener)
			s.setSession(ctx, currentSession)
			ctx.Response.Header.Add(headers.SetCookie, (&http.Cookie{
				Name:    s.Provider.CookieName(),
				Value:   currentSession.Id(),
				Expires: time.Now().Add(s.Config.Valid),
				Path:    "/",
			}).String())
		}
		ctx.Chain()
	}
}
