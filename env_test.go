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
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testEnvCase struct {
	env        string
	val        any
	beforeCall func()
	expectCall func() any
}

func TestEnv(t *testing.T) {
	a := New()
	cases := make([]*testEnvCase, 0)

	{
		// test for setServerMaxHeaderSize
		cases = append(cases, &testEnvCase{envServerMaxHeaderSize, 100, func() { a.setServerMaxHeaderSize() }, func() any { return a.Config.Server.MaxHeaderSize }})
		// test for setServerReadTimeout
		cases = append(cases, &testEnvCase{envServerReadTimeout, "10s", func() { a.setServerReadTimeout() }, func() any { return a.Config.Server.ReadTimeout }})
		// test for setServerReadHeaderTimeout
		cases = append(cases, &testEnvCase{envServerReadHeaderTimeout, "10s", func() { a.setServerReadHeaderTimeout() }, func() any { return a.Config.Server.ReadHeaderTimeout }})
		// test for setServerWriteTimeout
		cases = append(cases, &testEnvCase{envServerWriteTimeout, "10s", func() { a.setServerWriteTimeout() }, func() any { return a.Config.Server.WriteTimeout }})
		// test for setServerIdleTimeout
		cases = append(cases, &testEnvCase{envServerIdleTimeout, "10s", func() { a.setServerIdleTimeout() }, func() any { return a.Config.Server.IdleTimeout }})
		// test for setConfigFile
		cases = append(cases, &testEnvCase{envConfigFile, "app.yml", func() { a.setConfigFile() }, func() any { return a.ConfigFile }})
		// test for setServerHost
		cases = append(cases, &testEnvCase{envServerHost, "0.0.0.0", func() { a.setServerHost() }, func() any { return a.Config.Server.Host }})
		// test for setServerPort
		cases = append(cases, &testEnvCase{envServerPort, 1080, func() { a.setServerPort() }, func() any { return a.Config.Server.Port }})
		// test for setBannerEnable
		cases = append(cases, &testEnvCase{envBannerEnable, true, func() { a.setBannerEnable() }, func() any { return a.Config.Banner.Enable }})
		// test for setBannerType
		cases = append(cases, &testEnvCase{envBannerType, "default", func() { a.setBannerType() }, func() any { return a.Config.Banner.Type }})
		// test for setBannerText
		cases = append(cases, &testEnvCase{envBannerText, "hello world -- GO GO GO", func() { a.setBannerText() }, func() any { return a.Config.Banner.Text }})
		// test for setBannerFile
		cases = append(cases, &testEnvCase{envBannerFile, "banner.txt", func() { a.setBannerFile() }, func() any { return a.Config.Banner.File }})
		// test for setServerTLSEnable
		cases = append(cases, &testEnvCase{envServerTLSEnable, true, func() { a.setServerTLSEnable() }, func() any { return a.Config.Server.TLS.Enable }})
		// test for envServerTLSCertFile
		cases = append(cases, &testEnvCase{envServerTLSCertFile, "cert.pem", func() { a.setServerTLSCertFile() }, func() any { return a.Config.Server.TLS.CertFile }})
		// test for setServerTLSKeyFile
		cases = append(cases, &testEnvCase{envServerTLSKeyFile, "key.pem", func() { a.setServerTLSKeyFile() }, func() any { return a.Config.Server.TLS.KeyFile }})
		// test for setTemplateCache
		cases = append(cases, &testEnvCase{envTemplateCache, true, func() { a.setTemplateCache() }, func() any { return a.Config.Template.Cache }})
		// test for setTemplateRoot
		cases = append(cases, &testEnvCase{envTemplateRoot, "/to/path", func() { a.setTemplateRoot() }, func() any { return a.Config.Template.Root }})
		// test for setTemplateSuffix
		cases = append(cases, &testEnvCase{envTemplateSuffix, ".tpl", func() { a.setTemplateSuffix() }, func() any { return a.Config.Template.Suffix }})

	}

	for _, c := range cases {
		// set env
		_ = os.Setenv(c.env, fmt.Sprintf("%v", c.val))
		// before call
		c.beforeCall()
		expectVal := c.expectCall()
		// start require
		switch expectVal.(type) {
		default:
			require.Equal(t, c.val, expectVal)
		case time.Duration:
			duration, _ := time.ParseDuration(fmt.Sprintf("%v", c.val))
			require.Equal(t, duration, expectVal)
		}

	}
}
