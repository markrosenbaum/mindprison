/*==============================================================================
   Copyright 2018 Jeff Sharpe

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License
=============================================================================*/

package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"mindprison/internal/endpoint"
)

type Endpoint struct {
	endpoint.BaseEndpoint
	server *http.Server

	started *sync.Mutex
}

func NewEndpoint(port int) *Endpoint {

	startMutex := sync.Mutex{}
	startMutex.Lock()
	e := Endpoint{BaseEndpoint: endpoint.BaseEndpoint{Port: port}, started: &startMutex, server: nil}

	return &e
}

func (e *Endpoint) Start() {
	e.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", e.Port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      e,
	}
	log.Printf("Configuring HTTP listerner for port %d", e.Port)

	go func() {
		log.Printf("Endpoint[%d]: Starting HTTP listener", e.Port)
		go func() {
			err := e.server.ListenAndServe()

			if err != nil {
				log.Printf("Error managing HTTP endpoint: %v", err)
			}
		}()
		time.Sleep(100 * time.Millisecond)
		e.started.Unlock()
	}()
}

func (e *Endpoint) Stop() {
	shutdownContext, _ := context.WithTimeout(context.Background(), 10*time.Second)
	e.server.Shutdown(shutdownContext)
}

func (e *Endpoint) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	log.Printf("Endpoint[%d]: Got HTTP request for %s", e.Port, request.URL)

	response.WriteHeader(http.StatusTeapot)
	response.Write([]byte("Payload"))
}

func (e *Endpoint) WaitForStartup() {
	e.started.Lock()
	e.started.Unlock()
}
