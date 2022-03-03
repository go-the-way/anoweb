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
	"fmt"
	"net/http"

	"github.com/go-the-way/anoweb/rest"
	"github.com/go-the-way/anoweb/util"
)

// Controller Route REST-ful Controller
func (a *App) Controller(c ...rest.Controller) *App {
	a.controllers = append(a.controllers, c...)
	return a
}

func (a *App) routeRestControllers() *App {
	for _, c := range a.controllers {
		prefix := util.TrimSpecialChars(c.Prefix())
		if c.Get() != nil {
			a.Route(http.MethodGet, fmt.Sprintf("%s/{RESTFUL_KEY}", prefix), c.Get())
		}
		if c.Gets() != nil {
			a.Route(http.MethodGet, prefix, c.Gets())
		}
		if c.Post() != nil {
			a.Route(http.MethodPost, prefix, c.Post())
		}
		if c.Put() != nil {
			a.Route(http.MethodPut, fmt.Sprintf("%s/{RESTFUL_KEY}", prefix), c.Put())
		}
		if c.Delete() != nil {
			a.Route(http.MethodDelete, fmt.Sprintf("%s/{RESTFUL_KEY}", prefix), c.Delete())
		}
	}
	return a
}
