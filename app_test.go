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
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestAppRun(t *testing.T) {
	errCh := make(chan error, 1)
	okCh := make(chan bool, 1)
	go func() {
		New().Run()
	}()
	time.AfterFunc(time.Millisecond*100, func() {
		http.DefaultClient.Timeout = time.Second
		response, err := http.Get("http://localhost:9494")
		if err != nil {
			errCh <- err
		} else if response.StatusCode > 0 {
			okCh <- true
		}
	})
	select {
	case <-time.After(time.Second):
		t.Error("test fail: timeout")
	case err := <-errCh:
		t.Errorf("test err: %v", err)
	case <-okCh:
		t.Log("test ok")
	}
}

func TestAppRunWithTLS(t *testing.T) {
	defer func() {
		if re := recover(); re != nil {
			t.Log("test ok!")
		}
	}()
	_ = ioutil.WriteFile("ca.crt", []byte(certFile), 0700)
	_ = ioutil.WriteFile("ca.key", []byte(keyFile), 0700)
	defer func() {
		_ = os.Remove("ca.crt")
		_ = os.Remove("ca.key")
	}()
	go func() {
		app := New()
		app.Config.Server.TLS.Enable = true
		app.Config.Server.TLS.CertFile = "ca.crt"
		app.Config.Server.TLS.KeyFile = "ca.key"
		app.Config.Server.Port = 9595
		app.Run()
	}()
	time.Sleep(time.Second)
	conn, _ := net.DialTimeout("tcp", "localhost:9595", time.Second)
	require.NotNil(t, conn)
	_ = conn.Close()
}

