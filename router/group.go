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

// Group struct
type Group struct {
	prefix  string
	routers []*Router
}

// NewGroup return new group
func NewGroup(prefix string) *Group {
	return &Group{
		prefix:  prefix,
		routers: make([]*Router, 0),
	}
}

// Add routers
func (g *Group) Add(r ...*Router) *Group {
	g.routers = append(g.routers, r...)
	return g
}

// Prefix return group's prefix
func (g *Group) Prefix() string {
	return g.prefix
}

// Routers return group's routers
func (g *Group) Routers() []*Router {
	return g.routers
}
