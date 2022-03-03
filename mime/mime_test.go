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

package mime

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMime(t *testing.T) {
	for _, ms := range [][]string{
		{TEXT, "text/plain;charset=utf-8"},
		{HTML, "text/html;charset=utf-8"},
		{JS, "text/javascript;charset=utf-8"},
		{CSS, "text/css;charset=utf-8"},
		{JSON, "application/json;charset=utf-8"},
		{XML, "application/xml;charset=utf-8"},
		{BMP, "image/bmp"},
		{JPG, "image/jpg"},
		{PNG, "image/png"},
		{GIF, "image/gif"},
		{ICO, "image/ico"},
		{ZIP, "application/zip"},
		{TAR, "application/x-tar"},
		{GZIP, "application/x-gzip"},
		{TGZ, "application/x-tgz"},
		{RAR, "application/x-rar-compressed"},
		{ZIP7, "application/x-7z-compressed"},
		{XLS, "application/vnd.ms-excel"},
		{XLSX, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
		{DOC, "application/msword"},
		{DOCX, "application/vnd.openxmlformats-officedocument.wordprocessingml.document"},
		{PPT, "application/vnd.ms-powerpoint"},
		{PPTX, "application/vnd.openxmlformats-officedocument.presentationml.presentation"},
		{BINARY, "application/octet-stream"},
	} {
		require.Equal(t, ms[1], ms[0])
	}
}
