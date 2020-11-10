// Package pkg contains all API functionality and helpers.
package pkg

import (
	"github.com/anexia-it/go-anxcloud/pkg/client"
	"github.com/anexia-it/go-anxcloud/pkg/ipam"
	"github.com/anexia-it/go-anxcloud/pkg/test"
	"github.com/anexia-it/go-anxcloud/pkg/vsphere"
)

// API contains all API calls structured their location in the API.
type API interface {
	IPAM() ipam.API
	Test() test.API
	VSphere() vsphere.API
}

type api struct {
	ipam    ipam.API
	test    test.API
	vsphere vsphere.API
}

func (a api) IPAM() ipam.API {
	return a.ipam
}

func (a api) Test() test.API {
	return a.test
}

func (a api) VSphere() vsphere.API {
	return a.vsphere
}

// NewAPI creates a new API instance with the given client.
func NewAPI(c client.Client) API {
	return &api{
		ipam.NewAPI(c),
		test.NewAPI(c),
		vsphere.NewAPI(c),
	}
}
