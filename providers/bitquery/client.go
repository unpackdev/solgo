package bitquery

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type BitQueryProvider struct {
	ctx    context.Context
	opts   *Options
	client *http.Client
}

func NewBitQueryProvider(ctx context.Context, opts *Options) (*BitQueryProvider, error) {
	if opts == nil {
		return nil, errors.New("bitquery provider is not configured")
	}

	if opts.Endpoint == "" {
		return nil, errors.New("bitquery provider endpoint is not configured")
	}

	if opts.Key == "" {
		return nil, errors.New("bitquery provider key is not configured")
	}

	return &BitQueryProvider{
		ctx:    ctx,
		opts:   opts,
		client: &http.Client{Timeout: time.Second * 30},
	}, nil
}

func (b *BitQueryProvider) GetContractCreationInfo(ctx context.Context, query map[string]string) (*ContractCreationInfo, error) {
	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, errors.New("failed to marshal query: " + err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", b.opts.Endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, errors.New("failed to create request: " + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", b.opts.Key)

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, errors.New("failed to send request: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("received non-OK response: " + resp.Status)
	}

	var info ContractCreationInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, errors.New("failed to decode response: " + err.Error())
	}

	return &info, nil
}
