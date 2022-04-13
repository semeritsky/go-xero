package xero

import (
	"context"
	"net/http"
)

type TenantsService service

func (s *TenantsService) GetTenants(ctx context.Context) ([]*Tenant, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "connections", nil)
	if err != nil {
		return nil, nil, err
	}

	var c []*Tenant
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, nil, err
	}

	return c, resp, nil
}

type Tenant struct {
	Id             string `json:"id"`
	AuthEventId    string `json:"authEventId"`
	TenantId       string `json:"tenantId"`
	TenantType     string `json:"tenantType"`
	TenantName     string `json:"tenantName"`
	CreatedDateUtc string `json:"createdDateUtc"`
	UpdatedDateUtc string `json:"updatedDateUtc"`
}
