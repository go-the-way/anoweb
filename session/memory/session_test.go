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
	"errors"
	"fmt"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-the-way/anoweb"
	"github.com/go-the-way/anoweb/context"
	"github.com/go-the-way/anoweb/middleware"
	s "github.com/go-the-way/anoweb/session"
)

var _port = int32(9000)

func nextPort() int {
	return int(atomic.AddInt32(&_port, 1))
}

func currSessionId(resp *http.Response) string {
	sessionId := ""
	cookies := resp.Cookies()
	for _, c := range cookies {
		if c.Name == "GOSESSID" {
			sessionId = c.Value
			break
		}
	}
	return sessionId
}

func TestSession(t *testing.T) {
	sessionCh := make(chan string, 1)
	port := nextPort()
	go func() {
		app := anoweb.New()
		app.Config.Server.Port = port
		app.Get("/", func(ctx *context.Context) {
			sessionCh <- middleware.GetSession(ctx).Id()
		}).UseSession(Provider(), &s.Config{Valid: time.Minute}, nil).Run()
	}()
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	time.AfterFunc(time.Millisecond*500, func() {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			errCh <- err
		}

		sessionId := ""
		cookies := resp.Cookies()
		for _, c := range cookies {
			if c.Name == "GOSESSID" {
				sessionId = c.Value
				break
			}
		}

		select {
		case <-time.After(time.Millisecond * 200):
			errCh <- errors.New("get session id timeout")
		case currSession := <-sessionCh:
			if sessionId == currSession {
				okCh <- struct{}{}
			} else {
				errCh <- errors.New("session id not equal")
			}
		}

	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}

}

func TestSessionRenew(t *testing.T) {
	port := nextPort()
	invalidatedCh := make(chan bool, 1)
	go func() {
		app := anoweb.New()
		app.Config.Server.Port = port
		app.Get("/renew", func(ctx *context.Context) {
			currSession := middleware.GetSession(ctx)
			invalidatedCh <- currSession.Invalidated()
		}).UseSession(Provider(), &s.Config{Valid: time.Minute}, nil).Run()
	}()
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	time.AfterFunc(time.Millisecond*500, func() {
		resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			errCh <- err
		}

		sessionId := currSessionId(resp)
		t.Logf("get session: %s", sessionId)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:%d/renew", port), nil)
		if err != nil {
			errCh <- err
		}
		req.AddCookie(&http.Cookie{Name: "GOSESSID", Value: sessionId})
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			errCh <- err
		}

		select {
		case <-time.After(time.Millisecond * 200):
			errCh <- errors.New("get session invalidated timeout")
		case invalidated := <-invalidatedCh:
			if !invalidated {
				okCh <- struct{}{}
			} else {
				errCh <- errors.New("session must be not invalidated")
			}
		}
	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}
}

func TestSessionInvalidated(t *testing.T) {
	port := nextPort()
	invalidatedCh := make(chan bool, 1)
	go func() {
		app := anoweb.New()
		app.Config.Server.Port = port
		app.Get("/", func(ctx *context.Context) {
			currSession := middleware.GetSession(ctx)
			invalidatedCh <- currSession.Invalidated()
		}).UseSession(Provider(), &s.Config{Valid: time.Minute}, nil).Run()
	}()
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	time.AfterFunc(time.Millisecond*500, func() {
		_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			errCh <- err
		}
		select {
		case <-invalidatedCh:
			okCh <- struct{}{}
		}
	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}

}

func TestSessionInvalidate(t *testing.T) {
	port := nextPort()
	invalidatedCh := make(chan bool, 1)
	go func() {
		app := anoweb.New()
		app.Config.Server.Port = port
		app.Get("/", func(ctx *context.Context) {
			currSession := middleware.GetSession(ctx)
			currSession.Invalidate()
			invalidatedCh <- currSession.Invalidated()
		}).UseSession(Provider(), &s.Config{Valid: time.Minute}, nil).Run()
	}()
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	time.AfterFunc(time.Millisecond*500, func() {
		_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			errCh <- err
		}

		select {
		case invalidated := <-invalidatedCh:
			if invalidated {
				okCh <- struct{}{}
			} else {
				errCh <- errors.New("not expected")
			}
		}
	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}
}

func TestSessionGet(t *testing.T) {
	port := nextPort()
	dataCh := make(chan string, 1)
	go func() {
		app := anoweb.New()
		app.Config.Server.Port = port
		app.Get("/", func(ctx *context.Context) {
			currSession := middleware.GetSession(ctx)
			currSession.Set("apple", "100")
			dataCh <- currSession.Get("apple").(string)
		}).UseSession(Provider(), &s.Config{Valid: time.Minute}, nil).Run()
	}()
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	time.AfterFunc(time.Millisecond*500, func() {
		_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			errCh <- err
		}

		select {
		case apple := <-dataCh:
			if apple == "100" {
				okCh <- struct{}{}
			} else {
				errCh <- errors.New("not expected")
			}
		}
	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}
}

