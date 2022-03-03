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

package context

import (
	"net/http"
	"testing"

	"github.com/go-the-way/anoweb/config"

	"github.com/stretchr/testify/require"
)

func buildParamReq() *http.Request {
	return buildParamReqWithPS("")
}

func buildParamReqWithPS(ps string) *http.Request {
	req, _ := http.NewRequest("", "/?apple=100&orange=200&pear=300&apple=110"+ps, nil)
	return req
}

func TestParamMap(t *testing.T) {
	ctx := New(buildParamReq(), &config.Template{})
	require.Equal(t, map[string][]string{"apple": {"100", "110"}, "orange": {"200"}, "pear": {"300"}}, ctx.ParamMap())
}

func TestParam(t *testing.T) {
	ctx := New(buildParamReq(), &config.Template{})
	require.Equal(t, "", ctx.Param("banana"))
	require.Equal(t, "100", ctx.Param("apple"))
}

func TestParams(t *testing.T) {
	ctx := New(buildParamReq(), &config.Template{})
	require.Equal(t, []string{}, ctx.Params("banana", []string{}))
	require.Equal(t, []string{"100", "110"}, ctx.Params("apple", nil))
}

func TestHasParam(t *testing.T) {
	ctx := New(buildParamReq(), &config.Template{})
	require.Equal(t, false, ctx.HasParam("banana"))
	require.Equal(t, true, ctx.HasParam("apple"))
}

func TestIntParam(t *testing.T) {
	ctx := New(buildParamReqWithPS("&banana=hello"), &config.Template{})
	require.Equal(t, 10, int(ctx.IntParam("banana", 10)))
	require.Equal(t, 100, int(ctx.IntParam("apple", 0)))
	require.Equal(t, 10, int(ctx.IntParam("banana", 10)))
	require.Equal(t, 20, int(ctx.IntParam("banana2", 20)))
}

func TestFloatParam(t *testing.T) {
	ctx := New(buildParamReqWithPS("&banana=hello"), &config.Template{})
	require.Equal(t, float64(10), ctx.FloatParam("banana", 10))
	require.Equal(t, float64(100), ctx.FloatParam("apple", 0))
	require.Equal(t, float64(10), ctx.FloatParam("banana", 10))
	require.Equal(t, float64(20), ctx.FloatParam("banana2", 20))
}

func TestSingleParamMap(t *testing.T) {
	ctx := New(buildParamReq(), &config.Template{})
	require.Equal(t, map[string]string{"apple": "100", "orange": "200", "pear": "300"}, ctx.SingleParamMap())
	ctx.SetParamMap(map[string][]string{"apple": {}}, true)
	require.Equal(t, map[string]string{"apple": ""}, ctx.SingleParamMap())
}

func TestJoinedParamMap(t *testing.T) {
	ctx := New(buildParamReq(), &config.Template{})
	require.Equal(t, map[string]string{"apple": "100|110", "orange": "200", "pear": "300"}, ctx.JoinedParamMap("|"))
	ctx.SetParamMap(map[string][]string{"apple": {}}, true)
	require.Equal(t, map[string]string{"apple": ""}, ctx.JoinedParamMap("|"))
}

func TestSetParamMap(t *testing.T) {
	ctx := New(buildParamReq(), &config.Template{})
	ctx.SetParamMap(map[string][]string{"apple": {"apple"}}, false)
	require.Equal(t, "apple", ctx.ParamDefault("apple", "apple"))
}

func TestKey(t *testing.T) {
	ctx := New(buildParamReq(), &config.Template{})
	ctx.SetParamMap(map[string][]string{"RESTFUL_KEY": {"apple"}}, false)
	require.Equal(t, "apple", ctx.Key())
}

func TestIntKey(t *testing.T) {
	ctx := New(buildParamReq(), &config.Template{})
	ctx.SetParamMap(map[string][]string{"RESTFUL_KEY": {"100"}}, false)
	require.Equal(t, int64(100), ctx.IntKey())
}

func TestIntKeyPanic(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	ctx := New(buildParamReq(), &config.Template{})
	ctx.SetParamMap(map[string][]string{"RESTFUL_KEY": {"hello"}}, false)
	ctx.IntKey()
}
