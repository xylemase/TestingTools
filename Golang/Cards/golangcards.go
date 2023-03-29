package card_api

import (
	"net/http"
	"net/http/httptest"
	"testing"
        "io"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func TestAPI(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "API Suite")
}

var _ = ginkgo.Describe("API", func() {
	var (
		server *httptest.Server
	)

	ginkgo.BeforeEach(func() {
		// Set up the server with your API routes and handlers
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, World!"))
		})

		server = httptest.NewServer(handler)
	})

	ginkgo.AfterEach(func() {
		server.Close()
	})

	ginkgo.It("returns a 200 status code", func() {
		resp, err := http.Get(server.URL)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(resp.StatusCode).To(gomega.Equal(http.StatusOK))
	})

	ginkgo.It("returns the correct response body", func() {
		resp, err := http.Get(server.URL)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		bodyBytes, err := io.ReadAll(resp.Body)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(string(bodyBytes)).To(gomega.Equal("Hello, World!"))
	})
})
