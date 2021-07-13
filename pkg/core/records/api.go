package records

import (
	"context"
	"github.com/anexia-it/go-anxcloud/pkg/client"
)

// API contains methods for tag control.
type API interface {
	List(ctx context.Context, zoneName string) ([]Summary, error)
	Get(ctx context.Context, id string) (Info, error)
	AddRecord(ctx context.Context, zoneName, recordType, recordData string) ([]Summary, error)
	RemoveRecord(ctx context.Context, zoneName, recordId string) error
}

type api struct {
	client client.Client
}

// NewAPI creates a new tags API instance with the given client.
func NewAPI(c client.Client) API {
	return api{c}
}
