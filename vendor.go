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
	"strings"
)

const version = "v1.0.6"

func (a *App) printVendor() {
	_, _ = fmt.Println(strings.ReplaceAll(`
::anoweb:: 

The lightweight and powerful web framework using the new way for Go. Another go the way.

{{ Version @VER }}
{{ Powered by go-the-way }}
{{ https://github.com/go-the-way/anoweb }}
`, "@VER", version))
}
