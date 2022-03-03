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
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBannerWithDefault(t *testing.T) {
	outFile := "std.out"
	file, err := os.Create(outFile)
	if err != nil {
		require.Error(t, err)
	}
	defer func() {
		_ = os.Remove(outFile)
	}()
	os.Stdout = file
	a := New()
	banner := "+>>>>>>>>>>>>>>>>>>>+ hello world +>>>>>>>>>>>>>>>>>>>"
	banners = []string{banner}
	a.printBanner()
	_ = file.Close()
	bytes, err := ioutil.ReadFile(outFile)
	if err != nil {
		require.Error(t, err)
	}
	require.Equal(t, banner, strings.TrimSpace(string(bytes)))
}

func TestBannerWithText(t *testing.T) {
	outFile := "std.out"
	file, err := os.Create(outFile)
	if err != nil {
		require.Error(t, err)
	}
	defer func() {
		_ = os.Remove(outFile)
	}()
	os.Stdout = file
	a := New()
	banner := "+>>>>>>>>>>>>>>>>>>>+ hello world +>>>>>>>>>>>>>>>>>>>"
	a.Config.Banner.Type = "text"
	a.Config.Banner.Text = banner
	a.printBanner()
	_ = file.Close()
	bytes, err := ioutil.ReadFile(outFile)
	if err != nil {
		require.Error(t, err)
	}
	require.Equal(t, banner, strings.TrimSpace(string(bytes)))
}

func TestBannerWithFile(t *testing.T) {
	var (
		err        error
		banner     = "+>>>>>>>>>>>>>>>>>>>+ hello world +>>>>>>>>>>>>>>>>>>>"
		outFile    = "std.out"
		bannerFile = "banner.txt"
		file, _    = os.Create(outFile)
	)
	err = ioutil.WriteFile(bannerFile, []byte(banner), 0700)
	if err != nil {
		require.Error(t, err)
	}
	defer func() {
		_ = os.Remove(outFile)
		_ = os.Remove(bannerFile)
	}()
	os.Stdout = file
	a := New()
	a.Config.Banner.Type = "file"
	a.Config.Banner.File = bannerFile
	a.printBanner()
	_ = file.Close()
	bytes, err := ioutil.ReadFile(outFile)
	if err != nil {
		require.Error(t, err)
	}
	require.Equal(t, banner, strings.TrimSpace(string(bytes)))
}
