package tests_test

import (
	"context"
	"github.com/anexia-it/go-anxcloud/pkg/clouddns/zone"
	"os"
	"time"

	"github.com/anexia-it/go-anxcloud/pkg/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Clouddns API endpoint tests", func() {

	var cli client.Client

	BeforeEach(func() {
		var err error
		os.Setenv("ANEXIA_TOKEN", "tNSwT-Fnp0Ac2mQv1ntt2OBraSnw7Atu")
		cli, err = client.New(client.AuthFromEnv(false))
		Expect(err).ToNot(HaveOccurred())
	})

	Context("Zones endpoint", func() {

		It("Should list all created records", func() {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
			defer cancel()
			_, err := zone.NewAPI(cli).List(ctx, "dev.se.anx.io")
			Expect(err).NotTo(HaveOccurred())
		})

	})
})
