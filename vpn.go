package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const vpnPath = "/v2/vpns"

// VPNService is the interface to interact with the VPN endpoints on the Vultr API
// TODO: link
type VPNService interface {
	Create(ctx context.Context, vpnReq *VPNCreateReq) (*VPN, *http.Response, error)
	Get(ctx context.Context, vpnID string) (*VPN, *http.Response, error)
	Update(ctx context.Context, vpnID string, vpnReq *VPNUpdateReq) (*VPN, *http.Response, error)
	Delete(ctx context.Context, vpnID string) error
	List(ctx context.Context, options *ListOptions) ([]VPN, *Meta, *http.Response, error)

	CreateEndpoint(ctx context.Context, vpnID string, endpointReq *VPNEndpointReq) (*VPNEndpoint, *http.Response, error)
	GetEndpoint(ctx context.Context, vpnID, endpointID string) (*VPNEndpoint, *http.Response, error)
	UpdateEndpoint(ctx context.Context, vpnID, endpointID string, req *VPNEndpointReq) (*VPNEndpoint, *http.Response, error)
	DeleteEndpoint(ctx context.Context, vpnID, endpointID string) error
	ListEndpoints(ctx context.Context, vpnID string, options *ListOptions) ([]VPNEndpoint, *Meta, *http.Response, error)

	GetEndpointConfigFile(ctx context.Context, vpnID, endpointID string) (*VPNEndpointConfigFile, *http.Response, error)
}

// VPNService handles the interaction between the VPN service and the Vultr API.
type VPNServiceHandler struct {
	client *Client
}

// VPN represents a Vultr VPN subscriptions.
type VPN struct {
	ID           string     `json:"id"`
	VPCID        string     `json:"vpc_id"`
	NATGatewayID string     `json:"nat_gateway_id"`
	Label        string     `json:"label"`
	Status       string     `json:"status"`
	DateCreated  string     `json:"date_created"`
	Protocol     string     `json:"protocol"`
	Subnet       string     `json:"subnet"`
	SubnetPrefix int        `json:"subnet_prefix"`
	PublicKey    string     `json:"public_key"`
	Billing      VPNBilling `json:"billing"`
}

// VPNBilling represents the billing information for a VPN.
type VPNBilling struct {
	Charges     float32 `json:"charges"`
	Monthly     float32 `json:"monthly"`
	DateInvoice string  `json:"next_invoice_date"`
}

type vpnsBody struct {
	VPNs []VPN `json:"vpns"`
	Meta *Meta `json:"meta"`
}

type vpnBody struct {
	VPN *VPN `json:"vpn"`
}

// VPNCreateReq represents a VPN create request.
type VPNCreateReq struct {
	Label        string `json:"label"`
	Protocol     string `json:"protocol"`
	VPCID        string `json:"vpc_id"`
	Subnet       string `json:"subnet,omitempty"`
	SubnetPrefix int    `json:"subnet_prefix,omitempty"`
}

// VPNUpdateReq represents a VPN create request.
type VPNUpdateReq struct {
	Label string `json:"label"`
}

// VPNEndpoint represents a VPN endpoint.
type VPNEndpoint struct {
	ID            string                    `json:"id"`
	Label         string                    `json:"label"`
	DateCreated   string                    `json:"created_at"`
	DateUpdated   string                    `json:"updated_at"`
	RemoteSubnets []VPNEndpointRemoteSubnet `json:"remote_subnets"`
	Config        VPNEndpointConfig         `json:"config"`
}

// VPNEndpointRemoteSubnet represents a VPN endpoint remote subnet.
type VPNEndpointRemoteSubnet struct {
	Subnet string `json:"subnet"`
	Prefix int    `json:"subnet_prefix"`
}

