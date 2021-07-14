package zone

import (
	"context"
	"encoding/json"
	"fmt"
	"k8s.io/klog"
	"net/http"
	"strings"
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

func (a api) List(ctx context.Context, zoneName string) ([]Summary, error) {
	klog.Infof("baseurl: %s; pathprefix: %s; zonename: %s", a.client.BaseURL(), pathPrefix, zoneName)
	url := fmt.Sprintf(
		"%s%s/%s/records",
		a.client.BaseURL(),
		pathPrefix, zoneName,
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

func (a api) AddRecord(ctx context.Context, zoneName, jsonString string) error {
	url := fmt.Sprintf(
		"%s%s/%s/records",
		a.client.BaseURL(),
		pathPrefix, zoneName,
	)

	/* example obj
	{
	    "type": "A",
	    "rdata": "9.9.9.9"
	}
	*/

	//reader, _ := client.NewReader(ctx,obj)
	reader := strings.NewReader(jsonString) //falscher reader ???? strings Reader !== io Reader
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reader)
	if err != nil {
		return fmt.Errorf("could not attach record post request: %w", err)
	}

	httpResponse, err := a.client.Do(req)
	if httpResponse.StatusCode >= 500 && httpResponse.StatusCode < 600 {
		return fmt.Errorf("could not execute attach record request, got response %s", httpResponse.Status)
	}

	var summary []Summary
	err = json.NewDecoder(httpResponse.Body).Decode(&summary)
	_ = httpResponse.Body.Close()
	if err != nil {
		return fmt.Errorf("could not decode attach record response: %w", err)
	}

	return nil
}

func (a api) RemoveRecord(ctx context.Context, zoneName, recordId string) error {
	url := fmt.Sprintf(
		"%s%s/%s/records/%s",
		a.client.BaseURL(), pathPrefix, zoneName, recordId,
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