func TestSessionGetAll(t *testing.T) {
	port := nextPort()
	dataCh := make(chan map[string]any, 1)
	go func() {
		app := anoweb.New()
		app.Config.Server.Port = port
		app.Get("/", func(ctx *context.Context) {
			currSession := middleware.GetSession(ctx)
			dataCh <- currSession.GetAll()
		}).UseSession(Provider(), &s.Config{Valid: time.Minute}, &s.Listener{Created: func(session s.Session) {
			session.SetAll(map[string]any{
				"apple":  "100",
				"banana": "200",
			}, false)
		}}).Run()
	}()
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	time.AfterFunc(time.Millisecond*500, func() {
		_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			errCh <- err
		}

		select {
		case data := <-dataCh:
			if data["apple"].(string) == "100" && data["banana"].(string) == "200" {
				okCh <- struct{}{}
			} else {
				errCh <- errors.New("not expected")
			}
		}
	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}
}

func TestSessionSet(t *testing.T) {
	port := nextPort()
	dataCh := make(chan string, 1)
	go func() {
		app := anoweb.New()
		app.Config.Server.Port = port
		app.Get("/", func(ctx *context.Context) {
			currSession := middleware.GetSession(ctx)
			currSession.Set("apple", "100")
			dataCh <- currSession.Get("apple").(string)
		}).UseSession(Provider(), &s.Config{Valid: time.Minute}, nil).Run()
	}()
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	time.AfterFunc(time.Millisecond*500, func() {
		_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			errCh <- err
		}

		select {
		case apple := <-dataCh:
			if apple == "100" {
				okCh <- struct{}{}
			} else {
				errCh <- errors.New("not expected")
			}
		}
	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}
}

func TestSessionSetAll(t *testing.T) {
	port := nextPort()
	dataCh := make(chan map[string]any, 1)
	go func() {
		app := anoweb.New()
		app.Config.Server.Port = port
		app.Get("/", func(ctx *context.Context) {
			currSession := middleware.GetSession(ctx)
			currSession.SetAll(nil, false)
			currSession.Set("apple", "100")
			currSession.SetAll(map[string]any{
				"apple":  "100",
				"banana": "200",
			}, true)
			dataCh <- currSession.GetAll()
		}).UseSession(Provider(), &s.Config{Valid: time.Minute}, nil).Run()
	}()
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	time.AfterFunc(time.Millisecond*500, func() {
		_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			errCh <- err
		}

		select {
		case data := <-dataCh:
			if data["apple"].(string) == "100" && data["banana"].(string) == "200" {
				okCh <- struct{}{}
			} else {
				errCh <- errors.New("not expected")
			}
		}
	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}
}

func TestSessionDel(t *testing.T) {
	port := nextPort()
	dataCh := make(chan any, 1)
	go func() {
		app := anoweb.New()
		app.Config.Server.Port = port
		app.Get("/", func(ctx *context.Context) {
			currSession := middleware.GetSession(ctx)
			currSession.Del("apple")
			dataCh <- currSession.Get("apple")
		}).UseSession(Provider(), &s.Config{Valid: time.Minute}, &s.Listener{Created: func(session s.Session) {
			session.Set("apple", "100")
		}}).Run()
	}()
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	time.AfterFunc(time.Millisecond*500, func() {
		_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			errCh <- err
		}

		select {
		case data := <-dataCh:
			if data == nil {
				okCh <- struct{}{}
			} else {
				errCh <- errors.New("not expected")
			}
		}
	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}
}

func TestSessionClear(t *testing.T) {
	port := nextPort()
	dataCh := make(chan any, 1)
	go func() {
		app := anoweb.New()
		app.Config.Server.Port = port
		app.Get("/", func(ctx *context.Context) {
			currSession := middleware.GetSession(ctx)
			currSession.Clear()
			dataCh <- currSession.Get("apple")
		}).UseSession(Provider(), &s.Config{Valid: time.Minute}, &s.Listener{Created: func(session s.Session) {
			session.Set("apple", "100")
		}}).Run()
	}()
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	time.AfterFunc(time.Millisecond*500, func() {
		_, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			errCh <- err
		}

		select {
		case data := <-dataCh:
			if data == nil {
				okCh <- struct{}{}
			} else {
				errCh <- errors.New("not expected")
			}
		}
	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}
}
