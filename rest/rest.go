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

package rest

import "github.com/go-the-way/anoweb/context"

// Controller interface
type Controller interface {
	// Prefix route pattern prefix
	Prefix() string
	// Gets route => GET: /${Prefix}
	Gets() func(ctx *context.Context)
	// Get route => GET: /${Prefix}/${RESTFUL_KEY}
	Get() func(ctx *context.Context)
	// Post route => POST: /${Prefix}
	Post() func(ctx *context.Context)
	// Put route => PUT: /${Prefix}/${RESTFUL_KEY}
	Put() func(ctx *context.Context)
	// Delete route => DELETE: /${Prefix}/${RESTFUL_KEY}
	Delete() func(ctx *context.Context)
}
