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
	"net/http"
	"testing"
	"time"

	se "github.com/go-the-way/anoweb/session"

	"github.com/stretchr/testify/require"
)

func TestProvider(t *testing.T) {
	p := Provider()
	require.NotNil(t, p.GetAll())
}

func TestProviderCookieName(t *testing.T) {
	p := Provider()
	require.Equal(t, "GOSESSID", p.CookieName())
}

func TestProviderGetId(t *testing.T) {
	p := Provider()
	req, _ := http.NewRequest("", "", nil)
	req.AddCookie(&http.Cookie{Name: p.CookieName(), Value: "hello---cookie---"})
	require.Equal(t, "hello---cookie---", p.GetId(req))
}

func TestProviderExists(t *testing.T) {
	p := Provider()
	currSession := p.New(&se.Config{Valid: time.Minute}, nil)
	require.NotNil(t, true, p.Exists(currSession.Id()))
}

func TestProviderGet(t *testing.T) {
	{
		p := Provider()
		time.Sleep(time.Millisecond * 100)
		require.Nil(t, p.Get("none"))
	}
	{
		p := Provider()
		currSession := p.New(&se.Config{Valid: time.Minute}, nil)
		time.Sleep(time.Millisecond * 100)
		require.NotNil(t, p.Get(currSession.Id()))
	}
}

func TestProviderDel(t *testing.T) {
	p := Provider()
	currSession := p.New(&se.Config{Valid: time.Minute}, nil)
	p.Del(currSession.Id())
	require.NotNil(t, false, p.Exists(currSession.Id()))
}

func TestProviderGetAll(t *testing.T) {
	p := Provider()
	_ = p.New(&se.Config{Valid: time.Minute}, nil)
	_ = p.New(&se.Config{Valid: time.Minute}, nil)
	p.Clear()
	_ = p.New(&se.Config{Valid: time.Minute}, nil)
	require.Equal(t, 1, len(p.GetAll()))
}

func TestProviderClear(t *testing.T) {
	p := Provider()
	p.Clear()
	require.Equal(t, 0, len(p.GetAll()))
}

func TestProviderRefresh(t *testing.T) {
	p := Provider()
	c := se.Config{Valid: time.Minute}
	currSession := p.New(&c, nil)
	p.Refresh(currSession, &c, &se.Listener{Refreshed: func(s se.Session) {}})
	require.Equal(t, false, currSession.Invalidated())
}

func TestProviderCleanSession(t *testing.T) {
	p := Provider()
	s1 := p.New(&se.Config{}, nil)
	s2 := p.New(&se.Config{}, nil)
	s3 := p.New(&se.Config{}, nil)
	cc := 0
	ccP := &cc
	time.Sleep(time.Millisecond * 100)
	p.cleanSession(&se.Listener{
		Destroyed:   func(s se.Session) { *ccP++ },
		Invalidated: func(s se.Session) { *ccP++ },
	})
	time.Sleep(time.Second * 2)
	require.Nil(t, p.Get(s1.Id()))
	require.Nil(t, p.Get(s2.Id()))
	require.Nil(t, p.Get(s3.Id()))
	require.Equal(t, 6, cc)
}
