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
	"io/ioutil"
	"strings"
)

var banners []string

func init() {
	banners = append(banners, `                                    _`)
	banners = append(banners, `                                   | |`)
	banners = append(banners, ` _____  ____    ___   _ _ _  _____ | |__`)
	banners = append(banners, `(____ ||  _ \  / _ \ | | | || ___ ||  _ \ `)
	banners = append(banners, `/ ___ || | | || |_| || | | || ____|| |_) )`)
	banners = append(banners, `\_____||_| |_| \___/  \___/ |_____)|____/ `)
}

func (a *App) printBanner() {
	if a.Config.Banner.Enable {
		switch a.Config.Banner.Type {
		case "default":
			_, _ = fmt.Println(strings.Join(banners, "\n"))
		case "text":
			_, _ = fmt.Println(a.Config.Banner.Text)
		case "file":
			if bytes, _ := ioutil.ReadFile(a.Config.Banner.File); bytes != nil {
				_, _ = fmt.Printf("%v\n", string(bytes))
			}
		}
	}
}
