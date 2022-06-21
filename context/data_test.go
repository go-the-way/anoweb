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
	"github.com/go-the-way/anoweb/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetData(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	ctx.SetData("apple", 100)
	require.NotNil(t, ctx.dataMap["apple"])
	require.Equal(t, ctx.dataMap["apple"], 100)
}

func TestSetDataMap(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	ctx.SetDataMap(map[string]interface{}{"apple": 100, "orange": 200}, false)
	require.NotNil(t, ctx.dataMap["apple"])
	require.Equal(t, ctx.dataMap["apple"], 100)
	require.NotNil(t, ctx.dataMap["orange"])
	require.Equal(t, ctx.dataMap["orange"], 200)
}

func TestSetDataMapWithFlush(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	ctx.SetData("apple", 10)
	ctx.SetDataMap(map[string]interface{}{"apple": 100, "orange": 200}, true)
	require.NotNil(t, ctx.dataMap["apple"])
	require.Equal(t, ctx.dataMap["apple"], 100)
	require.NotNil(t, ctx.dataMap["orange"])
	require.Equal(t, ctx.dataMap["orange"], 200)
}

func TestGetData(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	ctx.SetData("apple", 100)
	require.NotNil(t, ctx.GetData("apple"))
	require.Equal(t, ctx.GetData("apple"), 100)
}

func TestGetDataMap(t *testing.T) {
	ctx := New()
	ctx.Allocate(buildReq(""), &config.Template{})
	ctx.SetData("apple", 100)
	require.NotNil(t, ctx.GetDataMap())
	require.Equal(t, ctx.GetDataMap()["apple"], 100)
}
