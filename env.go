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
	"os"
	"strconv"
	"time"
)

const (
	envServerMaxHeaderSize     = "SERVER_MAX_HEADER_SIZE"
	envServerReadTimeout       = "SERVER_READ_TIMEOUT"
	envServerReadHeaderTimeout = "SERVER_READ_HEADER_TIMEOUT"
	envServerWriteTimeout      = "SERVER_WRITE_TIMEOUT"
	envServerIdleTimeout       = "SERVER_IDLE_TIMEOUT"
	envConfigFile              = "CONFIG_FILE"
	envServerHost              = "SERVER_HOST"
	envServerPort              = "SERVER_PORT"
	envServerTLSEnable         = "SERVER_TLS_ENABLE"
	envServerTLSCertFile       = "SERVER_TLS_CERT_FILE"
	envServerTLSKeyFile        = "SERVER_TLS_KEY_FILE"
	envBannerEnable            = "BANNER_ENABLE"
	envBannerType              = "BANNER_TYPE"
	envBannerText              = "BANNER_TEXT"
	envBannerFile              = "BANNER_FILE"
	envTemplateCache           = "TEMPLATE_CACHE"
	envTemplateRoot            = "TEMPLATE_ROOT"
	envTemplateSuffix          = "TEMPLATE_SUFFIX"
)

func (a *App) setServerMaxHeaderSize() {
	maxHeaderSize, err := intEnv(envServerMaxHeaderSize)
	if err == nil {
		a.Config.Server.MaxHeaderSize = maxHeaderSize
	}
}

func (a *App) setServerReadTimeout() {
	readTimeout, err := durationEnv(envServerReadTimeout)
	if err == nil {
		a.Config.Server.ReadTimeout = readTimeout
	}
}

func (a *App) setServerReadHeaderTimeout() {
	readHeaderTimeout, err := durationEnv(envServerReadHeaderTimeout)
	if err == nil {
		a.Config.Server.ReadHeaderTimeout = readHeaderTimeout
	}
}

func (a *App) setServerWriteTimeout() {
	writeTimeout, err := durationEnv(envServerWriteTimeout)
	if err == nil {
		a.Config.Server.WriteTimeout = writeTimeout
	}
}

func (a *App) setServerIdleTimeout() {
	idleTimeout, err := durationEnv(envServerIdleTimeout)
	if err == nil {
		a.Config.Server.IdleTimeout = idleTimeout
	}
}

func (a *App) setConfigFile() {
	configFile := stringEnv(envConfigFile)
	if configFile != "" {
		a.ConfigFile = configFile
	}
}

func (a *App) setServerHost() {
	host := stringEnv(envServerHost)
	if host != "" {
		a.Config.Server.Host = host
	}
}

func (a *App) setServerPort() {
	port, err := intEnv(envServerPort)
	if err == nil {
		a.Config.Server.Port = port
	}
}

func (a *App) setBannerEnable() {
	if stringEnv(envBannerEnable) != "" {
		a.Config.Banner.Enable = boolEnv(envBannerEnable)
	}
}

func (a *App) setBannerType() {
	bannerType := stringEnv(envBannerType)
	if bannerType != "" {
		a.Config.Banner.Type = bannerType
	}
}

func (a *App) setBannerText() {
	bannerText := stringEnv(envBannerText)
	if bannerText != "" {
		a.Config.Banner.Text = bannerText
	}
}

func (a *App) setBannerFile() {
	bannerFile := stringEnv(envBannerFile)
	if bannerFile != "" {
		a.Config.Banner.File = bannerFile
	}
}

func (a *App) setServerTLSEnable() {
	if stringEnv(envServerTLSEnable) != "" {
		a.Config.Server.TLS.Enable = boolEnv(envServerTLSEnable)
	}
}

func (a *App) setServerTLSCertFile() {
	serverTLSCertFile := stringEnv(envServerTLSCertFile)
	if serverTLSCertFile != "" {
		a.Config.Server.TLS.CertFile = serverTLSCertFile
	}
}

func (a *App) setServerTLSKeyFile() {
	serverTLSKeyFile := stringEnv(envServerTLSKeyFile)
	if serverTLSKeyFile != "" {
		a.Config.Server.TLS.KeyFile = serverTLSKeyFile
	}
}

func (a *App) setTemplateCache() {
	if stringEnv(envTemplateCache) != "" {
		a.Config.Template.Cache = boolEnv(envTemplateCache)
	}
}

func (a *App) setTemplateRoot() {
	templateRoot := stringEnv(envTemplateRoot)
	if templateRoot != "" {
		a.Config.Template.Root = templateRoot
	}
}

func (a *App) setTemplateSuffix() {
	templateSuffix := stringEnv(envTemplateSuffix)
	if templateSuffix != "" {
		a.Config.Template.Suffix = templateSuffix
	}
}

func (a *App) parseEnv() {
	a.setConfigFile()
	if a.ConfigFile != "" {
		a.parseYml()
	}
	a.setServerMaxHeaderSize()
	a.setServerReadTimeout()
	a.setServerReadHeaderTimeout()
	a.setServerWriteTimeout()
	a.setServerIdleTimeout()
	a.setServerHost()
	a.setServerPort()
	a.setServerTLSEnable()
	a.setServerTLSCertFile()
	a.setServerTLSKeyFile()
	a.setBannerEnable()
	a.setBannerType()
	a.setBannerText()
	a.setBannerFile()
	a.setTemplateCache()
	a.setTemplateRoot()
	a.setTemplateSuffix()
}

func stringEnv(key string) string {
	return os.Getenv(key)
}

func boolEnv(key string) bool {
	be := os.Getenv(key)
	b, err := strconv.ParseBool(be)
	return err == nil && b
}

func intEnv(key string) (int, error) {
	ie := os.Getenv(key)
	return strconv.Atoi(ie)
}

func durationEnv(key string) (time.Duration, error) {
	de := os.Getenv(key)
	return time.ParseDuration(de)
}
