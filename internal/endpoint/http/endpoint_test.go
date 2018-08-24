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
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func buildTestRequest(port int, method string, path string) *http.Request {
	testReq, reqError := http.NewRequest(method, fmt.Sprintf("http://localhost:%d%s", port, path), &bytes.Buffer{})

	if reqError != nil {
		log.Fatalf("Error creating test URL: %v", reqError)
	}

	return testReq
}

func TestEndpoint(t *testing.T) {
	port := 6666
	client := http.Client{}

	e := NewEndpoint(6666)

	log.Printf("Starting Test HTTP Endpoint.")
	e.Start()
	e.WaitForStartup()
	log.Printf("Test HTTP endpoint started.")

	testReq := buildTestRequest(port, "GET", "/test/path?var=val")

	resp, respErr := client.Do(testReq)

	require.NoError(t, respErr, "Error talking to HTTP endpoint.")

	log.Printf("Response: %d %s (%d bytes)", resp.StatusCode, resp.Status, resp.ContentLength)

	time.Sleep(4 * time.Second)

	e.Stop()
	log.Printf("Test HTTP endpoint shut down.")
}
