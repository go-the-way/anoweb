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

package anoweb

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-the-way/anoweb/context"
	"github.com/stretchr/testify/require"
)

type _responseWriter struct {
	statusCode int
	buf        strings.Builder
}

func (w *_responseWriter) Header() http.Header {
	return http.Header{}
}

func (w *_responseWriter) Write(bytes []byte) (int, error) {
	return w.buf.Write(bytes)
}

func (w *_responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func TestNewDispatcher(t *testing.T) {
	a := New()
	require.Equal(t, a.newDispatcher().App, dispatcher{a}.App)
}

func TestDispatcherServeHTTP(t *testing.T) {
	message := `hello world`
	a := New().Get("/", func(ctx *context.Context) {
		ctx.Text(message)
	}).parseRouters()
	responseWriter := &_responseWriter{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost", nil)
	if err != nil {
		require.Error(t, err)
	}
	a.newDispatcher().ServeHTTP(responseWriter, req)
	require.Equal(t, responseWriter.buf.String(), message)
	require.Equal(t, responseWriter.statusCode, http.StatusOK)
}

func TestDispatcherDispatch(t *testing.T) {
	message := `hello world`
	a := New().Get("/", func(ctx *context.Context) {
		ctx.Text(message)
	}).parseRouters()
	responseWriter := &_responseWriter{}
	req, err := http.NewRequest(http.MethodGet, "http://localhost", nil)
	if err != nil {
		require.Error(t, err)
	}
	_dispatcher := a.newDispatcher()
	_dispatcher.dispatch(req, responseWriter)
	require.Equal(t, responseWriter.buf.String(), message)
	require.Equal(t, responseWriter.statusCode, http.StatusOK)
}

func TestDispatcherWriteDone(t *testing.T) {
	message := `hello world`
	a := New()
	responseWriter := &_responseWriter{}
	_dispatcher := a.newDispatcher()
	_dispatcher.writeDone(context.Builder().Data([]byte(message)).Cookies([]*http.Cookie{{Name: "haha"}}).Build(), responseWriter)
	require.Equal(t, responseWriter.buf.String(), message)
	require.Equal(t, responseWriter.statusCode, http.StatusOK)
}