const (
	certFile = `-----BEGIN CERTIFICATE-----
MIIDETCCAfkCFHyMmBP9DYIZRoqW17cQvSupfKISMA0GCSqGSIb3DQEBCwUAMEUx
CzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRl
cm5ldCBXaWRnaXRzIFB0eSBMdGQwHhcNMjExMjE4MDEyNzQ1WhcNMzExMjE2MDEy
NzQ1WjBFMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UE
CgwYSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAxDXnGF0QWnNtRA4TL8uAhDE0tuorSwq6DZI2rCXYkaY4+6SN
WYcGbpLTnJVP0yRBbf+8kI04yGWGFObmfcyUniDS4N0QrKcdTE31Egvc+vb5RKbj
Ht0jcj/ofyWSb50xZUaGolqKJ9rb3A0Y4ia9TnLGnf6KSDlYAvjfu7wY1RGFkgRx
HnpxKGT/Ry2J+pSWYrVOYOHub2M2eI2RLzY6jrEHWkif+Bybtwt84qIWG6fKkgFC
d7hQwbaqiD6AuwuUtsSzBzJ5S0cOLvuCHLG0dRVcLkkHkirLJ1WMEfzC/SlsOHCW
hS92v6IH72UbPO2nHGskikbtnQqC8CnKfiq1XwIDAQABMA0GCSqGSIb3DQEBCwUA
A4IBAQAiPJ9nDjqTsCbq4pgfNaVFNY+bGgi2cGyiD1+MOGxp8BycfLUfTmHiBUPM
OmOhg5hXxiknJ5eAhRBP+c6NlemoT61oTUmsr1jmTpoiryBZmnZnKCu8TxkxhG/r
Mk9ZwXJdtmO8iJT/9vBF7i4vHCUBMimZztowsH1azds3trJtmINQOORiXeoIbs38
tfxRK0ScDrr+dt+MrLNThasdHBpMjfI9NvnqjlPmDHzY+nQicotkmGASnyQ1F4K9
CuhVzAi5ud33TpIH6CsrINIbhHnybGyGfWmtFpJXf7Wnl6Bja5IJBaOi1ODGAdlK
eehq4+5HPiABnRBh+B6APRB/qodg
-----END CERTIFICATE-----`
	keyFile = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAxDXnGF0QWnNtRA4TL8uAhDE0tuorSwq6DZI2rCXYkaY4+6SN
WYcGbpLTnJVP0yRBbf+8kI04yGWGFObmfcyUniDS4N0QrKcdTE31Egvc+vb5RKbj
Ht0jcj/ofyWSb50xZUaGolqKJ9rb3A0Y4ia9TnLGnf6KSDlYAvjfu7wY1RGFkgRx
HnpxKGT/Ry2J+pSWYrVOYOHub2M2eI2RLzY6jrEHWkif+Bybtwt84qIWG6fKkgFC
d7hQwbaqiD6AuwuUtsSzBzJ5S0cOLvuCHLG0dRVcLkkHkirLJ1WMEfzC/SlsOHCW
hS92v6IH72UbPO2nHGskikbtnQqC8CnKfiq1XwIDAQABAoIBACs+ZgxslmoY/n/9
SiVCiLSZ07Bss9X6Kz9KdlpCjRSsuepcPfr5U2WTXqgoEEvMtc70ii6hsV4ZYg/B
RBN9v1OKkG+WyVIEEuT6WYT8sFtvi0iiL3Rh8KoBg9BiC4Al+PkFLi8iHUjjZ4l8
KXvOZfKgQT4ZF4kLemZNS6IotqBetR7jrKPLdYJrH5quZvZliFyTIL8O1EOaQ/EJ
mtYPX/C7e/2hcj9Fhix6t0WoiT3KwJsIlw4QCNmlRBpKR64ZT6Z6UFUTpayrP+ev
KGoXCN8rVOZXc9VaAsWD7/rzIhsdMIxxoybG3WD8knMJgTv2DGZtoveAzLi3vZG/
INFaoEECgYEA8ykfG5FPE12HNQJQPdnvxgdTX0//OmsQKlDx3KpCzgcprYn7WCiF
SL2gigZ0dK14vQPltolUUzeTnBPwyo8wtgS2yX1XuD2zknhOTYOBz3qJXgnlF+6B
1Iw8UurdPjBQZv+tghrCD55Tch7bagk16sP3i7gO7bDhczPHr3T/tAkCgYEAzpIk
cEm/OqqZIN69pjhieomLPe7K9y6IXkNM8ZDARggkOj7xVSBYD4SjhuGR0vHDmvU8
5qEPkwa7YLMlDdxvgTB+zCDKC466v6Y+5FLemMyIYDr/3YwE+29hJ5c5F8unrw5U
EO8/deD4maZbQGIGR8lZshw8mpZesCTMZuJmCCcCgYEAqRLht2htJFj3B2vJYYhl
CTvUw8Q7AmKpRdMsqTOV6e3fE/SKWL0sF+0KcI6WcP6hokPQeQC6KnbNY0wWNLIl
u5pBgo5t7QSyFNkkEQ+sthhM5Z9ZtS85BRJRa5I1LeWoMkX7Xii+4N9ExGgiRnOL
EucZ/AOKFcnUqSbK5PwkRAkCgYAUEl84CfJq4OjAKOSEojXvci31dp6CJiNaBXAU
iNwl8eSTREpu2xWzbE/3azOgK522EN46CqxYvO64FrAjCKhNBUlMzGLVfKjotl6m
EOdQMY+OyizSeiiBxfDKyAbkKQXCHMJOYvDno1SEmYWEXAIAN7Bffh7lZncM5oZ1
+MmxQQKBgQCHLD8XXmB6mlWiwr4zTNUts2XPRgiDjp4kDcFbNkgOsbJGwwBAWlsU
KjOVgp30TxXY/+XTc8h/qG5LeBga9Ul3eW6hoCr6l4Sdez+JeYRn0r8JIX22ClL3
lB0CQqCfRU9GQjCMkgx/FpQ0Al1w70tDuDspjQxZgnPxp7/Nv4YKAA==
-----END RSA PRIVATE KEY-----`
)
