package records

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const pathPrefix = "/api/clouddns/v1/zone.json"

// Summary describes a resource in short.
type Summary struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// Type is part of info.
type Type struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// Info contains all information about a resource.
type Info struct {
	Identifier string   `json:"identifier"`
	Name       string   `json:"name"`
	Type       Type     `json:"resource_type"`
	Tags       []string `json:"tags"`
}

type listResponse struct {
	Data []Summary `json:"data"`
}

func (a api) List(ctx context.Context) ([]Summary, error) {
	url := fmt.Sprintf(
		"%s%s/%s/records",
		a.client.BaseURL(),
		pathPrefix,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create record list request: %w", err)
	}

	httpResponse, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not execute record list request: %w", err)
	}
	if httpResponse.StatusCode >= 500 && httpResponse.StatusCode < 600 {
		return nil, fmt.Errorf("could not execute record list request, got response %s", httpResponse.Status)
	}

	var responsePayload listResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&responsePayload)
	_ = httpResponse.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not decode record list response: %w", err)
	}

	return responsePayload.Data, nil
}

func (a api) Get(ctx context.Context, id string) (Info, error) {
	url := fmt.Sprintf(
		"%s%s/%s",
		a.client.BaseURL(),
		pathPrefix,
		id,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Info{}, fmt.Errorf("could not create record get request: %w", err)
	}

	httpResponse, err := a.client.Do(req)
	if err != nil {
		return Info{}, fmt.Errorf("could not execute record get request: %w", err)
	}
	if httpResponse.StatusCode >= 500 && httpResponse.StatusCode < 600 {
		return Info{}, fmt.Errorf("could not execute record get request, got response %s", httpResponse.Status)
	}

	var info Info
	err = json.NewDecoder(httpResponse.Body).Decode(&info)
	_ = httpResponse.Body.Close()
	if err != nil {
		return Info{}, fmt.Errorf("could not decode record get response: %w", err)
	}

	return info, nil
}

func (a api) AddRecord(ctx context.Context, zonename string) ([]Summary, error) {
	url := fmt.Sprintf(
		"%s%s/%v/records",
		a.client.BaseURL(),
		pathPrefix, zonename,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not attach record post request: %w", err)
	}

	httpResponse, err := a.client.Do(req)
	if httpResponse.StatusCode >= 500 && httpResponse.StatusCode < 600 {
		return nil, fmt.Errorf("could not execute attach record request, got response %s", httpResponse.Status)
	}

	var summary []Summary
	err = json.NewDecoder(httpResponse.Body).Decode(&summary)
	_ = httpResponse.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not decode attach record response: %w", err)
	}

	return summary, nil
}

func (a api) RemoveRecord(ctx context.Context, zonename, recordId string) error {
	url := fmt.Sprintf(
		"%s%s/%v/records/%v",
		a.client.BaseURL(), zonename, recordId,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("could not create record delete request: %w", err)
	}

	httpResponse, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("could not execute record delete request: %w", err)
	}
	if httpResponse.StatusCode >= 500 && httpResponse.StatusCode < 600 {
		return fmt.Errorf("could not execute record delete request, got response %s", httpResponse.Status)
	}

	return httpResponse.Body.Close()
}
