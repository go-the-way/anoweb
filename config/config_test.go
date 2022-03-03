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

package config

import (
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"net/http"
	"testing"
	"time"
)

func TestUnmarshal(t *testing.T) {
	c := &Config{}
	bytes, _ := yaml.Marshal(Default())
	require.Nil(t, c.Unmarshal(bytes))
}

func TestDefault(t *testing.T) {
	c := &Config{
		Server: &Server{
			Host: "0.0.0.0",
			Port: 9494,
			TLS: &TLS{
				Enable:   false,
				CertFile: "",
				KeyFile:  "",
			},
			MaxHeaderSize:     http.DefaultMaxHeaderBytes,
			ReadTimeout:       time.Minute,
			ReadHeaderTimeout: time.Minute,
			WriteTimeout:      time.Minute,
			IdleTimeout:       time.Second,
		},
		Banner: &Banner{
			Enable: true,
			Type:   "default",
			File:   "banner.txt",
			Text:   "anoweb",
		},
		Template: &Template{
			Cache:  true,
			Root:   "templates",
			Suffix: ".html",
		},
	}

	require.Equal(t, c, Default())
}
