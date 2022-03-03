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
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sync/atomic"
	"testing"
	"time"
)

var (
	_port = int32(10000)
)

func nextPort() int {
	return int(atomic.AddInt32(&_port, 1))
}

type testHTTPCase struct {
	method  string
	reqPath string
	body    string
	forms   url.Values
	expect  string
}

func testHTTP(t *testing.T, a *App, cases ...*testHTTPCase) {
	errCh := make(chan error, 1)
	okCh := make(chan struct{}, 1)
	port := nextPort()
	go func(port int) {
		a.Config.Server.Port = port
		// override env server
		_ = os.Setenv(envServerHost, "localhost")
		// override env port
		_ = os.Setenv(envServerPort, fmt.Sprintf("%d", a.Config.Server.Port))
		// override env tls enable
		_ = os.Setenv(envServerTLSEnable, "false")
		a.Run()
	}(port)
	time.AfterFunc(time.Millisecond*500, func() {
		for _, thc := range cases {
			var buf bytes.Buffer
			if thc.body != "" {
				buf.WriteString(thc.body)
			}
			req, err := http.NewRequest(thc.method, fmt.Sprintf("http://localhost:%d%s", port, thc.reqPath), &buf)
			if err != nil {
				errCh <- err
				return
			}
			req.Form = thc.forms
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				errCh <- err
				return
			}
			readAll, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				errCh <- err
				return
			}
			if reflect.DeepEqual(string(readAll), thc.expect) {
				okCh <- struct{}{}
			} else {
				errCh <- errors.New("not equal")
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
