package tests_test

import (
	"context"
	"fmt"
	"github.com/anexia-it/go-anxcloud/pkg/client"
	"github.com/anexia-it/go-anxcloud/pkg/clouddns/zone"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"os"
	"time"
)

var _ = Describe("Zone", func() {

	var cli client.Client
	const zonename = "dev.se.anx.io"

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
			_, err := zone.NewAPI(cli).List(ctx, zonename)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Should add records", func() {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
			defer cancel()
			minKeyName := 1
			maxKeyName := 999999
			minIP := 1
			maxIP := 250
			var jsonStr = fmt.Sprintf(`{"name":"testrecord%d", "type":"A", "rdata":"%d.%d.%d.%d", "region":"default", "ttl":3600}`, rand.Intn(maxKeyName-minKeyName)+minKeyName, rand.Intn(maxIP-minIP)+minIP, rand.Intn(maxIP-minIP)+minIP, rand.Intn(maxIP-minIP)+minIP, rand.Intn(maxIP-minIP)+minIP)
			err := zone.NewAPI(cli).AddRecord(ctx, zonename, jsonStr)
			Expect(err).NotTo(HaveOccurred())
		})

	})

})
