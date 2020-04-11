/*
 * This file is part of the CDI project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2018 Red Hat, Inc.
 *
 */

package uploadserver

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"kubevirt.io/containerized-data-importer/pkg/common"
	"kubevirt.io/containerized-data-importer/pkg/importer"
	"kubevirt.io/containerized-data-importer/pkg/util/cert"
	"kubevirt.io/containerized-data-importer/pkg/util/cert/triple"
	"kubevirt.io/containerized-data-importer/tests/reporters"
)

func TestUploadServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Upload Server Suite", reporters.NewReporters())
}

func newServer() *uploadServerApp {
	server := NewUploadServer("127.0.0.1", 0, "disk.img", "", "", "", "", "")
	return server.(*uploadServerApp)
}

func newTLSServer(clientCertName, expectedName string) (*uploadServerApp, *triple.KeyPair, *x509.Certificate) {
	serverCA, err := triple.NewCA("server")
	Expect(err).ToNot(HaveOccurred())

	clientCA, err := triple.NewCA("client")
	Expect(err).ToNot(HaveOccurred())

	serverKeyPair, err := triple.NewServerKeyPair(serverCA, "localhost", "localhost", "default", "local", []string{"127.0.0.1"}, []string{"localhost"})
	Expect(err).ToNot(HaveOccurred())

	tlsKey := string(cert.EncodePrivateKeyPEM(serverKeyPair.Key))
	tlsCert := string(cert.EncodeCertPEM(serverKeyPair.Cert))
	clientCert := string(cert.EncodeCertPEM(clientCA.Cert))

	server := NewUploadServer("127.0.0.1", 0, "disk.img", tlsKey, tlsCert, clientCert, expectedName, "").(*uploadServerApp)

	clientKeyPair, err := triple.NewClientKeyPair(clientCA, clientCertName, []string{})
	Expect(err).ToNot(HaveOccurred())

	return server, clientKeyPair, serverCA.Cert
}

func newHTTPClient(clientKeyPair *triple.KeyPair, serverCACert *x509.Certificate) *http.Client {
	clientCert, err := tls.X509KeyPair(cert.EncodeCertPEM(clientKeyPair.Cert), cert.EncodePrivateKeyPEM(clientKeyPair.Key))
	Expect(err).ToNot(HaveOccurred())

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert.EncodeCertPEM(serverCACert))

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	return client
}

func saveProcessorSuccess(stream io.ReadCloser, dest, imageSize, contentType string) error {
	return nil
}

func saveProcessorFailure(stream io.ReadCloser, dest, imageSize, contentType string) error {
	return fmt.Errorf("Error using datastream")
}

func withProcessorSuccess(f func()) {
	replaceProcessorFunc(saveProcessorSuccess, f)
}

func withProcessorFailure(f func()) {
	replaceProcessorFunc(saveProcessorFailure, f)
}

func replaceProcessorFunc(replacement func(io.ReadCloser, string, string, string) error, f func()) {
	origProcessorFunc := uploadProcessorFunc
	uploadProcessorFunc = replacement
	defer func() {
		uploadProcessorFunc = origProcessorFunc
	}()
	f()
}

type AsyncMockDataSource struct {
}

// Info is called to get initial information about the data.
func (amd *AsyncMockDataSource) Info() (importer.ProcessingPhase, error) {
	return importer.ProcessingPhaseTransferDataFile, nil
}

// Transfer is called to transfer the data from the source to the passed in path.
func (amd *AsyncMockDataSource) Transfer(path string) (importer.ProcessingPhase, error) {
	return importer.ProcessingPhasePause, nil
}

// TransferFile is called to transfer the data from the source to the passed in file.
func (amd *AsyncMockDataSource) TransferFile(fileName string) (importer.ProcessingPhase, error) {
	return importer.ProcessingPhasePause, nil
}

// Process is called to do any special processing before giving the url to the data back to the processor
func (amd *AsyncMockDataSource) Process() (importer.ProcessingPhase, error) {
	return importer.ProcessingPhaseConvert, nil
}

// Close closes any readers or other open resources.
func (amd *AsyncMockDataSource) Close() error {
	return nil
}

// GetURL returns the url that the data processor can use when converting the data.
func (amd *AsyncMockDataSource) GetURL() *url.URL {
	return nil
}

// GetResumePhase returns the next phase to process when resuming
func (amd *AsyncMockDataSource) GetResumePhase() importer.ProcessingPhase {
	return importer.ProcessingPhaseComplete
}

func saveAsyncProcessorSuccess(stream io.ReadCloser, dest, imageSize, contentType string) (*importer.DataProcessor, error) {
	return importer.NewDataProcessor(&AsyncMockDataSource{}, "", "", "", ""), nil
}

func saveAsyncProcessorFailure(stream io.ReadCloser, dest, imageSize, contentType string) (*importer.DataProcessor, error) {
	return importer.NewDataProcessor(&AsyncMockDataSource{}, "", "", "", ""), fmt.Errorf("Error using datastream")
}

func withAsyncProcessorSuccess(f func()) {
	replaceAsyncProcessorFunc(saveAsyncProcessorSuccess, f)
}

