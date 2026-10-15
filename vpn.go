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
	Create(ctx context.Context, req *VPNCreateReq) (*VPN, *http.Response, error)
	Get(ctx context.Context, vpnID string) (*VPN, *http.Response, error)
	Update(ctx context.Context, vpnID string, req *VPNUpdateReq) (*VPN, *http.Response, error)
	Delete(ctx context.Context, vpnID string) error
	List(ctx context.Context, options *ListOptions) ([]VPN, *Meta, *http.Response, error)
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

// Create sends a create request for a VPN
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

// Get retrieves a VPN
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

// Update sends an update request for a VPN
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

// Delete removes a VPN
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

// List retrieves all VPNs
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
