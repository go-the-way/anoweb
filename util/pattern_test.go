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

package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReBuildPath(t *testing.T) {

	type _case struct {
		pattern string
		expect  string
	}

	cases := []*_case{
		{"hello", "/hello"},
		{"world", "/world"},
		{"hello1a2b3c", "/hello1a2b3c"},
		{"hello,{}[]-=*/-12vb", "/hello,{}[]-=*/-12vb"},
		{"/hello/world", "/hello/world"},
		{"/hello_world/abc/xyz", "/hello_world/abc/xyz"},
	}

	for _, c := range cases {
		require.Equal(t, c.expect, ReBuildPath(c.pattern))
	}

}

func TestTrimSpecialChars(t *testing.T) {
	type _case struct {
		pattern string
		expect  string
	}

	cases := []*_case{
		{"hello,{}[]-=*/-12vb", "/hello{}-/-12vb"},
		{"/hello_world/abc/xyz", "/hello_world/abc/xyz"},
		{"/index//*-0.3484", "/index/-0.3484"},
		{"/index/we/{abc}/_=-hello/***/-world", "/index/we/{abc}/_-hello/-world"},
		{"/{id}/{name}/hello.world/index.%￥！@#￥%……&*（）——+~！-*/+go_go_go-wold/", "/{id}/{name}/hello.world/index.-/go_go_go-wold"},
	}

	for _, c := range cases {
		require.Equal(t, c.expect, TrimSpecialChars(c.pattern))
	}

}
