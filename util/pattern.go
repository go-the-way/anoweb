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

package util

import (
	"regexp"
	"strings"
)

// TrimSpecialChars string
func TrimSpecialChars(str string) string {
	re := regexp.MustCompile(`[^/\w-._{}]`)
	return ReBuildPath(re.ReplaceAllString(str, ""))
}

// ReBuildPath re-build pattern
func ReBuildPath(pattern string) string {
	ps := strings.Split(pattern, "/")
	newPs := make([]string, 0)
	// append head
	newPs = append(newPs, "")
	for _, p := range ps {
		p = strings.TrimSpace(p)
		if p != "" {
			newPs = append(newPs, p)
		}
	}
	return strings.Join(newPs, "/")
}