// VPNEndpointConfig represents the config data of a VPN endpoint.
type VPNEndpointConfig struct {
	IPAddress        string `json:"ip_address"`
	PreSharedKey     string `json:"pre_shared_key"`
	RemoteIdentifier string `json:"remote_identifier"`
	PrivateKey       string `json:"private_key"`
	PublicKey        string `json:"public_key"`
	ServerPublicKey  string `json:"server_public_key"`
}

// VPNEndpointReq represents the body of a VPN endpoint create request.
type VPNEndpointReq struct {
	Label         string                    `json:"label"`
	IPAddress     string                    `json:"ip_address,omitempty"`
	RemoteSubnets []VPNEndpointRemoteSubnet `json:"remote_subnets,omitemtpy"`
	PreSharedKey  string                    `json:"pre_shared_key,omitempty"`
}

type vpnEndpointsBody struct {
	Endpoints []VPNEndpoint `json:"endpoints"`
	Meta      *Meta         `json:"meta"`
}

type vpnEndpointBody struct {
	Endpoint *VPNEndpoint `json:"endpoint"`
}

// VPNEndpointConfigFile represents the config for a VPN endpoint.
type VPNEndpointConfigFile struct {
	Config string `json:"config"`
}

// Create sends a create request for a VPN.
func (h *VPNServiceHandler) Create(ctx context.Context, vpnReq *VPNCreateReq) (*VPN, *http.Response, error) {
	req, errReq := h.client.NewRequest(ctx, http.MethodPost, vpnPath, vpnReq)
	if errReq != nil {
		return nil, nil, errReq
	}

	v := new(vpnBody)
	resp, errResp := h.client.DoWithContext(ctx, req, &v)
	if errResp != nil {
		return nil, resp, errResp
	}

	return v.VPN, resp, nil
}

// Get retrieves a VPN.
func (h *VPNServiceHandler) Get(ctx context.Context, vpnID string) (*VPN, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", vpnPath, vpnID)

	req, errReq := h.client.NewRequest(ctx, http.MethodGet, path, nil)
	if errReq != nil {
		return nil, nil, errReq
	}

	v := new(vpnBody)
	resp, errResp := h.client.DoWithContext(ctx, req, &v)
	if errResp != nil {
		return nil, resp, errResp
	}

	return v.VPN, resp, nil
}

// Update sends an update request for a VPN.
func (h *VPNServiceHandler) Update(ctx context.Context, vpnID string, vpnReq *VPNUpdateReq) (*VPN, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", vpnPath, vpnID)

	req, errReq := h.client.NewRequest(ctx, http.MethodPut, path, vpnReq)
	if errReq != nil {
		return nil, nil, errReq
	}

	v := new(vpnBody)
	resp, errResp := h.client.DoWithContext(ctx, req, &v)
	if errResp != nil {
		return nil, resp, errResp
	}

	return v.VPN, resp, nil
}

// Delete removes a VPN.
func (h *VPNServiceHandler) Delete(ctx context.Context, vpnID string) error {
	path := fmt.Sprintf("%s/%s", vpnPath, vpnID)

	req, errReq := h.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if errReq != nil {
		return errReq
	}

	if _, errResp := h.client.DoWithContext(ctx, req, nil); errResp != nil {
		return errResp
	}

	return nil
}

// List retrieves all VPNs.
func (h *VPNServiceHandler) List(ctx context.Context, options *ListOptions) ([]VPN, *Meta, *http.Response, error) {
	req, errReq := h.client.NewRequest(ctx, http.MethodGet, vpnPath, nil)
	if errReq != nil {
		return nil, nil, nil, errReq
	}

	qStrings, errQ := query.Values(options)
	if errQ != nil {
		return nil, nil, nil, errQ
	}

	req.URL.RawQuery = qStrings.Encode()

	vpnList := new(vpnsBody)
	resp, errResp := h.client.DoWithContext(ctx, req, &vpnList)
	if errResp != nil {
		return nil, nil, resp, errResp
	}

	return vpnList.VPNs, vpnList.Meta, resp, nil
}

