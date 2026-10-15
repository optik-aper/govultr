package govultr

import (
	"net/http"
	"reflect"
	"testing"
)

func TestVPNServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpns", testJSONResponseHandlerFunc(http.StatusCreated, `
{
	"vpn": {
        "id": "a993ed18-9fc2-4217-a1a4-e860bddc3684",
        "vpc_id": "7ceab21c-e35b-4efb-9e96-6026ff5de5aa",
        "nat_gateway_id": "ab75b684-b3ca-443d-9550-65ce812cce6d",
        "label": "govultr-test",
        "status": "active",
        "date_created": "2026-10-13 12:36:04",
        "protocol": "ipsec",
        "subnet": "10.0.16.0",
        "subnet_prefix": 20,
        "public_key": "",
        "billing": {
                "charges": 0.04,
                "monthly": 25
        }
	}
}
`))

	vpnReq := &VPNCreateReq{
		Label:    "govultr-test",
		Protocol: "ipsec",
		VPCID:    "7ceab21c-e35b-4efb-9e96-6026ff5de5aa",
	}
	v, _, err := client.VPN.Create(ctx, vpnReq)
	if err != nil {
		t.Errorf("VPN.Create returned %+v, expected %+v", err, nil)
	}

	expected := &VPN{
		ID:           "a993ed18-9fc2-4217-a1a4-e860bddc3684",
		VPCID:        "7ceab21c-e35b-4efb-9e96-6026ff5de5aa",
		NATGatewayID: "ab75b684-b3ca-443d-9550-65ce812cce6d",
		Label:        "govultr-test",
		Status:       "active",
		DateCreated:  "2026-10-13 12:36:04",
		Protocol:     "ipsec",
		Subnet:       "10.0.16.0",
		SubnetPrefix: 20,
		PublicKey:    "",
		Billing: VPNBilling{
			Charges: 0.04,
			Monthly: 25,
		},
	}

	if !reflect.DeepEqual(v, expected) {
		t.Errorf("ISO.Create returned %+v, expected %+v", v, expected)
	}
}

func TestVPNServiceHandler_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpns/a993ed18-9fc2-4217-a1a4-e860bddc3684", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"vpn": {
        "id": "a993ed18-9fc2-4217-a1a4-e860bddc3684",
        "vpc_id": "7ceab21c-e35b-4efb-9e96-6026ff5de5aa",
        "nat_gateway_id": "ab75b684-b3ca-443d-9550-65ce812cce6d",
        "label": "govultr-test",
        "status": "active",
        "date_created": "2026-10-13 12:36:04",
        "protocol": "ipsec",
        "subnet": "10.0.16.0",
        "subnet_prefix": 20,
        "public_key": "",
        "billing": {
                "charges": 0.04,
                "monthly": 25
        }
	}
}
`))

	v, _, err := client.VPN.Get(ctx, "a993ed18-9fc2-4217-a1a4-e860bddc3684")
	if err != nil {
		t.Errorf("VPN.Get returned %+v, expected %+v", err, nil)
	}

	expected := &VPN{
		ID:           "a993ed18-9fc2-4217-a1a4-e860bddc3684",
		VPCID:        "7ceab21c-e35b-4efb-9e96-6026ff5de5aa",
		NATGatewayID: "ab75b684-b3ca-443d-9550-65ce812cce6d",
		Label:        "govultr-test",
		Status:       "active",
		DateCreated:  "2026-10-13 12:36:04",
		Protocol:     "ipsec",
		Subnet:       "10.0.16.0",
		SubnetPrefix: 20,
		PublicKey:    "",
		Billing: VPNBilling{
			Charges: 0.04,
			Monthly: 25,
		},
	}

	if !reflect.DeepEqual(v, expected) {
		t.Errorf("ISO.Get returned %+v, expected %+v", v, expected)
	}
}

func TestVPNServiceHandler_Update(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpns/a993ed18-9fc2-4217-a1a4-e860bddc3684", testJSONResponseHandlerFunc(http.StatusOK, `
{
	"vpn": {
        "id": "a993ed18-9fc2-4217-a1a4-e860bddc3684",
        "vpc_id": "7ceab21c-e35b-4efb-9e96-6026ff5de5aa",
        "nat_gateway_id": "ab75b684-b3ca-443d-9550-65ce812cce6d",
        "label": "govultr-test-updated",
        "status": "active",
        "date_created": "2026-10-13 12:36:04",
        "protocol": "ipsec",
        "subnet": "10.0.16.0",
        "subnet_prefix": 20,
        "public_key": "",
        "billing": {
                "charges": 0.04,
                "monthly": 25
        }
	}
}
`))

	vpnReq := &VPNUpdateReq{Label: "govultr-test-updated"}

	v, _, err := client.VPN.Update(ctx, "a993ed18-9fc2-4217-a1a4-e860bddc3684", vpnReq)
	if err != nil {
		t.Errorf("VPN.Update returned %+v, expected %+v", err, nil)
	}

	expected := &VPN{
		ID:           "a993ed18-9fc2-4217-a1a4-e860bddc3684",
		VPCID:        "7ceab21c-e35b-4efb-9e96-6026ff5de5aa",
		NATGatewayID: "ab75b684-b3ca-443d-9550-65ce812cce6d",
		Label:        "govultr-test-updated",
		Status:       "active",
		DateCreated:  "2026-10-13 12:36:04",
		Protocol:     "ipsec",
		Subnet:       "10.0.16.0",
		SubnetPrefix: 20,
		PublicKey:    "",
		Billing: VPNBilling{
			Charges: 0.04,
			Monthly: 25,
		},
	}

	if !reflect.DeepEqual(v, expected) {
		t.Errorf("ISO.Update returned %+v, expected %+v", v, expected)
	}
}

func TestVPNServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v2/vpns/a993ed18-9fc2-4217-a1a4-e860bddc3684", testJSONResponseHandlerFunc(http.StatusNoContent, ""))

	if err := client.VPN.Delete(ctx, "a993ed18-9fc2-4217-a1a4-e860bddc3684"); err != nil {
		t.Errorf("VPN.Delete returned %+v, expected %+v", err, nil)
	}
}
