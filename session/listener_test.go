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

package session

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListener(t *testing.T) {

	var (
		p           = int32(0)
		created     = func(session Session) { atomic.AddInt32(&p, 1) }
		destroyed   = func(session Session) { atomic.AddInt32(&p, 1) }
		invalidated = func(session Session) { atomic.AddInt32(&p, 1) }
		refreshed   = func(session Session) { atomic.AddInt32(&p, 1) }
	)

	ls := &Listener{
		Created:     created,
		Destroyed:   destroyed,
		Invalidated: invalidated,
		Refreshed:   refreshed,
	}

	ls.Created(nil)
	require.Equal(t, p, int32(1))

	ls.Destroyed(nil)
	require.Equal(t, p, int32(2))

	ls.Invalidated(nil)
	require.Equal(t, p, int32(3))

	ls.Refreshed(nil)
	require.Equal(t, p, int32(4))

}
