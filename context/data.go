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

// SetData Set data into context
func (ctx *Context) SetData(name string, value interface{}) *Context {
	ctx.dataMap[name] = value
	return ctx
}

// SetDataMap Set data map into context
func (ctx *Context) SetDataMap(dataMap map[string]interface{}, flush bool) *Context {
	if flush {
		for k := range ctx.dataMap {
			delete(ctx.dataMap, k)
		}
	}
	if dataMap != nil {
		for k, v := range dataMap {
			ctx.SetData(k, v)
		}
	}
	return ctx
}

// GetData get data from context
func (ctx *Context) GetData(name string) interface{} {
	return ctx.dataMap[name]
}

// GetDataMap get data map from context
func (ctx *Context) GetDataMap() map[string]interface{} {
	return ctx.dataMap
}