// CreateEndpoint sends a create request for a VPN endpoint.
func (h *VPNServiceHandler) CreateEndpoint(ctx context.Context, vpnID string, endpointReq *VPNEndpointReq) (*VPNEndpoint, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/endpoints", vpnPath, vpnID)

	req, errReq := h.client.NewRequest(ctx, http.MethodPost, path, endpointReq)
	if errReq != nil {
		return nil, nil, errReq
	}

	e := new(vpnEndpointBody)
	resp, errResp := h.client.DoWithContext(ctx, req, &e)
	if errResp != nil {
		return nil, resp, errResp
	}

	return e.Endpoint, resp, nil
}

// GetEndpoint retrieves a VPN endpoint.
func (h *VPNServiceHandler) GetEndpoint(ctx context.Context, vpnID, endpointID string) (*VPNEndpoint, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/endpoints/%s", vpnPath, vpnID, endpointID)

	req, errReq := h.client.NewRequest(ctx, http.MethodGet, path, nil)
	if errReq != nil {
		return nil, nil, errReq
	}

	e := new(vpnEndpointBody)
	resp, errResp := h.client.DoWithContext(ctx, req, &e)
	if errResp != nil {
		return nil, resp, errResp
	}

	return e.Endpoint, resp, nil
}

// UpdateEndpoint sends an update request for a VPN endpoint.
func (h *VPNServiceHandler) UpdateEndpoint(ctx context.Context, vpnID, endpointID string, vpnReq *VPNEndpointReq) (*VPNEndpoint, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/endpoints/%s", vpnPath, vpnID, endpointID)

	req, errReq := h.client.NewRequest(ctx, http.MethodPut, path, vpnReq)
	if errReq != nil {
		return nil, nil, errReq
	}

	e := new(vpnEndpointBody)
	resp, errResp := h.client.DoWithContext(ctx, req, &e)
	if errResp != nil {
		return nil, resp, errResp
	}

	return e.Endpoint, resp, nil
}

// DeleteEndpoint removes a VPN endpoint.
func (h *VPNServiceHandler) DeleteEndpoint(ctx context.Context, vpnID, endpointID string) error {
	path := fmt.Sprintf("%s/%s/endpoints/%s", vpnPath, vpnID, endpointID)

	req, errReq := h.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if errReq != nil {
		return errReq
	}

	if _, errResp := h.client.DoWithContext(ctx, req, nil); errResp != nil {
		return errResp
	}

	return nil
}

// ListEndpoints retrieves all VPN endpoints.
func (h *VPNServiceHandler) ListEndpoints(ctx context.Context, vpnID string, options *ListOptions) ([]VPNEndpoint, *Meta, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/endpoints", vpnPath, vpnID)

	req, errReq := h.client.NewRequest(ctx, http.MethodGet, path, nil)
	if errReq != nil {
		return nil, nil, nil, errReq
	}

	qStrings, errQ := query.Values(options)
	if errQ != nil {
		return nil, nil, nil, errQ
	}

	req.URL.RawQuery = qStrings.Encode()

	e := new(vpnEndpointsBody)
	resp, errResp := h.client.DoWithContext(ctx, req, &e)
	if errResp != nil {
		return nil, nil, resp, errResp
	}

	return e.Endpoints, e.Meta, resp, nil
}

// GetEndpointConfigFile retrieves the config file used to connect to a VPN endpoint.
func (h *VPNServiceHandler) GetEndpointConfigFile(ctx context.Context, vpnID, endpointID string) (*VPNEndpointConfigFile, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/endpoints/%s/config", vpnPath, vpnID, endpointID)

	req, errReq := h.client.NewRequest(ctx, http.MethodGet, path, nil)
	if errReq != nil {
		return nil, nil, errReq
	}

	cfg := new(VPNEndpointConfigFile)
	resp, errResp := h.client.DoWithContext(ctx, req, &cfg)
	if errResp != nil {
		return nil, resp, errResp
	}

	return cfg, resp, nil
}
