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
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-the-way/anoweb/config"
	"github.com/go-the-way/anoweb/middleware"
	"github.com/go-the-way/anoweb/rest"
	"github.com/go-the-way/anoweb/router"
)

// App struct
type App struct {
	logger         *log.Logger
	ConfigFile     string
	Config         *config.Config
	controllers    []rest.Controller
	groups         []*router.Group
	routers        []*router.Router
	parsedRouters  *router.ParsedRouter
	middlewares    []middleware.Middleware
	defaultMWState *defaultMWState
}

// Default the default App
var Default = New()

// New return new App
func New() *App {
	return &App{
		logger:         log.New(os.Stdout, "[anoweb] ", log.LstdFlags),
		ConfigFile:     "app.yml",
		Config:         config.Default(),
		controllers:    make([]rest.Controller, 0),
		groups:         make([]*router.Group, 0),
		routers:        []*router.Router{router.NewRouter()},
		parsedRouters:  &router.ParsedRouter{Simples: make(router.SimpleM), Dynamics: make(router.DynamicM)},
		middlewares:    make([]middleware.Middleware, 6),
		defaultMWState: &defaultMWState{header: true, faviconFile: "favicon.ico", faviconRoute: "/favicon.ico"},
	}
}

// Run App
func (a *App) Run() {
	a.parseYml()
	a.parseEnv()
	a.printBanner()
	a.printVendor()
	a.routeRestControllers()
	a.useDefaultMWs()
	a.parseRouters()
	a.serve()
}

func (a *App) serve() {
	host := a.Config.Server.Host
	port := a.Config.Server.Port
	tlsEnable := a.Config.Server.TLS.Enable
	addr := host + ":" + strconv.Itoa(port)
	server := &http.Server{
		Addr:              addr,
		Handler:           a.newDispatcher(),
		MaxHeaderBytes:    a.Config.Server.MaxHeaderSize,
		ReadTimeout:       a.Config.Server.ReadTimeout,
		ReadHeaderTimeout: a.Config.Server.ReadHeaderTimeout,
		WriteTimeout:      a.Config.Server.WriteTimeout,
		IdleTimeout:       a.Config.Server.IdleTimeout,
	}
	if tlsEnable {
		server.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		certFile := a.Config.Server.TLS.CertFile
		keyFile := a.Config.Server.TLS.KeyFile
		a.logger.Printf("Server started on https://%s\n", addr)
		_, _ = fmt.Fprintln(os.Stderr, server.ListenAndServeTLS(certFile, keyFile))
	} else {
		a.logger.Printf("Server started on http://%s\n", addr)
		_, _ = fmt.Fprintln(os.Stderr, server.ListenAndServe())
	}
}
