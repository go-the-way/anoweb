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

	"github.com/go-the-way/anoweb/mime"
)

// Text Response text
func (ctx *Context) Text(text string) {
	ctx.Binary([]byte(text), mime.TEXT)
}

// TextFile Response text file
func (ctx *Context) TextFile(textFile string) {
	ctx.File(textFile, mime.TEXT)
}

// JSON Response JSON
func (ctx *Context) JSON(data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	ctx.Binary(jsonData, mime.JSON)
}

// JSONText Response JSON text
func (ctx *Context) JSONText(json string) {
	ctx.Binary([]byte(json), mime.JSON)
}

// JSONFile Response JSON file
func (ctx *Context) JSONFile(jsonFile string) {
	ctx.File(jsonFile, mime.JSON)
}

// XML Response XML
func (ctx *Context) XML(data any) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		panic(err)
	}
	ctx.Binary(xmlData, mime.XML)
}

// XMLText Response XML text
func (ctx *Context) XMLText(xml string) {
	ctx.Binary([]byte(xml), mime.XML)
}

// XMLFile Response XML file
func (ctx *Context) XMLFile(xmlFile string) {
	ctx.File(xmlFile, mime.XML)
}

// HTML Response html
func (ctx *Context) HTML(html string) {
	ctx.Binary([]byte(html), mime.HTML)
}

// HTMLFile Response html file
func (ctx *Context) HTMLFile(htmlFile string) {
	ctx.File(htmlFile, mime.HTML)
}

// CSS Response css
func (ctx *Context) CSS(css string) {
	ctx.Binary([]byte(css), mime.CSS)
}

// CSSFile Response css file
func (ctx *Context) CSSFile(cssFile string) {
	ctx.File(cssFile, mime.CSS)
}

// JS Response js
func (ctx *Context) JS(js string) {
	ctx.Binary([]byte(js), mime.JS)
}

// JSFile Response js file
func (ctx *Context) JSFile(jsFile string) {
	ctx.File(jsFile, mime.JS)
}
