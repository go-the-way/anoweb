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

import v "github.com/go-the-way/validator"

func (ctx *Context) Validate(structPtr interface{}, call func()) {
	ctx.ValidateWithParams(structPtr, "message", "Parameters is invalid", "code", 500, call)
}

func (ctx *Context) ValidateWithParams(structPtr interface{}, messageName, message, codeName string, code int, call func()) {
	result := v.New(structPtr).Validate()
	if result.Passed {
		call()
	} else {
		msg := result.Messages()
		if msg == "" {
			msg = message
		}
		ctx.JSON(map[string]interface{}{messageName: msg, codeName: code})
	}
}

func (ctx *Context) BindAndValidate(structPtr interface{}, call func()) {
	ctx.Bind(structPtr)
	ctx.Validate(structPtr, call)
}

func (ctx *Context) BindAndValidateWithParams(structPtr interface{}, messageName, message, codeName string, code int, call func()) {
	ctx.Bind(structPtr)
	ctx.ValidateWithParams(structPtr, messageName, message, codeName, code, call)
}
