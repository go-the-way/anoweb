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

package router

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGroup(T *testing.T) {
	{
		g := NewGroup("")
		require.Equal(T, "", g.prefix)
	}
	{
		g := NewGroup("/group")
		require.Equal(T, "/group", g.prefix)
	}
}

func TestGroupPrefix(t *testing.T) {
	g := NewGroup("/awesome")
	require.Equal(t, "/awesome", g.Prefix())
}

func TestGroupAdd(t *testing.T) {
	g := NewGroup("")
	g.Add(&Router{})
	g.Add(&Router{})
	g.Add(&Router{})
	require.Equal(t, 3, len(g.routers))
}

func TestGroupRouters(t *testing.T) {
	g := NewGroup("")
	g.Add(&Router{})
	g.Add(&Router{})
	g.Add(&Router{})
	require.Equal(t, 3, len(g.Routers()))
}
