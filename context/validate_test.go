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
	"github.com/go-the-way/anoweb/mime"

	"github.com/stretchr/testify/require"
)

func TestContextValidate(t *testing.T) {
	type _model struct {
		ID   int    `json:"id" validate:"msg(id验证未通过) min(10)"`
		Name string `json:"name" validate:"msg(name验证未通过) maxlength(10)"`
	}
	// test for validated
	{
		m := _model{100, "hello"}
		ctx := New()
		ctx.Allocate(buildReq(""), &config.Template{})
		called := false
		ctx.Validate(&m, func() { called = true })
		require.Equal(t, true, called)
	}
	// test for not validated
	{
		m := _model{1, "hello"}
		ctx := New()
		ctx.Allocate(buildReq(""), &config.Template{})
		called := false
		ctx.Validate(&m, func() { called = true })
		require.Equal(t, false, called)
		rendered := ctx.Response
		require.Equal(t, http.StatusOK, rendered.Status)
		require.Equal(t, mime.JSON, rendered.ContentType)
		require.Equal(t, `{"code":500,"message":"id验证未通过"}`, string(rendered.Data))
	}
	// test for not validated with 2 fields
	{
		m := _model{1, "hello-world-world"}
		ctx := New()
		ctx.Allocate(buildReq(""), &config.Template{})
		called := false
		ctx.Validate(&m, func() { called = true })
		require.Equal(t, false, called)
		rendered := ctx.Response
		require.Equal(t, http.StatusOK, rendered.Status)
		require.Equal(t, mime.JSON, rendered.ContentType)
		require.Equal(t, `{"code":500,"message":"id验证未通过,name验证未通过"}`, string(rendered.Data))
	}
}

func TestValidateWithParams(t *testing.T) {
	// test for not validated
	{
		type _model struct {
			ID   int    `json:"id" validate:"min(10)"`
			Name string `json:"name" validate:"maxlength(10)"`
		}
		m := _model{1, "hello"}
		ctx := New()
		ctx.Allocate(buildReq(""), &config.Template{})
		called := false
		ctx.ValidateWithParams(&m, "_message_", "fail", "_code_", 15000, func() { called = true })
		require.Equal(t, false, called)
		rendered := ctx.Response
		require.Equal(t, http.StatusOK, rendered.Status)
		require.Equal(t, mime.JSON, rendered.ContentType)
		require.Equal(t, `{"_code_":15000,"_message_":"fail"}`, string(rendered.Data))
	}
	// test for not validated with 2 fields
	{
		type _model struct {
			ID   int    `json:"id" validate:"msg(id验证未通过) min(10)"`
			Name string `json:"name" validate:"msg(name验证未通过) maxlength(10)"`
		}
		m := _model{1, "hello-world-world"}
		ctx := New()
		ctx.Allocate(buildReq(""), &config.Template{})
		called := false
		ctx.ValidateWithParams(&m, "_message_", "fail", "_code_", 15000, func() { called = true })
		require.Equal(t, false, called)
		rendered := ctx.Response
		require.Equal(t, http.StatusOK, rendered.Status)
		require.Equal(t, mime.JSON, rendered.ContentType)
		require.Equal(t, `{"_code_":15000,"_message_":"id验证未通过,name验证未通过"}`, string(rendered.Data))
	}
}

func TestBindAndValidate(t *testing.T) {
	type _model struct {
		ID   int    `json:"id" validate:"msg(id验证未通过) min(10)"`
		Name string `json:"name" validate:"msg(name验证未通过) maxlength(10)"`
	}
	m := _model{}
	// test for validated
	{
		ctx := New()
		ctx.Allocate(buildReq(`{"id":100,"name":"hello"}`), &config.Template{})
		called := false
		ctx.BindAndValidate(&m, func() { called = true })
		require.Equal(t, true, called)
	}
	// test for not validated
	{
		ctx := New()
		ctx.Allocate(buildReq(`{"id":1,"name":"hello"}`), &config.Template{})
		called := false
		ctx.BindAndValidate(&m, func() { called = true })
		require.Equal(t, false, called)
		rendered := ctx.Response
		require.Equal(t, http.StatusOK, rendered.Status)
		require.Equal(t, mime.JSON, rendered.ContentType)
		require.Equal(t, `{"code":500,"message":"id验证未通过"}`, string(rendered.Data))
	}
	// test for not validated with 2 fields
	{
		ctx := New()
		ctx.Allocate(buildReq(`{"id":1,"name":"hello-world-world"}`), &config.Template{})
		called := false
		ctx.BindAndValidate(&m, func() { called = true })
		require.Equal(t, false, called)
		rendered := ctx.Response
		require.Equal(t, http.StatusOK, rendered.Status)
		require.Equal(t, mime.JSON, rendered.ContentType)
		require.Equal(t, `{"code":500,"message":"id验证未通过,name验证未通过"}`, string(rendered.Data))
	}
}

func TestBindAndValidateWithParams(t *testing.T) {
	type _model struct {
		ID   int    `json:"id" validate:"min(10)"`
		Name string `json:"name" validate:"maxlength(10)"`
	}
	m := _model{}
	// test for validated
	{
		ctx := New()
		ctx.Allocate(buildReq(`{"id":100,"name":"hello"}`), &config.Template{})
		called := false
		ctx.BindAndValidateWithParams(&m, "", "", "", 0, func() { called = true })
		require.Equal(t, true, called)
	}
	// test for not validated
	{
		ctx := New()
		ctx.Allocate(buildReq(`{"id":1,"name":"hello"}`), &config.Template{})
		called := false
		ctx.BindAndValidateWithParams(&m, "_message_", "fail", "_code_", 15000, func() { called = true })
		require.Equal(t, false, called)
		rendered := ctx.Response
		require.Equal(t, http.StatusOK, rendered.Status)
		require.Equal(t, mime.JSON, rendered.ContentType)
		require.Equal(t, `{"_code_":15000,"_message_":"fail"}`, string(rendered.Data))
	}
	// test for not validated with 2 fields
	{
		ctx := New()
		ctx.Allocate(buildReq(`{"id":1,"name":"hello-world-world"}`), &config.Template{})
		called := false
		ctx.BindAndValidateWithParams(&m, "_message_", "fail", "_code_", 15000, func() { called = true })
		require.Equal(t, false, called)
		rendered := ctx.Response
		require.Equal(t, http.StatusOK, rendered.Status)
		require.Equal(t, mime.JSON, rendered.ContentType)
		require.Equal(t, `{"_code_":15000,"_message_":"fail"}`, string(rendered.Data))
	}
}
