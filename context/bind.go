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
	"encoding/json"
	"encoding/xml"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func (ctx *Context) BindJSON(ptr any) error {
	readAll, _ := ioutil.ReadAll(ctx.Request.Body)
	return json.Unmarshal(readAll, ptr)
}

func (ctx *Context) BindXML(ptr any) error {
	readAll, _ := ioutil.ReadAll(ctx.Request.Body)
	return xml.Unmarshal(readAll, ptr)
}

func (ctx *Context) BindYAML(ptr any) error {
	readAll, _ := ioutil.ReadAll(ctx.Request.Body)
	return yaml.Unmarshal(readAll, ptr)
}