func withAsyncProcessorFailure(f func()) {
	replaceAsyncProcessorFunc(saveAsyncProcessorFailure, f)
}

func replaceAsyncProcessorFunc(replacement func(io.ReadCloser, string, string, string) (*importer.DataProcessor, error), f func()) {
	origProcessorFuncAsync := uploadProcessorFuncAsync
	uploadProcessorFuncAsync = replacement
	defer func() {
		uploadProcessorFuncAsync = origProcessorFuncAsync
	}()
	f()
}

var _ = Describe("Upload server tests", func() {
	It("GET fails", func() {
		withProcessorSuccess(func() {
			req, err := http.NewRequest("GET", common.UploadPathSync, nil)
			Expect(err).ToNot(HaveOccurred())

			rr := httptest.NewRecorder()

			server := newServer()
			server.ServeHTTP(rr, req)

			status := rr.Code
			Expect(status).To(Equal(http.StatusNotFound))
		})
	})

	It("healthz", func() {
		req, err := http.NewRequest("GET", healthzPath, nil)
		Expect(err).ToNot(HaveOccurred())

		rr := httptest.NewRecorder()

		app := uploadServerApp{}
		server, _ := app.createHealthzServer()
		server.Handler.ServeHTTP(rr, req)

		status := rr.Code
		Expect(status).To(Equal(http.StatusOK))

	})

	table.DescribeTable("Process unavailable", func(uploadPath string) {
		withProcessorSuccess(func() {
			req, err := http.NewRequest("POST", common.UploadPathAsync, strings.NewReader("data"))
			Expect(err).ToNot(HaveOccurred())

			rr := httptest.NewRecorder()

			server := newServer()
			server.uploading = true
			server.ServeHTTP(rr, req)

			status := rr.Code
			Expect(status).To(Equal(http.StatusServiceUnavailable))
		})
	},
		table.Entry("async", common.UploadPathAsync),
		table.Entry("sync", common.UploadPathSync),
	)

	table.DescribeTable("Completion conflict", func(uploadPath string) {
		withAsyncProcessorSuccess(func() {
			req, err := http.NewRequest("POST", uploadPath, strings.NewReader("data"))
			Expect(err).ToNot(HaveOccurred())

			rr := httptest.NewRecorder()

			server := newServer()
			server.done = true
			server.ServeHTTP(rr, req)

			status := rr.Code
			Expect(status).To(Equal(http.StatusConflict))
		})
	},
		table.Entry("async", common.UploadPathAsync),
		table.Entry("sync", common.UploadPathSync),
	)

	It("Success", func() {
		withProcessorSuccess(func() {
			req, err := http.NewRequest("POST", common.UploadPathSync, strings.NewReader("data"))
			Expect(err).ToNot(HaveOccurred())

			rr := httptest.NewRecorder()

			server := newServer()
			server.ServeHTTP(rr, req)

			status := rr.Code
			Expect(status).To(Equal(http.StatusOK))
		})
	})

	table.DescribeTable("Success, async", func(method string) {
		withAsyncProcessorSuccess(func() {
			req, err := http.NewRequest(method, common.UploadPathAsync, strings.NewReader("data"))
			Expect(err).ToNot(HaveOccurred())

			rr := httptest.NewRecorder()

			server := newServer()
			server.ServeHTTP(rr, req)

			status := rr.Code
			Expect(status).To(Equal(http.StatusOK))
		})
	},
		table.Entry("POST", "POST"),
		table.Entry("HEAD", "HEAD"),
	)

	table.DescribeTable("Stream fail", func(uploadPath string) {
		withAsyncProcessorFailure(func() {
			req, err := http.NewRequest("POST", uploadPath, strings.NewReader("data"))
			Expect(err).ToNot(HaveOccurred())

			rr := httptest.NewRecorder()

			server := newServer()
			server.ServeHTTP(rr, req)

			status := rr.Code
			Expect(status).To(Equal(http.StatusInternalServerError))
		})
	},
		table.Entry("async", common.UploadPathAsync),
		table.Entry("sync", common.UploadPathSync),
	)

	table.DescribeTable("Real upload with client", func(certName string, expectedName string, expectedResponse int) {
		withProcessorSuccess(func() {
			server, clientKeyPair, serverCACert := newTLSServer(certName, expectedName)

			client := newHTTPClient(clientKeyPair, serverCACert)

			ch := make(chan struct{})

			go func() {
				server.Run()
				close(ch)
			}()

			for i := 0; i < 10; i++ {
				if server.bindPort != 0 {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}

			Expect(server.bindPort).ToNot(Equal(0))

			url := fmt.Sprintf("https://localhost:%d%s", server.bindPort, common.UploadPathSync)
			stringReader := strings.NewReader("nothing")

			resp, err := client.Post(url, "application/x-www-form-urlencoded", stringReader)
			Expect(err).ToNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(expectedResponse))

			if !server.done {
				close(server.doneChan)
			}

			<-ch
		})
	},
		table.Entry("Valid data", "client", "client", 200),
		table.Entry("Invalid data", "foo", "bar", 401),
	)
})
