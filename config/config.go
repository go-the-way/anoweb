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
	t "html/template"
	"net/http"
	"time"

	"gopkg.in/yaml.v3"
)

// Config struct
type Config struct {
	Server   *Server   `yaml:"server"`
	Banner   *Banner   `yaml:"banner"`
	Template *Template `yaml:"template"`
}

// Unmarshal yaml
func (c *Config) Unmarshal(bytes []byte) error {
	return yaml.Unmarshal(bytes, c)
}

// Server Config Server
type Server struct {
	Host              string        `yaml:"host"`
	Port              int           `yaml:"port"`
	TLS               *TLS          `yaml:"tls"`
	MaxHeaderSize     int           `yaml:"max_header_size"`
	ReadTimeout       time.Duration `yaml:"read_timeout"`
	ReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	WriteTimeout      time.Duration `yaml:"write_timeout"`
	IdleTimeout       time.Duration `yaml:"idle_timeout"`
}

// Banner Config Banner
type Banner struct {
	Enable bool   `yaml:"enable"`
	Type   string `yaml:"type"`
	File   string `yaml:"file"`
	Text   string `yaml:"text"`
}

// Template Config Template
type Template struct {
	Cache   bool   `yaml:"cache"`
	Root    string `yaml:"root"`
	Suffix  string `yaml:"suffix"`
	FuncMap t.FuncMap
}

// TLS Config TLS
type TLS struct {
	Enable   bool   `yaml:"enable"`
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
}

// Default Config
func Default() *Config {
	return &Config{
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
}
